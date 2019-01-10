package main

import "fmt"

// js中的对象？

type Vertex struct {
	X int
	Y int
}

func main(){
	v := Vertex{1,2} 
	v.X = 33
	p := &v
	p.X = 1e9
	fmt.Println(v.X)
}