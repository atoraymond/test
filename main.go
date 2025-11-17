package main

import "fmt"

func main(){
	var x = []int{10, 20, 30, 40, 50}
	var y = [5]int(x)
	fmt.Println(y)
}