package main

import (
	"fmt"
)

type bird struct {
	num_of_wings int
	color        string
	name         string
}

func main() {
	bird_a := bird{2, "yellow", "a"}
	fmt.Println(bird_a.fly())
	fmt.Println(bird_a.color)

	bird_b := bird{num_of_wings: 2, color: "yellow"}
	fmt.Println(bird_b.fly())
	fmt.Println(bird_b.name)

	ptr := &bird_a
	get_bird_details(ptr)

}

func (b bird) fly() string {
	return b.name + " is flying"
}

func get_bird_details(bird *bird) {
	fmt.Println(bird.name)
	fmt.Println(bird.color)
	fmt.Println(bird.num_of_wings)
}
