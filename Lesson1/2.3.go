package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println(a, "килограмм это", a / 1000, " тонн")
}