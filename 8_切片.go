package main

import "fmt"

var slice []int

func main() {
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	fmt.Println(slice2)
	s3 := slice2[:]
	fmt.Println(s3)
	s3 = slice2[1:3]
	fmt.Println(s3)
	s3 = slice2[:3]
	fmt.Println(s3)
	s3 = slice2[1:]
	fmt.Println(s3)
	s3 = append(s3, 12345)
	fmt.Println(s3)

	var numbers = make([]int, 3, 5) // 3是长度，5是容量
	printSlice(numbers)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
