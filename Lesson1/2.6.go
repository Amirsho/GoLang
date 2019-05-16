package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(n, "секунд с начало суток это", n/3600, "часов")
	fmt.Println((n%3600)/60, " минут прошло с начала очередного часа")
	fmt.Println(n%60, "секунд прошло с начала очередной минуты")
}
