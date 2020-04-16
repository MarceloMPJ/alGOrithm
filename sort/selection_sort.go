package sort

// SelectionSort is a sorting algorithm with complexity O(n**2)
func SelectionSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		minimumIdx := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[minimumIdx] {
				minimumIdx = j
			}
		}
		array[i], array[minimumIdx] = array[minimumIdx], array[i]
	}
	return array
}
