package main

import "fmt"

func main() {
	var n, sum int
	fmt.Print("masukkan bilangan bulat: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		sum += i

	}
	fmt.Printf("hasil penjumlahan dari 1 hingga %d adalah : %d\n", n, sum)
}
