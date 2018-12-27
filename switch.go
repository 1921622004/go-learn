package main

import "fmt"
/*
func selectFunc(){
	const a = 1
	const b = 2
	select {
	case a == 1:
		fmt.Println("a")
	case b == 2:
		fmt.Println("b")
	case a > b: 
		fmt.Println("a > b")
	default :
		fmt.Println("nothing")
	}
}
*/

func main(){
	const a = 1
	const b = 2
	switch  {
	// 不用写break  
	case a > b:
		fmt.Println("a > b")
	case a < b:
		fmt.Println("a < b")
	default : 
		fmt.Println("a == b")
	}
//	selectFunc()
}