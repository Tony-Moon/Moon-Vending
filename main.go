package main

import (
	"fmt"
	"github.com/Tony-Moon/Moon-Vending/vend"
)

func main(){
	var cewmachine = vend.Machine {
		Model:    "18239342",
		Address:  vend.NewAddress("101 Mitchell St", "Arlington", "TX", "76010"),
		Capacity: 100,
	}

	var p vend.Printer = cewmachine
	fmt.Print(p)
}