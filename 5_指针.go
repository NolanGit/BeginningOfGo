package main

import "fmt"

func main() {
	a := 10
	var p *int
	p = &a
	fmt.Println(p)
	fmt.Println(*p)

	//空指针
	var ptr *string
	fmt.Println(ptr == nil)
}
