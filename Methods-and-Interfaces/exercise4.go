package main

import "fmt"

type empty interface{}

func main(){
	var empty empty
	empty = 15

	fmt.Printf("type is : %T\n" , empty)

	empty = 15.45
	fmt.Printf("type is : %T\n" , empty)

	empty = []int{1,4,7}
	fmt.Printf("type is : %T\n" , empty)

	empty = append(empty.([]int), 10)
	fmt.Println("slice is : " , empty)
}