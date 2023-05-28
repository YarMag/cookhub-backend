package island_algorithm

import (
	. "ga_operators"
	"math/rand"
)

type OptimizationCriterionType int

const (
	MinPrice OptimizationCriterionType = iota
	MinCookingTime
	MaxProteins
)

/*
	Calculates chromosome value from the point of optimization criterion
	Requires recipes metadata (food values, costs, other properties)
	Needs to know how to map passed chromosome to actual recipes set 
*/
type CriterionEvaluator interface {
	RangePopulation([]GeneticAlgorithmChromosome)
	EvaluateChromosome(GeneticAlgorithmChromosome) float32
}

type OptimizationCriteriaFactory interface {
	CreateCriterion(OptimizationCriterionType, []int) CriterionEvaluator
}

/*
	Initial parameters for GA work
*/
type GeneticAlgorithmParams struct {
	TotalIterationsCount int
	MigrationIterationsCount int
	IslandsCount int
	Operators []GeneticOperatorMetadata
	Requirements ChromosomeStructureRequirements
	RecipesIds []int
}

type GeneticAlgorithmResult struct {
	BestSolutions []GeneticAlgorithmChromosome
}

type IslandExecutorParams struct {
	TotalIterationsCount int
	MigrationIterationsCount int
	Operators []GeneticOperatorMetadata
	Requirements ChromosomeStructureRequirements
	PopulationUpdateRate float32 // [0;1] - percent of updated chromosomes in population via selected genetic operator
	RecipesIds []int
	DefaultPopulationSize int
	MigrationPopulationSize int
}

type IslandExecutorResult struct {
	SenderId string
	Chromosomes []GeneticAlgorithmChromosome
	IsTerminal bool
}

/*
type IslandExecutor interface {
	runs with some initial population
	internally executes genetic operators with predefined probabilities
	send back data for migration
	returns some subpopulation as the result
}
*/
type IslandExecutor interface {
	RunCalculations(IslandExecutorParams) (error)
	SetupMigrationChannel(chan IslandExecutorResult) (chan IslandExecutorResult, error)
	IsReadyToWork() bool
	GetId() string
}

type IslandExecutorFactory interface {
	CreateExecutor(OptimizationCriterionType, []int) (IslandExecutor, error)
}

/*
type IslandMaster interface {
	acceps initial parameters
	splits large recipes set into small parts and distributes to executors
	generates required amount of executors
	rules the migration and shuffles migrations blocks between executors
}
*/
type IslandMaster interface {
	StartOptimization(GeneticAlgorithmParams) (chan GeneticAlgorithmResult)
}

type GeneticOperatorMetadata struct {
	Type GeneticOperator
	Probability float32 // [0; 1]
}

func SelectOperator(operators []GeneticOperatorMetadata) GeneticOperator {
	generatedProbability := rand.Float32()
	var accum float32
	accum = 0
	var resultOperator GeneticOperator
	for _, operatorInfo := range operators {
		accum += operatorInfo.Probability
		if generatedProbability <= accum {
			resultOperator = operatorInfo.Type
			break
		}
	}
	return resultOperator
}
