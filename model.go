package main

//go:generate msgp
//go:generate msgpackgen

type Person struct {
	Name    string `msg:"name"`
	Address string `msg:"address"`
	Age     int    `msg:"age"`
}
