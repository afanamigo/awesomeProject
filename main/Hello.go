package main

import "fmt"

var maxNum = 9

func main() {
	var s = "Hello World, Hello Golang! --insider main"
	fmt.Println(s)
	fmt.Println()

	fmt.Println("Counting From 0 ~ 9")
	for i := 0; i <= maxNum; i++ {
		fmt.Print(i)
		if i < maxNum {
			fmt.Print(",")
		}
	}

	fmt.Println()
	fmt.Println("Counting End.")
}
