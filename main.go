package main

import (
    "fmt"
)
func printBytes(s string) {
    fmt.Printf("Bytes: ")
    for i:= 0; i< len(s); i++ {
        fmt.Printf("%x", s[i])
    }
}

func main() {
    name:= "HELLO WORLD"
    fmt.Printf("String: %s\n", name)
    printBytes(name)
}
