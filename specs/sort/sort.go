package sort

import "reflect"

const errorMessage = "Error! Actual: %v, and Expected = %v"

var tableTest = [][][]int{
	{
		{9, 8, 7, 6, 5, 4, 3, 2, 1},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
	{
		{3, 3, 3, 3, 4, 5, 1, 2, 2, 2, 2},
		{1, 2, 2, 2, 2, 3, 3, 3, 3, 4, 5},
	},
}

/*
SortingAlgorithmTest is a function to test sorting algoritms
returning true if all test passed
returning false if any test not passed
*/
func SortingAlgorithmTest(sortAlgorithm func(array []int) []int) (bool, []int, []int) {
	for _, intSlice := range tableTest {
		actual := sortAlgorithm(intSlice[0])
		expected := intSlice[1]
		if !reflect.DeepEqual(actual, expected) {
			return false, actual, expected
		}
	}
	return true, nil, nil
}
