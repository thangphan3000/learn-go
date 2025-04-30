package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	// NOTE: we need to think about what type of int we need to improving the memory use in our app
	// it can be int8/ int16/ int32/ int64
	var b, c int8 = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// NOTE: alternative type: int, string
	var e bool
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}
