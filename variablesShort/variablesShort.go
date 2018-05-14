package main

import (
    "fmt"
)

func main() {
    intVar01 := 101
    intVar02, strVar01 := 102, "TestString"

    fmt.Printf("variables values, intVar01 %d, intVar02 %d, strVar01 %s\n", intVar01, intVar02, strVar01)
}
