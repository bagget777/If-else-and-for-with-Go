package main

import "fmt"

type User struct {
	name 	 string
	age 	 uint
	isActive bool
}

type Product struct {
	Title 	 string
	price 	 float64
	Quantity uint
}

type Address struct {
	City   string
	Street string
}

type Company struct {
	Name 		 string
	LegalAddress Address
}
func main() {
	user1 := User{
		name: 	  "Alex",
		age:  	  17,
		isActive: true,
	}
	fmt.Println(user1)
	ProductPhone := Product{
		Title:    "Iphon",
		price: 	  1399.9,
		Quantity: 12,
	}
	fmt.Println("Имеется продукт", ProductPhone.Title, "Его цена:", ProductPhone.price, "Кол-во на складе:", ProductPhone.Quantity)
	address := Address{
		City: 	"New-York",
		Street: "Sweet-City",
	}
	company := Company{
		Name: "ALex-GOne",
		LegalAddress: address,
	}
	fmt.Println(company)

	PrintProduct(ProductPhone)
}

func PrintProduct(product Product) {
	fmt.Println("У нас имеется", product.Title)
	fmt.Println("ЕГо цена:", product.price)
	fmt.Println("кол-во на складе:", product.Quantity)
}
