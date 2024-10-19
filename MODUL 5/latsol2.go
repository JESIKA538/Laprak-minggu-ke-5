package main

import "fmt"

func main() {
	var panjang_jejari, tinggi float64
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n ; i++ {
		fmt.Print("masukkan panjang, jejari alas, tinggi : ")
		fmt.Scan(&panjang_jejari, &tinggi)
		volume:= 1.0/3.0 * 3.14 * panjang_jejari*panjang_jejari* tinggi
		fmt.Println(volume)

	}
	


}