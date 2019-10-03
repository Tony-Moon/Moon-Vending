package main

import "fmt"

func main(){
	var cewmachine = machine{
		model:    "182",
		address:  address {
			street : "101 Mitchell St",
			city:    "Arlington",
			state:   "TX",
			zip:     "76010",
		},
		capacity: 100,
	}

	var p printer = cewmachine
	fmt.Print(p)
}