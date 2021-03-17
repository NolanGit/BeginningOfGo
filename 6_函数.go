package main

import "fmt"

func max(num1 int, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func testCallback(f func()) {
	fmt.Println("before callback...")
	f()
}

func callback() {
	fmt.Println("in callback function...")
}
func main() {
	a := 1
	b := 2
	fmt.Println(max(a, b))
	testCallback(callback)

	add_func := add(1, 2)
	fmt.Println(add_func())
	fmt.Println(add_func())
	fmt.Println(add_func())
}

// 闭包使用方法
func add(x1, x2 int) func() (int, int) {
	i := 0
	return func() (int, int) {
		i++
		return i, x1 + x2
	}
}
