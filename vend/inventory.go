package vend

import "strconv"

// Inventory about individual rows
type Inventory struct {
	Row      string 
	Beverage string
	Stock    int
}

// This function dispenses one beverage
func (d *Inventory) Dispense() int{
	d.Stock = d.Stock - 1
	return d.Stock
}

// Print the readout of the row and its stock
func (r Inventory) String() string{
	return r.Row + " - " + r.Beverage +  " - " + strconv.Itoa(r.Stock) + " in stock.\n"
}