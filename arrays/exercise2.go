package main

import "fmt"



func main(){
	nums := [...]int{30, -1, -6, 90, -6}
	// Write a Go program that counts the number of positive even numbers in the array.
	positiveCounter := 0

	for _ , v := range nums{
		value := v
		if value > 0 {
			positiveCounter++
		}
	}

	// print result of positive numbers

	fmt.Println("The number of positive number is : " , positiveCounter)
}