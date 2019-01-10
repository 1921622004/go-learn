package main

import "fmt"

func main(){
	fmt.Printf("start ")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Printf("end")
}