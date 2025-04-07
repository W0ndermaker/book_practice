package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RUR
)

func main() {
	symbol := [...]string{USD: "1", EUR: "2", GBP: "21", RUR: "3"}
	fmt.Println(RUR, symbol[RUR])
	fmt.Println(len(symbol))
}
