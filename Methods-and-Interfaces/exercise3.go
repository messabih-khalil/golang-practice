package main

import "fmt"


type vehicle interface {
    License() string
    Name() string
}

type car struct {
    licenceNo string
    brand     string
}


func (c car) License(){
	fmt.Println("car licence number : " , c.licenceNo)
}

func (c car) Name() {
	fmt.Println("car brand : " , c.brand)
}

func main()  {
	// 
	car := car{"25548f7" , "Bmw"}

	car.Name()
	car.License()
}