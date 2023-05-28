package island_algorithm

import (
	. "ga_operators"
	"github.com/google/uuid"
	"math/rand"
	"log"
)

type SimpleIslandExecutorImpl struct {
	id string
	chMigrationInput chan IslandExecutorResult
	chMigrationOutput chan IslandExecutorResult
	criterion CriterionEvaluator
}

func NewSimpleIslandExecutorImpl(criterion CriterionEvaluator) IslandExecutor {
	return &SimpleIslandExecutorImpl {
		id: uuid.New().String(),
		criterion: criterion,
	}
}

func (executor *SimpleIslandExecutorImpl)RunCalculations(params IslandExecutorParams) (error) {
	go func(params IslandExecutorParams) {
		// генерация начальной популяции
		islandPopulation := executor.generateInitialPopulation(params)
		defaultPopulationSize := params.DefaultPopulationSize

		migrationsCount := params.TotalIterationsCount / params.MigrationIterationsCount
		// цикл от 0 до params.TotalIterationsCount
		for migrationIndex := 0; migrationIndex < migrationsCount; migrationIndex++ {
			for iterationIndex := 0; iterationIndex < params.MigrationIterationsCount; iterationIndex++ {
				// генерация вероятности для выбора генетического оператора
				// генерация процента хромосом, которые затронет оператор
				operatorToExecute := SelectOperator(params.Operators)
				affectedChromosomesCount := int(float32(defaultPopulationSize) * params.PopulationUpdateRate)

				childChromosomes := make([]GeneticAlgorithmChromosome, 0)
				for i := 0; i < affectedChromosomesCount; i++ {
					parents := make([]GeneticAlgorithmChromosome, 1)
					switch operatorToExecute {
					case Inversion, Mutation:
						parents[0] = islandPopulation[rand.Intn(defaultPopulationSize)]
					case SinglePointCrossingover, TwoPointCrossingover:
						firstParentIndex := rand.Intn(defaultPopulationSize)
						secondParentIndex := 0
						for {
							secondParentIndex = rand.Intn(defaultPopulationSize)
							if (secondParentIndex != firstParentIndex) {
								break
							}
						}

						parents[0] = islandPopulation[firstParentIndex]
						parents = append(parents, islandPopulation[secondParentIndex])
					}

					// применение генетических операторов
					result := ApplyOperator(operatorToExecute, params.Requirements, parents)

					childChromosomes = append(childChromosomes, result...)
				}

				childChromosomesCount := len(childChromosomes)
				actualPopulationSize := len(islandPopulation)
				totalChromosomesCount := childChromosomesCount + actualPopulationSize
				totalChromosomes := make([]GeneticAlgorithmChromosome, totalChromosomesCount)
				copy(totalChromosomes[0:actualPopulationSize+1], islandPopulation[:])
				copy(totalChromosomes[actualPopulationSize:], childChromosomes[:])
				
				executor.criterion.RangePopulation(totalChromosomes)
				islandPopulation = make([]GeneticAlgorithmChromosome, defaultPopulationSize)

				// поскольку мы породили len(childChromosomes) дочерних хромосом, нужно столько же исключить из результирующей популяции
				// пусть y - количество получившихся дочерних хромосом, тогда исключим y/2 лучших хромосом и y/2 худших
				// в итоге из отсортированного массива решений totalChromosomes будут скопированы диапазоны индексов [0;y/2],[y;n-y],[n-y/2;n]
				excludedChromosomesCount := childChromosomesCount / 2
				copy(islandPopulation[0:excludedChromosomesCount], totalChromosomes[0:excludedChromosomesCount])
				copy(islandPopulation[childChromosomesCount/2:defaultPopulationSize-excludedChromosomesCount], totalChromosomes[childChromosomesCount:totalChromosomesCount-childChromosomesCount])
				copy(islandPopulation[defaultPopulationSize-excludedChromosomesCount:], totalChromosomes[totalChromosomesCount - excludedChromosomesCount:])
			}

			// отбор хромосом для миграции
			// для простоты отбираем migrationSize/2 лучших и столько же худших (плохое по одному критерию может быть хорошим по другому). Массив уже отсортирован, так что задача облегчается
			migrationChromosomes := make([]GeneticAlgorithmChromosome, params.MigrationPopulationSize)
			limit := params.MigrationPopulationSize / 2
			copy(migrationChromosomes[0:limit], islandPopulation[0:limit])
			copy(migrationChromosomes[limit:], islandPopulation[defaultPopulationSize - limit:])
			log.Printf("IslandExecutor::Executor %s sending chromosomes to channel", executor.GetId())
			// отправка хромосом в канал миграции
			executor.chMigrationOutput<-IslandExecutorResult{
				SenderId: executor.id,
				Chromosomes: migrationChromosomes,
				IsTerminal: false,
			}
			log.Printf("IslandExecutor::Executor %s::receiving chromosomes", executor.GetId())
			// прием хромосом из другого канала
			incomingChromosomesData := <-executor.chMigrationInput
			log.Printf("IslandExecutor::Executor %s::received migration data", executor.GetId())
			// замещаем migrationSize хромосом с конца - замещаем худшие решения, по сути
			copy(islandPopulation[defaultPopulationSize - len(incomingChromosomesData.Chromosomes):], incomingChromosomesData.Chromosomes[:])
		}
		// проранжировать популяцию
		// отобрать наилучшие решения
		// отправить решения в созданный канал	
		executor.criterion.RangePopulation(islandPopulation)
		log.Printf("IslandExecutor::Executor %s::sends final data of %d chromosomes", executor.GetId(), len(islandPopulation))
		finalChromosomes := make([]GeneticAlgorithmChromosome, params.MigrationPopulationSize)
		copy(finalChromosomes[0:params.MigrationPopulationSize], islandPopulation[0:params.MigrationPopulationSize])
		executor.chMigrationOutput<-IslandExecutorResult{
			SenderId: executor.id,
			Chromosomes: finalChromosomes,
			IsTerminal: true,
		}
		
	}(params)

	return nil
}

