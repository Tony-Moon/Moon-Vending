package main

type address struct {
	street string
	city   string
	state  string
	zip    string
}

func (a address) String() string{
	var output string
	output = a.street + "\n" + a.city + ", " + " " + a.zip + "\n"
	return output
}