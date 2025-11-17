package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("%v %v\n", f, c)
}

//var p = f()
//
//func f() *int {
//	v := 1
//	return &v
//}
