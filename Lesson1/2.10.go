package main

import "fmt"

func main() {
	var n, a, b, s int
	fmt.Scan(&n)
	a = n / 10
	b = n % 10
	s = a + b
	fmt.Println(a,"десятков")
	fmt.Println(b,"едениц")
	fmt.Println(s,"сумма цифр")
}
