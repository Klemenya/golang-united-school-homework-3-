package homework

func average(input [15]float32) (result float32) {
	var sum float32

	for _, num := range input {
		sum += num
	}
	return sum/float32(len(input))
}
