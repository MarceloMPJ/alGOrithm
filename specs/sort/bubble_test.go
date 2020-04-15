package sort

import (
	"reflect"
	"testing"

	"github.com/MarceloMPJ/algorithm/sort"
)

func TestBubble(t *testing.T) {
	for _, intSlice := range tableTest {
		actual := sort.Bubble(intSlice[0])
		expected := intSlice[1]
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf(errorMessage, actual, expected)
		}
	}
}
