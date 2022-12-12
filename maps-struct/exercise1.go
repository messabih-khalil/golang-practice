package main

import "fmt"

type Product struct{
	productName string
	productPrice float64
}

type Customer struct{
	customerName string
}

type Store struct{
	storeProducts map[string]Product
}

type Carte struct{
	carteCustomer Customer
	carteProduct []Product
}



func main(){
	// create list of products
	product := make(map[string]Product)
	
	product["product1"] = Product{
		productName: "first product",
		productPrice: 120.0,
	}

	product["product2"] = Product{
		productName: "secend product",
		productPrice: 125.5,
	}

	// create list of customer

	names := []string{"aldin" , "tars" , "khalil"}

	carte := make(map[string]Carte)
	for _ , value := range names{

		carte[value] = Carte{
			carteCustomer: Customer{
				customerName: value,
			},
			carteProduct: []Product{product["product2"]},
		}
	}

	fmt.Println(carte)
	
	 
}