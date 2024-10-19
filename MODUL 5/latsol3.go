package main

import "fmt"
import "math"

func main() {
	var  n, t float64
	fmt.Scan(&n, &t)
	pangkat:= math.Pow(n, t)
	fmt.Print(pangkat)
	}
