package main

import "fmt"

var (
	name		string 	= "John"
	age			int 	= 32
	location	string 	= "New York"
)

func main() {
	fmt.Println(name, age, location)
	fmt.Println(StatusAccepted)
	fmt.Println(big)
}

const (
	StatusOK		= 200
	StatusNotFound	= 404
	StatusAccepted	= 202
	big = 1 << 2 // shift left 2 bits
	// 00000001 << 2 = 00000100 = 0 * 2^0 + 0 * 2^1 + 1 * 2^2 + 0 * 2^3 = 4
)