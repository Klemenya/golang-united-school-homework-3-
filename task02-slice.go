package homework

func reverse(input []int64) (result []int64) {
	l := len(input)

	result = make([]int64, len(input))
	copy(result, input)

	for i := 0; i < l/2; i++ {
		result[i], result[l-1-i] = result[l-1-i], result[i]
	}

	return result
}
