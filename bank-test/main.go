package main

import (
	"bank-test/cars"
)

func main() {
	bmv_car := cars.BMV{}
	zhiga_car := cars.Zhiga{}

	cars.Ride(bmv_car)
	cars.Ride(zhiga_car)
}


