package cars

import "fmt"

type Cars interface {
	StepOnGas()
}

type BMV struct {}

type Zhiga struct {}

func (b BMV) StepOnGas() {
	fmt.Println("i bmv, drive...")
}

func (z Zhiga) StepOnGas() {
	fmt.Println("i zhiga, drive...")
}

func Ride(car Cars) {
	fmt.Println("завожусь")
	fmt.Println("разогреваюсь")
	fmt.Println("Я поехал")
	car.StepOnGas()
}
//изучить пакеты модули и указатели
