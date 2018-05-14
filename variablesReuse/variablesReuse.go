package main

import (
	"fmt"
)

func main() {
	power := 10000

	fmt.Printf("default power is %d\n", power)

	// this is compilation error
	// ./variablesReuse.go:13:8: no new variables on left side of :=
	// power := 1200

	// there should at least 1 new variable
	name, power := "Test", 9000
	fmt.Printf("%s's power is over %d\n", name, power)
}
