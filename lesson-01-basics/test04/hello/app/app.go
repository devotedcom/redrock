package app

import (
	"fmt"
	"redrock/teaching-exercises/lesson-01-basics/test04/hello/calc"
)

func Sum(a, b int) int {
	return calc.Add(a, b)
}

func Describe(a, b int) string {
	return fmt.Sprintf("%d + %d = %d", a, b, calc.Add(a, b))
}
