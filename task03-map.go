package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {

	temp := make([]int, 0, len(input))

	for i, _ := range input {
		temp = append(temp, i)
	}

	sort.Ints(temp)

	result = make([]string, 0, len(input))

	for r := range temp {
		result = append(result, input[temp[r]])
	}

	return
}
