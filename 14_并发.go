package main

import (
	"fmt"
)

/*
// 并发
func say(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func main() {
	go say("1")
	say("2")
}
*/

// 通道
func sum(l []int, c chan int) {
	sum := 0
	for _, i := range l {
		sum += i
	}
	c <- sum
	fmt.Println("result is", sum)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {

	/* 不带缓冲区的通道
	当有数据发送至管道后，发送端将阻塞直到数据被接收端取出 */
	l1 := []int{1, 2, 3, 4, 5} // 15
	l2 := []int{2, 3, 4, 5, 6} // 20
	c := make(chan int)
	// 通道先进后出，后进先出
	go sum(l1, c)
	go sum(l2, c)
	r1, r2 := <-c, <-c
	fmt.Println(r1)
	fmt.Println(r2)

	/* 带缓冲区的通道
	当发送端发送的数据在缓冲区的空间内的时候，发送端不会被阻塞 */
	c2 := make(chan int, 2)
	c2 <- 1
	c2 <- 2
	r1, r2 = <-c2, <-c2
	fmt.Println("r1, r2")
	fmt.Println(r1, r2)

	/* 管道关闭
	使用range遍历管道时， 使用close(chan)可以关闭管道，从而停止遍历 */
	c3 := make(chan int)
	go fibonacci(10, c3)
	for i := range c3 {
		fmt.Println(i)
	}
}
