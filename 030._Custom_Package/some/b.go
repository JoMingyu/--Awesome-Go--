package some

func Sum2(a ...int) int {
	sum := 0

	for _, value := range a {
		sum += value
	}

	return sum
}
