package main

import "fmt"

var a int = 12

const asd = "const123"

const (
	z = 2 * iota
	v
	n
)

func main() {
	b := "qwe"
	fmt.Println(b)
	var c = b
	c = "wer"
	fmt.Println(c)
	fmt.Println("Hello, World!")
	fmt.Println(asd)
	fmt.Println(len(asd))
	fmt.Println(z)
	fmt.Println(v)
	fmt.Println(n)
}
