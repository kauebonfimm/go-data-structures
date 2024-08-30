package array

func Arrays(values [10]int) [10]int {
	initial_array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // standard init

	for i := range values {
		values[i] = initial_array[i] // adding itens by position
	}

	return initial_array
}

func Slices(values [10]int) []int {
	initial_slice := make([]int, 0) // standard init

	for i := range values {
		initial_slice[i] = values[i]             // adding positions
		initial_slice = append(initial_slice, i) // push to last index
	}

	return initial_slice
}
