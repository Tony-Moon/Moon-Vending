package main

/*
This is a struct representng a vending machine 
It stores the model number, the address, and the capcity of items.
It has a String() method that overrides the default print behavior.
*/

type machine struct {
	model    string
	address  address
	capacity int
}

func (m machine) String() string{
	return "This machine is located at: \n" + m.address.String()
}