package main

import (
	"bank-test/cars"
)

func main() {
	car := cars.BMV{}
	car.StepOnGas()
	car2 := cars.Zhiga{}
	car2.StepOnGas()
	cars.Ride(car)
}


