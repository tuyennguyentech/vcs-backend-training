package main

import (
	"errors"
	"fmt"
)

// function trong go
func plus(a int, b int) int {
	return a + b
}

// variadic function, kiểu ...int dùng như slice []int
func sum(nums ...int) (total int) {
	for _, num := range nums {
		total += num
	}
	return
}

// multiple returns
func modulo(dividend int, divisor int) (quotient int, remainder int) {
	quotient = dividend / divisor
	remainder = dividend % divisor
	return
}

func innerFunc(phantom int) (err error) {
	if phantom == 0 {
		err = errors.New("phantom cannot be zero")
	}
	return
}

func outerFunc(phantom int) (err error) {
	err = innerFunc(phantom)
	if err != nil {
		err = fmt.Errorf("outerFunc error: %w", err)
	}
	return
}

func main() {
	_ = plus(1, 2)

	// variadic function có thể nhận vào 1 số lượng tham số bất kì
	_ = sum(1, 2)
	_ = sum(1, 2, 3)
	// có thể truyền vào slice theo cú pháp func(slice...)
	_ = sum([]int{1, 2, 3}...)
	_, _ = modulo(10, 4)
	inner := errors.New("inner error")
	outer := fmt.Errorf("Outer error: %w", inner)
	fmt.Println(outer)
}
