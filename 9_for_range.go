package main

import "fmt"

func main() {

	// 遍历切片
	nums := []int{1, 2, 3, 4, 5}
	var sum int
	for _, n := range nums {
		sum += n
	}
	fmt.Println(sum)
	for i, n := range nums {
		if n == 2 {
			fmt.Println(i)
		}
	}

	// 遍历map
	m := map[string]string{"a": "1"}
	for k := range m {
		fmt.Println(k)
	}

	// 坑
	a := []int{1, 2, 3, 4}
	var r []*int
	for _, v1 := range a {
		temp_v1 := v1
		r = append(r, &temp_v1)
	}
	for k, v2 := range r {
		fmt.Printf("r[%v]=%v\n", k, *v2)
	}
}
