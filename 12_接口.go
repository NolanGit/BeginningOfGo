package main

import "fmt"

/*
// 有参数
type Animal interface {
	eat(food string)
}

type Cat struct {
	name string
}

type Dog struct {
	name string
}

func (dog Dog) eat(food string) {
	fmt.Printf("%s is eating %s like a dog.\n", dog.name, food)
}

func (cat Cat) eat(food string) {
	fmt.Printf("%s is eating %s like a cat.\n", cat.name, food)
}

func main() {
	var a Animal
	a = Cat{"miaomiao"}
	a = Dog{"wangwang"}
	a.eat("fish")
	a.eat("shit")

}
*/

// 无参数
type Phone interface {
	call()
}

type iPhone struct{}

type Mi struct{}

func (iPhone) call() {
	fmt.Println("iPhone is calling.")
}

func (Mi) call() {
	fmt.Println("Mi is calling.")
}

func main() {
	var phone Phone
	phone = new(iPhone)
	phone.call()

	phone = new(Mi)
	phone.call()
}
