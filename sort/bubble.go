package sort

// Bubble Sort is the simplest sorting algorithm with complexity O(n**2)
func Bubble(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			if array[j] > array[i] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array
}
