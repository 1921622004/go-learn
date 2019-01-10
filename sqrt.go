package main

import "fmt"

func sqrt(x float64) float64 {
	s := 1.0
	var z float64
	for z != s  {
		s = z
		z = s - (( s * s - x ) / (2 * x) ) 
	}
	return z
}

func main (){
	fmt.Println(sqrt(2))
}