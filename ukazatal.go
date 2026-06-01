package main

import "fmt"

type Counter struct {
	Value int
}

type Employee struct {
	ID uint 
	Salary float64
}

type User struct {
	name *string
}

type BMV struct {
	owner string
}

type Product struct {
	title string
	price float64
}

func main() {
	n1 := 5
	result1 := InvertSign(&n1)
	fmt.Println(*result1)

	n2 := "hello"
	ResetString(&n2)

	n3, n4 := 2, 5
	Swap(&n3, &n4)

	n5 := Counter{
		Value: 1,
	}
	Increment(&n5)
	
	n6 := Employee{
		ID: 1,
		Salary: 24999.9,
	}
	GiveBonus(&n6, 5000.1)
	
	ChangeUkas()

	n7 := []int{}
	FillSlice(&n7)
	
	pn8 := ""
	n8 := User{
		name: &pn8,
	}

	n8e := CheckUser(&n8)
	fmt.Println(n8e)

	n9 := BMV{owner: "Baizak",}
	n9.ChangeOwner("asd")

	n10, n11 := NewValidProduct("Bazyk", 1234.2)
	fmt.Println(n10, n11)

}

func InvertSign(x *int) *int {
	n1 := *x * 2
	n2 := *x - n1
	return &n2
}

func ResetString(s *string) {
	*s = ""
	fmt.Println(*s)
}

func Swap(a, b *int) {
	*a, *b = *b, *a
	fmt.Println(*a, *b)
}

func Increment(c *Counter) {
	c.Value = c.Value + 1
	fmt.Println(*c)
}

func GiveBonus(emp *Employee, bonus float64) {
	emp.Salary = emp.Salary + bonus
	fmt.Println("Сотрудник", emp.ID, "Получил премию, его конечная зп =", emp.Salary)
}

func ChangeUkas() {
	x := 42
	p := &x
	pp := *p
	fmt.Println(pp)
	pp = 100
	fmt.Println(pp)
}

func FillSlice(s *[]int) {
	*s = append(*s, 1, 2, 3, 4)
	fmt.Println(*s)
}

func CheckUser(u *User) bool {
	if u.name == nil {
		return false
	} 
	return *u.name != ""
}

func (bmv *BMV) ChangeOwner(newOwner string) {
	fmt.Println(bmv.owner, "Стал")
	bmv.owner = newOwner
	fmt.Println(bmv.owner)
}

func NewValidProduct(title string, price float64) (*Product, error) {
	if title == "" {
		return nil, fmt.Errorf("Название не может быть пустым")
	}
	if price < 0 {
		return nil, fmt.Errorf("цена не может быть ниже 0")
	}

	return &Product{title: title, price: price}, nil
}
	
