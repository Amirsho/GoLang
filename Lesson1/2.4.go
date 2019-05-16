package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println(a, "метр это", a / 1000, " километр")
}