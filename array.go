package main 

import "fmt"

func main(){
	ary := []int{2,3,4,5,6}
	fmt.Println("ary == ", ary);
	for i := 0; i < len(ary); i++ {
		fmt.Printf("ary[%d] = %d", i, ary[i])
	}
}