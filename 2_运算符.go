package main

import (
	"fmt"
)

const A int = 10
const B int = 20

func main() {
	C := 30
	if A > 10 {
		fmt.Println(A)
	} else {
		fmt.Println(B)
	}

	// switch 语句
	switch C {
	case 10:
		fmt.Println("case 1")
	case 20:
		fmt.Println("case 2")
	default:
		fmt.Println("case 3")
	}

	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Println(" x 的类型 :%T", i)
	case int:
		fmt.Println("x 是 int 型")
	case float64:
		fmt.Println("x 是 float64 型")
	case func(int) float64:
		fmt.Println("x 是 func(int) 型")
	case bool, string:
		fmt.Println("x 是 bool 或 string 型")
	default:
		fmt.Println("未知型")
	}

	//使用 fallthrough 会强制执行后面的 case 语句

	switch {
	case true:
		fmt.Println("case 1")
		fallthrough
	case false:
		fmt.Println(("case 2"))
	}

	// select 语句

}
