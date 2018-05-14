package main
import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) != 2 {
       fmt.Println("missing or extra arguments !!!", os.Args[1:])
       os.Exit(1)
    }
    fmt.Println("It's over ", os.Args[1])
}
