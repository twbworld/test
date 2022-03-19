package main

import (
    "fmt"
)

type Aa struct {
    a int
    b string
}

func f1(a string, b int) (string, int){
    return a + "return", b+1
}

func main() {
    fmt.Println("===========\n")

    ab := Aa{a: 1, b: "c"}

    ab.a += 2

    fmt.Println("===========\n")
    ac, ad := f1(ab.b, ab.a)

    fmt.Println(ac, ad)

    fmt.Println("===========\n")


}
