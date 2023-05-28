package ga_operators

import (
	"testing"
	"reflect"
	"sort"
)

func TestObtainingGap(t *testing.T) {
	chromosome := []int{0, 1, 0, 1, 1, 0, 1, 0, 0, 1}
	results := make([]int, len(chromosome))

	for index, _ := range chromosome {
		results[index] = getSingleGap(len(chromosome))	
	}

	base := results[0]
	shouldFail := true
	for i:=1; i < len(results); i++ {
		if results[i] != base {
			shouldFail = false
		}
	}

	if shouldFail {
		t.Error("Generated gap should not be always constant!")
	}

}

func TestObtainingMultipleGaps(t *testing.T) {
	chromosome := []int{0, 1, 0, 1, 1, 0, 1, 0, 0, 1}
	gapsCount := 4
	results := make([][]int, 10)

	for index, _ := range chromosome {
		results[index] = getMultipleGaps(len(chromosome), gapsCount)

		if len(results[index]) != gapsCount {
			t.Error("Generated incorrect amount of gaps!")
		}

		if !sort.SliceIsSorted(results[index], func (a, b int) bool { return a < b }) {
			t.Error("Gaps indices should be sorted ascending!")
		}
	}

	for _, gaps := range results {
		base := gaps[0]
		for i:=1; i < len(gaps); i++ {
			if gaps[i] == base {
				t.Error("Generated multiple gaps should never be equal!")
			}
		}
	}
}

func TestApplyingMutationOperator(t *testing.T) {
	chromosome := []int{0, 1, 0, 1, 1, 0, 1, 0, 0, 1}
	gapPoint := 4
	expected := []int{0, 1, 0, 1, 0, 1, 1, 0, 0, 1}

	actual := applyMutation(chromosome, gapPoint)


	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Mutation did not replaced items at gap point %d", gapPoint)
	}
}

func TestApplyingInversionOperator(t *testing.T) {
	chromosome := []int{0, 1, 0, 1, 1, 0, 1, 0, 0, 1}
	gapPoint := 4
	expected := []int{0, 1, 0, 1, 1, 0, 0, 1, 0, 1}

	actual := applyInversion(chromosome, gapPoint)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Inversion didn't affect gens after %d", gapPoint)
	}	
}

func TestSplittingChromosome(t *testing.T) {
	chromosome := []int{0, 1, 0, 1, 1, 0, 1, 0, 0, 1}
	gapIndices := []int{2, 7}

	divided := splitIntoParts(chromosome, gapIndices)

	if !(len(divided[0]) + len(divided[1]) + len(divided[2]) == len(chromosome)) {
		t.Error("Incorrect dimensions for generated parts!")
	}

	if !(reflect.DeepEqual(chromosome[0:3], divided[0]) && reflect.DeepEqual(chromosome[3:8], divided[1]) && 
		reflect.DeepEqual(chromosome[8:10], divided[2]))  {
		t.Errorf("Split into %d parts works incorrectly", len(gapIndices) + 1)
	}
}

func TestApplyingCrossingover(t *testing.T) {
	parent1 := []int{0, 1, 0, 1, 1, 1, 0, 1, 0, 0}
	parent2 := []int{1, 0, 1, 1, 1, 0, 0, 0, 1, 0}

	gapIndices := []int{2, 7}

	expectedChild1 := []int{0, 1, 0, 1, 1, 0, 0, 0, 0, 0}
	expectedChild2 := []int{1, 0, 1, 1, 1, 1, 0, 1, 1, 0}

	actualChild1, actualChild2 := applyCrossingover(parent1, parent2, gapIndices)

	if !reflect.DeepEqual(actualChild1, expectedChild1) {
		t.Error("First crossingover child is broken")
	}
	if !reflect.DeepEqual(actualChild2, expectedChild2) {
		t.Error("Second crossingover child is broken")
	}
}
