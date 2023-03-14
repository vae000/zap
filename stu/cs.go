package main

import "fmt"

func main() {
	var x string = "cs2"
	//switch x {
	//case "cs":
	//	fmt.Println("success")
	//case "cs2":
	//	fmt.Println("faled")
	//
	//}
	for i, v := range x {
		fmt.Println(i, v)
	}

}
