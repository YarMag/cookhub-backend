package island_algorithm

import (
	. "ga_operators"
)

type MinPriceCriterionEvaluator struct {
	recipeInfoExtractor RecipeInfoExtractor
}

func (criterion *MinPriceCriterionEvaluator)RangePopulation(population []GeneticAlgorithmChromosome) {

}

type MaxProteinsCriterionEvaluator struct {
	recipeInfoExtractor RecipeInfoExtractor
}

func (criterion *MaxProteinsCriterionEvaluator)RangePopulation(population []GeneticAlgorithmChromosome) {

}

type MinCookingTimeCriterionEvaluator struct {
	recipeInfoExtractor RecipeInfoExtractor
}

func (criterion *MinCookingTimeCriterionEvaluator)RangePopulation(population []GeneticAlgorithmChromosome) {
	
}