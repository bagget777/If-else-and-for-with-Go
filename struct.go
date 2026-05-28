package main

import (
	"fmt"
	"errors"
)

type User struct {
	Name 	 string
	Age 	 uint
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

type BankAccount struct {
	Owner string
	Balance float64
}

type Point struct {
	X int
	Y int
}

func main() {
	user1 := User{
		Name: 	  "Alex",
		Age:  	  17,
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
		Name: "		  ALex-GOne",
		LegalAddress: address,
	}
	fmt.Println(company)

	PrintProduct(ProductPhone)
	ProductApple := Product{
		Title:	  "MacBook m1",
		price: 	  2599.9,
		Quantity: 12,
	}
	result := ApplyDiscount(&ProductApple, 399.9)
	fmt.Println(result)

	mew_user := NewUser("boy", 23)
	fmt.Println(mew_user)
		
	status := user1.GetStatus()
	fmt.Println(status)

	status2 := ProductPhone.Buy(6)
	fmt.Println(status2)

	bankUser := BankAccount{
		Owner: "Bazyk",
		Balance: 20.5,
	}

	up_balance := bankUser.Deposit(123.21)
	down_balance := bankUser.Withdraw(12)
	fmt.Println(up_balance)
	fmt.Println(down_balance)

	p1 := Point{3, 4}
	p2 := Point{3, 4}
	results := ComparePoints(p1, p2)
	fmt.Println(results)
	
}

func ComparePoints(p1, p2 Point) bool {
	return p1 == p2 
}

func PrintProduct(p Product) {
	fmt.Println("У нас имеется", p.Title)
	fmt.Println("ЕГо цена:", p.price)
	fmt.Println("кол-во на складе:", p.Quantity)
}

func ApplyDiscount(p *Product, discount float64) *Product {
	p.price = p.price - discount
	return p
}

func NewUser(name string, age uint) *User {
	new_user := User{
		Name: name,
		Age: age,
		isActive: true,
	}
	return &new_user
}

func (u User) GetStatus() string {
	if u.isActive == true {
		return "User in network"
	} else {
		return "User not in network"
	}
}

func (p *Product) Buy(amount uint) string {
	if p.Quantity < amount {
		return "недостаточно кол-во товара"
	}
	p.Quantity = p.Quantity - amount
	return fmt.Sprintf("Покупка успешна! Осталось: %d", p.Quantity)
}

func (b BankAccount) Deposit(amount float64) string {
	b.Balance = b.Balance + amount 
	return fmt.Sprintf("Ваш баланс: %d", b.Balance)
}

func (b BankAccount) Withdraw(amount float64) error {
	if b.Balance < amount {
		return errors.New("Недостаточно баланса")
	} else {
		b.Balance = b.Balance - amount
		return nil
	}
}


