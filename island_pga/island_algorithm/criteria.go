package island_algorithm

import (
	. "ga_operators"
	"sort"
	"log"
	"math"
)

const float64EqualityThreshold float64 = 1e-9

type MinPriceCriterionEvaluator struct {
	recipeInfoExtractor RecipeInfoExtractor
	recipeIds []int
}

func (criterion *MinPriceCriterionEvaluator)RangePopulation(population []GeneticAlgorithmChromosome) {
	calculatedValues := make(map[string]int)
	sort.Slice(population, func (i int, j int) bool {
		iId := population[i].GetId()
		iValue, ok := calculatedValues[population[i].GetId()]
		if !ok {
			iValue = criterion.evaluateChromosome(population[i])
			calculatedValues[iId] = iValue
		}

		jId := population[j].GetId()
		jValue, ok := calculatedValues[population[j].GetId()]
		if !ok {
			jValue = criterion.evaluateChromosome(population[j])
			calculatedValues[jId] = jValue
		}

		return iValue < jValue
	})
}

func (criterion *MinPriceCriterionEvaluator)evaluateChromosome(chromosome GeneticAlgorithmChromosome) int {
	chromosomeValues := chromosome.GetValues()
	accum := 0
	for index, val := range chromosomeValues {
		if val == 1 {
			recipeValue, err := criterion.recipeInfoExtractor.GetRecipePrice(criterion.recipeIds[index])
			if err != nil {
				log.Printf("MinPriceCriterionEvaluator::evaluateChromosome error %s", err)
			}
			accum += recipeValue
		}
	}
	return accum
}

func (criterion *MinPriceCriterionEvaluator)EvaluateChromosome(chromosome GeneticAlgorithmChromosome) float32 {
	return float32(criterion.evaluateChromosome(chromosome))
}

func (criterion *MinPriceCriterionEvaluator)CompareChromosomesFitnessValues(firstFV float32, secondFV float32) int {
	diff := float64(firstFV - secondFV)
	if diff > float64EqualityThreshold {
		return 1
	} else if math.Abs(diff) < float64EqualityThreshold {
		return 0
	} else {
		return -1
	}
}

type MaxProteinsCriterionEvaluator struct {
	recipeInfoExtractor RecipeInfoExtractor
	recipeIds []int
}

func (criterion *MaxProteinsCriterionEvaluator)RangePopulation(population []GeneticAlgorithmChromosome) {
	calculatedValues := make(map[string]float32)
	sort.Slice(population, func (i int, j int) bool {
		iId := population[i].GetId()
		iValue, ok := calculatedValues[population[i].GetId()]
		if !ok {
			iValue = criterion.evaluateChromosome(population[i])
			calculatedValues[iId] = iValue
		}

		jId := population[j].GetId()
		jValue, ok := calculatedValues[population[j].GetId()]
		if !ok {
			jValue = criterion.evaluateChromosome(population[j])
			calculatedValues[jId] = jValue
		}

		return jValue < iValue
	})
}

func (criterion *MaxProteinsCriterionEvaluator)evaluateChromosome(chromosome GeneticAlgorithmChromosome) float32 {
	chromosomeValues := chromosome.GetValues()
	var accum float32
	accum = 0
	for index, val := range chromosomeValues {
		if val == 1 {
			recipeValue, err := criterion.recipeInfoExtractor.GetRecipeProteins(criterion.recipeIds[index])
			if err != nil {
				log.Printf("MaxProteinsCriterionEvaluator::evaluateChromosome error %s", err)
			}
			accum += recipeValue
		}
	}
	return accum
}

func (criterion *MaxProteinsCriterionEvaluator)EvaluateChromosome(chromosome GeneticAlgorithmChromosome) float32 {
	return criterion.evaluateChromosome(chromosome)
}

func (criterion *MaxProteinsCriterionEvaluator)CompareChromosomesFitnessValues(firstFV float32, secondFV float32) int {
	diff := float64(firstFV - secondFV)
	if diff > float64EqualityThreshold {
		return -1
	} else if math.Abs(diff) < float64EqualityThreshold {
		return 0
	} else {
		return 1
	}
}

type MinCookingTimeCriterionEvaluator struct {
	recipeInfoExtractor RecipeInfoExtractor
	recipeIds []int
}

func (criterion *MinCookingTimeCriterionEvaluator)RangePopulation(population []GeneticAlgorithmChromosome) {
	calculatedValues := make(map[string]int)
	sort.Slice(population, func (i int, j int) bool {
		iId := population[i].GetId()
		iValue, ok := calculatedValues[population[i].GetId()]
		if !ok {
			iValue = criterion.evaluateChromosome(population[i])
			calculatedValues[iId] = iValue
		}

		jId := population[j].GetId()
		jValue, ok := calculatedValues[population[j].GetId()]
		if !ok {
			jValue = criterion.evaluateChromosome(population[j])
			calculatedValues[jId] = jValue
		}

		return iValue < jValue
	})	
}

func (criterion *MinCookingTimeCriterionEvaluator)EvaluateChromosome(chromosome GeneticAlgorithmChromosome) float32 {
	return float32(criterion.evaluateChromosome(chromosome))
}

func (criterion *MinCookingTimeCriterionEvaluator)CompareChromosomesFitnessValues(firstFV float32, secondFV float32) int {
	diff := float64(firstFV - secondFV)
	if diff > float64EqualityThreshold {
		return 1
	} else if math.Abs(diff) < float64EqualityThreshold {
		return 0
	} else {
		return -1
	}
}

func (criterion *MinCookingTimeCriterionEvaluator)evaluateChromosome(chromosome GeneticAlgorithmChromosome) int {
	chromosomeValues := chromosome.GetValues()
	accum := 0
	for index, val := range chromosomeValues {
		if val == 1 {
			recipeValue, err := criterion.recipeInfoExtractor.GetRecipeCookingTime(criterion.recipeIds[index])
			if err != nil {
				log.Printf("MinCookingTimeCriterionEvaluator::evaluateChromosome error %s", err)
			}
			accum += recipeValue
		}
	}
	return accum
}
