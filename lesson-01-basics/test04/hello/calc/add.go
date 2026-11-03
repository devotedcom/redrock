package calc

func Add(a, b int) int {
	return a + b
}

func helper(x int) int {
	return x * 2
}
func AddDoubled(a, b int) int {
	return helper(a) + helper(b)
}
