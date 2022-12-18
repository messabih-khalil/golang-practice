package main

import "fmt"

// Create a new type called money. Its underlying type is float64.

type Money float64


// Create a method called print that prints out the money value with exact 2 decimal points.

func (m Money) print(){
	fmt.Printf("%.2f\n", m)
}


func main(){
	moneys := []Money{12.2545478 , 578.455 , 7895.78545}

	for _ , val := range moneys{
		val.print()
	}
}