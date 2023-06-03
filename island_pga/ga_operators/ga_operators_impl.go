package ga_operators

import (
	"math/rand"
	"sort"
)

func getSingleGap(targetDimension int) (int) {
	return rand.Intn(targetDimension)
}

func getMultipleGaps(targetDimension int, gapsCount int) ([]int) {
	result := make([]int, 0)

	gapsMap := make(map[int]bool)
	for i:=0; i<gapsCount; i++ {
		isAdded := false
		for !isAdded {
			newGap := getSingleGap(targetDimension)
			if _, ok := gapsMap[newGap]; !ok {
				result = append(result, newGap)
				gapsMap[newGap] = true		
				isAdded = true
			}	
		}
		
		
	}
	return sort.IntSlice(result)
}

func applyMutation(chromosome []int, gapPoint int) ([]int) {
	result := make([]int, len(chromosome))

	for index, _ := range chromosome {
		switch index {
		case gapPoint:
			result[index] = chromosome[index+1]
			result[index+1] = chromosome[index]
		case gapPoint + 1:
			break // handled in previous index
		default:
			result[index] = chromosome[index]
		}
	}

	return result
}

func applyInversion(chromosome []int, gapPoint int) ([]int) {
	chromosomeLen := len(chromosome)
	result := make([]int, chromosomeLen)

	for i:=0; i < gapPoint; i++ {
		result[i] = chromosome[i]
	}
	for i:=0; i < chromosomeLen - gapPoint; i++ {
		result[gapPoint + i] = chromosome[chromosomeLen - i - 1]
	}

	return result
}

func splitIntoParts(chromosome []int, partsIndices []int) ([][]int) {
	divided := make([][]int, 0)
	extPartsIndices := append(partsIndices, len(chromosome) - 1)
	currIndex := 0
	for currPart := 0; currPart < len(extPartsIndices); currPart++ {
		divided = append(divided, chromosome[currIndex:(extPartsIndices[currPart]+1)])
		currIndex = extPartsIndices[currPart] + 1
	}
	return divided
}

func applyCrossingover(parent1 []int, parent2 []int, gapIndices []int) ([]int, []int) {
	parent1Parts := splitIntoParts(parent1, gapIndices)
	parent2Parts := splitIntoParts(parent2, gapIndices)

	child1 := make([]int, 0)
	child2 := make([]int, 0)

	for index, _ := range parent1Parts {
		if index % 2 == 0 {
			child1 = append(child1, parent1Parts[index]...)
			child2 = append(child2, parent2Parts[index]...)
		} else {
			child1 = append(child1, parent2Parts[index]...)
			child2 = append(child2, parent1Parts[index]...)
		}
	}

	return child1, child2
}

