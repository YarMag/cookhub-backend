package ga_operators


func getAffectedChromosomeBoundaries(gapPoint int, partsBoundsIndices []int) (int, int) {
	lowerBound := 0
	upperBound := 0
	for i := 0; i < len(partsBoundsIndices); i++ {
		upperBound = partsBoundsIndices[i]
		if (gapPoint >= lowerBound && gapPoint <= upperBound) {
			break
		}
		lowerBound = upperBound
	}

	return lowerBound, upperBound
}

func correctChromosome(chromosome []int, maxOnes int) {
	leftPtr := 0
	rightPtr := len(chromosome) - 1
	totalOnes := 0

	for ; leftPtr < rightPtr; {
		if chromosome[leftPtr] == 1 {
			totalOnes += 1
		}
		leftPtr++
		if totalOnes == maxOnes {
			break
		}
		if chromosome[rightPtr] == 1 {
			totalOnes += 1
		}
		rightPtr--
		if totalOnes == maxOnes {
			break
		}
	}

	for ; leftPtr < rightPtr; {
		chromosome[leftPtr] = 0
		chromosome[rightPtr] = 0
		leftPtr++
		rightPtr--
	}
}

func correctMutationGapPointIfNeeded(gapPoint int, partsBoundsIndices []int) int {
	for i := 0; i < len(partsBoundsIndices); i++ {
		if (gapPoint == partsBoundsIndices[i]) {
			return gapPoint - 1
		}
	}
	return gapPoint
}

func applyRestrictedInversion(chromosome []int, restrictions ChromosomeStructureRequirements) ([]int) {
	gapPoint := getSingleGap(len(chromosome))

	lowerBound, upperBound := getAffectedChromosomeBoundaries(gapPoint, restrictions.PartsBoundsIndices)
	
	result := applyInversion(chromosome[lowerBound:upperBound+1], gapPoint - lowerBound)
	correctChromosome(result, restrictions.MaxRecipesInPartsCount)
	
	childChromosome := make([]int, len(chromosome))
	copy(childChromosome[:], chromosome[:])
	copy(childChromosome[lowerBound:upperBound+1], result)

	return childChromosome
}

func applyRestrictedMutation(chromosome []int, restrictions ChromosomeStructureRequirements) ([]int) {
	gapPoint := getSingleGap(len(chromosome))
	gapPoint = correctMutationGapPointIfNeeded(gapPoint, restrictions.PartsBoundsIndices)
	lowerBound, upperBound := getAffectedChromosomeBoundaries(gapPoint, restrictions.PartsBoundsIndices)
	
	result := applyMutation(chromosome[lowerBound:upperBound+1], gapPoint - lowerBound)
	copy(chromosome[lowerBound:upperBound+1], result)
	
	childChromosome := make([]int, len(chromosome))
	copy(childChromosome[:], chromosome[:])
	copy(childChromosome[lowerBound:upperBound+1], result)

	return childChromosome
}

func applyRestrictedSinglePointCrossingover(firstParent []int, secondParent []int, restrictions ChromosomeStructureRequirements) ([]int, []int) {
	gapPointIndex := getSingleGap(len(restrictions.PartsBoundsIndices)-1) // because can't extract part after last bound
	gapPoint := restrictions.PartsBoundsIndices[gapPointIndex]

	chromosomeLen := len(firstParent)
	firstChild := make([]int, chromosomeLen)
	secondChild := make([]int, chromosomeLen)
	
	copy(firstChild[0:gapPoint+1], firstParent[0:gapPoint+1])
	copy(firstChild[gapPoint+1:], secondParent[gapPoint+1:])

	copy(secondChild[0:gapPoint+1], secondParent[0:gapPoint+1])
	copy(secondChild[gapPoint+1:], firstParent[gapPoint+1:])

	return firstChild, secondChild
}

func applyRestrictedTwoPointCrossingover(firstParent []int, secondParent []int, restrictions ChromosomeStructureRequirements) ([]int, []int) {
	// maybe one day i'll implement this. not this time, sorry
	chromosomeLen := len(firstParent)
	firstChild := make([]int, chromosomeLen)
	secondChild := make([]int, chromosomeLen)

	return firstChild, secondChild
}
