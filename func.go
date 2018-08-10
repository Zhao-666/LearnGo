package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("什么鬼" + op)
	}
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Println(opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func swap(a, b int) {
	b, a = a, b
}

func main() {
	fmt.Println(eval(3, 4, "*"))
	fmt.Println(apply(pow, 3, 4))
}
