package island_algorithm

import (
	"log"
	. "ga_operators"
	"reflect"
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
			newExecutor, _ := master.ExecutorFactory.CreateExecutor(availableCriterionTypes[i % len(availableCriterionTypes)], params.RecipesIds)
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

		criteria := make([]CriterionEvaluator, len(availableCriterionTypes))
		for i := 0; i < len(criteria); i++ {
			criteria[i] = executors[i].GetCriterion()
		}

		paretoFront := master.getBestParetoRangedGroup(totalChromosomes, criteria)
		uniqueSolutions := master.filterUnique(paretoFront)

		master.optimizationResultOutput<-GeneticAlgorithmResult {
			BestSolutions: uniqueSolutions,
		}
	}(params)
	return master.optimizationResultOutput
}

func (master *SimpleIslandMasterImpl)getBestParetoRangedGroup(solutions []GeneticAlgorithmChromosome, criteria []CriterionEvaluator) []GeneticAlgorithmChromosome {
	solutionsCount := len(solutions)
	criteriaCount := len(criteria)

	// формируем начальные условия, по которым считаем доминируемость
	valuesMatrix := make([][]float32, solutionsCount)
	for i := 0; i < solutionsCount; i++ {
		valuesMatrix[i] = make([]float32, criteriaCount)
		for j := 0; j < criteriaCount; j++ {
			valuesMatrix[i][j] = criteria[j].EvaluateChromosome(solutions[i])
		}
	}

	chromosomeCriterionValuesCache := make(map[string][]float32)

	// задаем квадратную булеву матрицу отношений хромосом, чтобы вычислить недоминируемые решения
	relationshipsMatrix := make([][]bool, solutionsCount)
	for i := 0; i < solutionsCount; i++ {
		relationshipsMatrix[i] = make([]bool, solutionsCount)

		iCriterionValues, ok := chromosomeCriterionValuesCache[solutions[i].GetId()]
		if !ok {
			iCriterionValues = make([]float32, criteriaCount)
			for index, criterion := range criteria {
				iCriterionValues[index] = criterion.EvaluateChromosome(solutions[i])
			}
			chromosomeCriterionValuesCache[solutions[i].GetId()] = iCriterionValues
		}

		for j := 0; j < solutionsCount; j++ {
			if i == j {
				continue
			}
			jCriterionValues, ok := chromosomeCriterionValuesCache[solutions[j].GetId()]
			if !ok {
				jCriterionValues = make([]float32, criteriaCount)
				for index, criterion := range criteria {
					jCriterionValues[index] = criterion.EvaluateChromosome(solutions[j])
				}
				chromosomeCriterionValuesCache[solutions[j].GetId()] = jCriterionValues
			}

			comparisonResults := make([]int, criteriaCount)
			hasGreaterCriterionValue := false
			isNonDominated := false
			for k := 0; k < criteriaCount; k++ {
				comparisonResults[k] = criteria[k].CompareChromosomesFitnessValues(iCriterionValues[k], jCriterionValues[k])
				if comparisonResults[k] == 1 {
					hasGreaterCriterionValue = true
				} else if comparisonResults[k] == -1 {
					isNonDominated = true
					break
				} // ignore 0, equal values don't matter
			}
			
			if !isNonDominated && hasGreaterCriterionValue {
				relationshipsMatrix[i][j] = true
			}
		}
	}

	// ищем столбцы, где все значения false - это будут недоминируемые решения, составляющие первую группу парето-фронта
	paretoFrontIndices := make([]int, 0)
	for j := 0; j < solutionsCount; j++ {
		hasDominatedValues := false
		for i := 0; i < solutionsCount; i++ {
			if relationshipsMatrix[i][j] == true {
				hasDominatedValues = true
				break
			}
		}
		if !hasDominatedValues {
			paretoFrontIndices = append(paretoFrontIndices, j)
		}
	}
	log.Printf("IslandMaster::Detected %d solutions in pareto front", len(paretoFrontIndices))

	bestParetoChromosomeGroup := make([]GeneticAlgorithmChromosome, 0)
	for _, paretoSolutionIndex := range paretoFrontIndices {
		bestParetoChromosomeGroup = append(bestParetoChromosomeGroup, solutions[paretoSolutionIndex])
	}

	return bestParetoChromosomeGroup
}

func (master *SimpleIslandMasterImpl)filterUnique(solutions []GeneticAlgorithmChromosome) []GeneticAlgorithmChromosome {
	solutionsCount := len(solutions)
	log.Printf("IslandMaster::Start filtering %d chromosomes", solutionsCount)
	uniqueSolutions := make([]GeneticAlgorithmChromosome, 0)
	shouldBeDiscardedFlags := make([]bool, solutionsCount)
	for i := 0; i < solutionsCount; i++ {
		if !shouldBeDiscardedFlags[i] {
			uniqueSolutions = append(uniqueSolutions, solutions[i])
		}
		for j := i+1; j < solutionsCount; j++ {
			if reflect.DeepEqual(solutions[i].GetValues(), solutions[j].GetValues()) {
				shouldBeDiscardedFlags[j] = true
			}
		} 
	}
	log.Printf("IslandMaster::Finish filtering - obtained %d unique chromosomes", len(uniqueSolutions))
	return uniqueSolutions
}
