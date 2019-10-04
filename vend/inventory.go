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

// This function restocks the beverage
func (r *Inventory) Restock(amount int) int{
	r.Stock = r.Stock + amount
	return r.Stock
}

// Print the readout of the row and its stock
func (p Inventory) String() string{
	return p.Row + " - " + p.Beverage +  " - " + strconv.Itoa(p.Stock) + " in stock.\n"
}