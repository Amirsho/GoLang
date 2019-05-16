package main

import "fmt"

func main() {
	var k, n, d int
	fmt.Scan(&k)
	n = k % 7
	fmt.Println(k, "день года это", n, "если 1 января это понидельник")
	n = (k + 1) % 7
	fmt.Println(k, "день года это", n, "если 1 января это вторник")
	//не стал использовать if или string чтоб вывести день недели словом
	fmt.Scan(&d)
	n = (k + d - 1) % 7
	fmt.Println(k, "день года это", n, "если 1 января это ",d,"-ый день недели")
}
