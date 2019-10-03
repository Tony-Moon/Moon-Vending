package vend

/*
Machine is a struct representng a vending machine 
It stores the model number, the address, and the capcity of items.
It has a String() method that overrides the default print behavior.
*/

type Machine struct {
	Model    string
	Address  *Address
	Capacity int
}

func (m Machine) String() string{
	return "This machine is located at: \n" + m.Address.String()
}