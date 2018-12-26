package main

import "fmt"

// iota  只能在定义常量中使用  代表自增  从0开始
/* 
const (
	a = iota
	b
	c
	d
)
*/

const (
	a = iota      // 0
	b = 3 << iota // 110 6
	c = 4 << iota // 10 1000 16
	d = 5 << iota // 101  101000     40
)
func main(){
	fmt.Println(a,b,c,d)
}