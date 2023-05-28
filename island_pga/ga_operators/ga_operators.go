package ga_operators

type GeneticOperator int

const (
	Inversion GeneticOperator = iota
	Mutation
	SinglePointCrossingover
	TwoPointCrossingover
)

type GeneticAlgorithmChromosome interface {
	GetValues() []int
	SetValues([]int)
}

type ChromosomeStructureRequirements struct {
	PartsBoundsIndices []int
	MaxRecipesInPartsCount int
}

type SimpleChromosome struct {
	Values []int
}

func (chromosome *SimpleChromosome)GetValues() []int {
	return chromosome.Values
}

func (chromosome *SimpleChromosome)SetValues(values []int) {
	chromosome.Values = values
}

func ApplyOperator(geneticOperator GeneticOperator, restrictions ChromosomeStructureRequirements, parents []GeneticAlgorithmChromosome) ([]GeneticAlgorithmChromosome) {
	var children []GeneticAlgorithmChromosome

	switch geneticOperator {
	case Inversion:
		children = make([]GeneticAlgorithmChromosome, 1)
		result := applyRestrictedInversion(parents[0].GetValues(), restrictions)
		children[0] = &SimpleChromosome { Values: result }
	case Mutation:
		children = make([]GeneticAlgorithmChromosome, 1)
		result := applyRestrictedMutation(parents[0].GetValues(), restrictions)
		children[0] = &SimpleChromosome { Values: result }
	case SinglePointCrossingover:
		children = make([]GeneticAlgorithmChromosome, 2)
		result0, result1 := applyRestrictedSinglePointCrossingover(parents[0].GetValues(), parents[1].GetValues(), restrictions)
		children[0] = &SimpleChromosome { Values: result0 }
		children[1] = &SimpleChromosome { Values: result1 }
	case TwoPointCrossingover:
		children = make([]GeneticAlgorithmChromosome, 2)
		result0, result1 := applyRestrictedTwoPointCrossingover(parents[0].GetValues(), parents[1].GetValues(), restrictions)
		children[0] = &SimpleChromosome { Values: result0 }
		children[1] = &SimpleChromosome { Values: result1 }
	}

	return children
}
