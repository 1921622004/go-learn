package main 

import "fmt"

func main(){
	i, j := 2, 27
	p := &i
	// & 生成指向其作用对象的指针
	fmt.Println(&p)
	// * 表示指针指向的底层的值
	*p = 1
	fmt.Println(i)

	p = &j
	*p = *p / 3
	fmt.Println(j)
	fmt.Println(i)
}