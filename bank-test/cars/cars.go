package cars

import "fmt"

type BMV struct {}

type Zhiga struct {}

type Cars interface {
	StepOnGas()
	StepOnBreak()
}

func (b BMV) StepOnBreak() {
	fmt.Println("i bmv, break...")
}

func (b BMV) StepOnGas() {
	fmt.Println("i bmv, drive...")
}

func (z Zhiga) StepOnBreak() {
	fmt.Println("i zhiga, break...")
}

func (z Zhiga) StepOnGas() {
	fmt.Println("i zhiga, drive...")
}

func Ride(car Cars) {
	fmt.Println("завожусь")
	fmt.Println("разогреваюсь")
	fmt.Println("Я поехал")
	fmt.Println("Нажал на тормоз")
	car.StepOnGas()
	car.StepOnBreak()
}
//изучить пакеты модули и указатели
