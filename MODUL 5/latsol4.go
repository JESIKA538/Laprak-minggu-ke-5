package main

import "fmt" 

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
func main() {
	var n int
	fmt.Print("masukkan bilbul: ")
	fmt.Scan(&n)
	hasil:= factorial(n)
	fmt.Println(hasil)
		
	}

