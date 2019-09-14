package functions

// SumAndDiffOperations: This function returns 2 values
func SumAndDiffOperations(x, y int) (sum, different int) {
	return x + y, x - y
}

func MultiSum(args ...int) int {
	sum := 0
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}
