package main

import "fmt"

type st struct {
	school string
}

type person struct {
	name string
	age  int
	st
}

func main() {

	fmt.Println(person{
		name: "gp",
		age:  1,
		st:   st{school: "cs"},
	})

}
