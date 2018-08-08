package main

import (
	"io/ioutil"
	"fmt"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "F"
	case score < 70:
		g = "C"
	case score < 90:
		g = "b"
	default:
		panic(fmt.Sprintf("Wrong score: %d", score))
	}
	return g
}

func main() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(grade(0),grade(80),grade(111))
}
