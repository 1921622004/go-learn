package main

import "fmt"

func main(){
	const LENGTH int = 10
	const HEIGHT int = 20
	const WEIGHT int = 30
	area := LENGTH * WEIGHT
	volume := LENGTH * WEIGHT * HEIGHT
	fmt.Printf("面积为 %d, 体积为%d", area, volume)
}