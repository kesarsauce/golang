package arrays

func SumArray(a []int) int {
	sum := 0
	for _, num := range a {
		sum += num
	}
	return sum
}

func SumAll(numberArrays ...[]int) []int {
	k := []int{}
	for _, numArray := range numberArrays {
		//fmt.Println(numArray)
		k = append(k, SumArray(numArray))
	}
	return k
}