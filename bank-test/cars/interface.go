package cars

import "fmt"

type BMV struct {}

func (b *BMV) StepOnGas() {
	fmt.Println("i bmv, drive...")
}

