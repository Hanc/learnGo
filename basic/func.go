package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func sum() int {
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	return sum
}

func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling  function %s with args"+"(%d, %d)", opName, a, b)
	return op(a, b)
}

func sumArgs(number ...int) int {
	sum := 0
	for i := range number {
		sum += number[i]
	}
	return sum
}

func main() {
	fmt.Println(sum())

	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(
			float64(a), float64(b)))
	}, 3, 4))

	fmt.Println(sumArgs(1, 2, 3, 5, 6))
}
