package sort

import (
	"testing"

	"github.com/MarceloMPJ/algorithm/sort"
)

func TestSelectionSort(t *testing.T) {
	tested, actual, expected := SortingAlgorithmTest(sort.SelectionSort)

	if !tested {
		t.Errorf(errorMessage, actual, expected)
	}
}
