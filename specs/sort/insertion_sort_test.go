package sort

import (
	"testing"

	"github.com/MarceloMPJ/algorithm/sort"
)

func TestInsertionSort(t *testing.T) {
	tested, actual, expected := SortingAlgorithmTest(sort.InsertionSort)

	if !tested {
		t.Errorf(errorMessage, actual, expected)
	}
}
