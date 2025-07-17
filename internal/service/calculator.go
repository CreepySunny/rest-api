package service

import "fmt"

func Add(a, b int) int {
	res := a + b
	return res
}

func Subtract(a, b int) int {
	res := a - b
	return res
}

func Multiply(a, b int) int {
	res := a * b
	return res
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	res := a / b
	return res, nil
}

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
