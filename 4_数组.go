package main

import "fmt"

var l = [10]int{1, 2, 3, 4}

func main() {
	l2 := [10]int{1, 2, 3, 4}
	l3 := [10]int{0: 1, 2: 2, 4: 3}
	fmt.Println(l)
	fmt.Println(l2)
	fmt.Println(l3)
	l3[6] = 4
	fmt.Println(l3)
	fmt.Println(l3[6])
}
