package main

import "fmt"

// 1. Create a new struct type called book with 2 fields: price and title of type float64 and string.

type book struct{
	title string
	price float64
} 

// 3. Create a method called discount that takes a book value as receiver and decreases the price of the book by 10%.

func (b * book) discount(){
	b.price = b.price - (b.price * 0.1)
}

func main(){
	book := book{title : "Learn go for dumies" , price: 45.25}

	fmt.Println(book)
	book.discount()
	fmt.Println(book)
}