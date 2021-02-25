package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	a1 = 3
	s2 = "jj"
	bb = true
)


func variableZeroValue()  {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue()  {
	var a, b int = 3, 4
	var s  string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction()  {
	var a, b, c , f = 3, 4, true, "def"
	fmt.Println(a, b, c, f)
}

func variableShorter()  {
	a, b, c, s := 3, 4, true, "def1"
	fmt.Println(a, b, c, s)
}

func euler()  {
	fmt.Println(cmplx.Exp(1i * math.Pi) + 1)
}

func triangle()  {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main()  {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(a1, bb, s2)
	euler()
	triangle()
}