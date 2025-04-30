package main

import "fmt"

func main() {
	var i int8 = 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := range 3 {
		fmt.Println(j)
	}
}
