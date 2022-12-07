package main

import "fmt"

func main(){
	// cities array
	var cities  [2]string
	// print cities
	fmt.Printf("%#v" , cities)
	// 2. Declare an array called grades of type [3]float64 and initialize only the first 2 elements using an array literal.
	var grades [3]float64
	grades[0] = 1.5
	grades[1] = 1.45

	// Print out the grades array and notice the value of its elements.
	fmt.Printf("\n%#v" , grades)
	// 3. Declare an array called taskDone using the ellipsis operator. The elements are of type bool. Print out taskDone.

	taskDone := [...]bool{true , false , true}
	fmt.Println("\n" , taskDone)

	// 4. Iterate over the array called cities using the classical for loop syntax and the len function and print out element by element.

	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	// 5. Iterate over grades using the range keyword and print out element by element.

	for i , v := range grades{
		fmt.Println(i , v)
	}


}