package main

import (
    "fmt"
)

func testFunc() int {
    return 103
}

func main() {
    var intVar01, intVar02, intVar03 int = 101, 102, testFunc()
    var strVar01 string = "TestString"

    fmt.Printf("variables values, intVar01 %d, intVar02 %d, intVar03 %d, strVar01 %s\n", intVar01, intVar02, intVar03, strVar01)
}
