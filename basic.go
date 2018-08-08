package main

import "fmt"

var (
	aa = 3
	bb = 4
)

func variableEmptyValue() {
	var a int
	var s string
	fmt.Println(a, s)
}

func variableInitValue() {
	var a int = 123
	var s string = "abc"
	fmt.Println(a, s)
}

func variableTypeValue() {
	var a, b, c, d = 123, 234, "abc", true
	fmt.Println(a, b, c, d)
}

func variableSimpleValue() {
	a, b, c, d := 123, 234, "abc", true
	fmt.Println(a, b, c, d)
}

func consts() {
	const (
		cpp = iota
		_
		php
		python
		jjj
	)
	fmt.Println(cpp, jjj, php, python)
}

func main() {
	fmt.Println("HelloWorld")
	variableEmptyValue()
	variableInitValue()
	variableTypeValue()
	variableSimpleValue()
	fmt.Println(aa, bb)
	consts()
}
