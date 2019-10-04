package main

import (
	"fmt"
	"github.com/Tony-Moon/Moon-Vending/vend"
)

func main(){
	// Build a new vending machine for the CEWF building
	var cewfmachine = vend.Machine {
		Model:    "18239342",
		Address:  vend.NewAddress("100 W Mitchell St", "Arlington", "TX", "76010"),
		Capacity: 100,
	}

	var p vend.Printer = cewfmachine
	fmt.Print(p)

	fmt.Println(" ")

	// Great! Now I can start adding stock to the vending machine. 
	var a1 = vend.Inventory {
		Row:      "A1",
		Beverage: "Coke",
		Stock:    5,
	}
	var a1p vend.Printer = a1
	fmt.Print(a1p)
	
	a1.Dispense()
	a1p = a1
	fmt.Print(a1p)

	a1.Restock(5)
	a1p = a1
	fmt.Print(a1p)

	// Row A1 down, but I should consider finding a way to generate the number of rows and filling them automatically

}