package main

import "fmt"

func main() {
	s := "中文测试!"
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Printf("%s\n", []byte(s))
	fmt.Println()
	for i, ch := range s { //ch is a rune
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) { //ch is a rune
		fmt.Printf("(%d %c)", i, ch)
	}
	fmt.Println()
}
