package main

import (
    "fmt"
    "hello/stringutils"
)

func main() {
    fmt.Println("Hello, world!")
    fmt.Println(stringutils.Reverse("Hello, world!"))

    a, b := stringutils.Swap("Hello,", "world")
    fmt.Println(a, b)
    fmt.Println(stringutils.Upper("Hello, world"))
}
