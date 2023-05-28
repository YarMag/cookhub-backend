package island_algorithm

import (
	"log"
	. "ga_operators"
)

type executorMigrationServiceInfo struct {
	InputChannel chan IslandExecutorResult
	LastReceivedResult IslandExecutorResult
	Id string
}

type SimpleIslandMasterImpl struct {
	ExecutorFactory IslandExecutorFactory
	optimizationResultOutput chan GeneticAlgorithmResult
	executorsInfo []executorMigrationServiceInfo
	executorsIdsMap map[string]int
}

const (
	populationUpdateRate float32 = 0.25
)


func (master *SimpleIslandMasterImpl)StartOptimization(params GeneticAlgorithmParams) (chan GeneticAlgorithmResult) {
	master.optimizationResultOutput = make(chan GeneticAlgorithmResult)

	go func(params GeneticAlgorithmParams) {
		// создать исполнителей с помощью фабрики
		// сформировать для каждого исполнителя начальные параметры для работы
		// отправить задания на выполнение
		// сохранить каналы для выходных данных на миграцию

		executors := []IslandExecutor{}
		
		masterInputChannel := make(chan IslandExecutorResult)
		master.executorsIdsMap = make(map[string]int)
		availableCriterionTypes := []OptimizationCriterionType{ MinPrice, MaxProteins, MinCookingTime }
		for i := 0; i < params.IslandsCount; i++ {
			newExecutor, _ := master.ExecutorFactory.CreateExecutor(availableCriterionTypes[i % len(availableCriterionTypes)])
			executors = append(executors, newExecutor)
			master.executorsIdsMap[newExecutor.GetId()] = i
			//log.Printf("IslandMaster::Created executor with id %s", newExecutor.GetId())

			executorInfo := executorMigrationServiceInfo{}
			migrationChannel, _ := newExecutor.SetupMigrationChannel(masterInputChannel)
			executorInfo.InputChannel = migrationChannel
			executorInfo.Id = newExecutor.GetId()
			master.executorsInfo = append(master.executorsInfo, executorInfo)

			if (newExecutor.IsReadyToWork()) {
				executorParams := IslandExecutorParams {
					TotalIterationsCount: params.TotalIterationsCount,
					MigrationIterationsCount: params.MigrationIterationsCount,
					Operators: params.Operators,
					Requirements: params.Requirements,
					PopulationUpdateRate: populationUpdateRate,
					MigrationPopulationSize: 10,
					DefaultPopulationSize: 40,
				}
				//log.Println("IslandMaster::Running executor")
				_ = newExecutor.RunCalculations(executorParams)
			} else {
				//log.Println("IslandMaster::Never started executor")
				// need somehow mark/restart failed executors
			}
		}


		// рассчитать количество итераций миграции - params.TotalIterationsCount / params.MigrationIterationsCount
		// цикл от 0 до количества итераций миграции
			// дождаться промежуточных результатов от каждого исполнителя
			// выполнить обмен данными
		migrationsIterations := params.TotalIterationsCount / params.MigrationIterationsCount
		activeExecutorsCount := len(master.executorsInfo)
		
		executorLeftMessagesCountMap := make(map[string]int)
		for i := 0; i < activeExecutorsCount; i++ {
			executorLeftMessagesCountMap[master.executorsInfo[i].Id] = migrationsIterations
		}

		totalChromosomes := make([]GeneticAlgorithmChromosome, 0)
		overallMessagesNumber := (migrationsIterations + 1) * activeExecutorsCount // migrations and one for total chromosomes appending
		for i := 0; i < overallMessagesNumber; i++ {
			log.Print("IslandMaster::Receiving...")
			curRes := <-masterInputChannel
			//log.Printf("IslandMaster::Received data from executor %s", curRes.SenderId)
			
			if executorLeftMessagesCountMap[curRes.SenderId] > 0 {
				// simple strategy - just send to next
				recieverExecutorIndex := (master.executorsIdsMap[curRes.SenderId] + 1) % activeExecutorsCount
				//log.Printf("IslandMaster::Sending data to next executor with index %d", recieverExecutorIndex)
				
				go func() {
					master.executorsInfo[recieverExecutorIndex].InputChannel <- IslandExecutorResult { SenderId: "", Chromosomes: curRes.Chromosomes }
					log.Print("IslandMaster::Successfully send")
				}()	
				
			} else {
				log.Printf("IslandMaster::Received final data with %d chromosomes.", len(curRes.Chromosomes))
				totalChromosomes = append(totalChromosomes, curRes.Chromosomes...)
			}
			executorLeftMessagesCountMap[curRes.SenderId] -= 1
		}
		
		
		// дождаться итоговых результатов от каждого исполнителя
		log.Printf("IslandMaster::received %d chromosomes from all executors", len(totalChromosomes))
		// сформировать Парето-фронт
		// отправить результат в master.optimizationResultOutput

		master.optimizationResultOutput<-GeneticAlgorithmResult {
			BestSolutions: totalChromosomes[0:11],
		}
	}(params)
	return master.optimizationResultOutput
}
