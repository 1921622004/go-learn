package main

import "fmt"

var a = "lover"
var b string = "fucker"
var c bool


func main(){
	// 局部变量不可以声明不使用
	// var e string = "aaa"
	fmt.Println(a,b,c)
	// 也可以这样声明
	d := 1
	e := d
	e = 2
	fmt.Println(d,e)
}