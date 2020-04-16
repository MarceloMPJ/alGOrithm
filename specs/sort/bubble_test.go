package sort

import (
	"testing"

	"github.com/MarceloMPJ/algorithm/sort"
)

func TestBubble(t *testing.T) {
	tested, actual, expected := SortingAlgorithmTest(sort.Bubble)

	if !tested {
		t.Errorf(errorMessage, actual, expected)
	}
}