func (executor *SimpleIslandExecutorImpl)SetupMigrationChannel(output chan IslandExecutorResult) (chan IslandExecutorResult, error) {
	executor.chMigrationInput = make(chan IslandExecutorResult)
	executor.chMigrationOutput = output

	return executor.chMigrationInput, nil
}

func (executor *SimpleIslandExecutorImpl)IsReadyToWork() bool {
	return executor.chMigrationInput != nil && executor.chMigrationOutput != nil
}

func (executor *SimpleIslandExecutorImpl)GetId() string {
	return executor.id
}

func (executor *SimpleIslandExecutorImpl)generateInitialPopulation(params IslandExecutorParams) ([]GeneticAlgorithmChromosome) {
	chromosomeLength := params.Requirements.PartsBoundsIndices[len(params.Requirements.PartsBoundsIndices) - 1] + 1

	chromosomes := make([]GeneticAlgorithmChromosome, params.DefaultPopulationSize)
	for chromosomeIndex := 0; chromosomeIndex < params.DefaultPopulationSize; chromosomeIndex++ {
		chromosomeValues := make([]int, chromosomeLength)

		lowerBound := 0
		for i := 0; i < len(params.Requirements.PartsBoundsIndices); i++ {
			upperBound := params.Requirements.PartsBoundsIndices[i]
			for j := 0; j < params.Requirements.MaxRecipesInPartsCount; j++ {
				valuableIndex := rand.Intn(upperBound - lowerBound) + lowerBound
				chromosomeValues[valuableIndex] = 1	
			}
			lowerBound = upperBound
		}
		chromosomes[chromosomeIndex] = &SimpleChromosome {Values: chromosomeValues}
	}

	return chromosomes
}
