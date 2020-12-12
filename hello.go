package main

import (
    "fmt"
    "rsc.io/quote"

    "greetings"
)

func main() {
    fmt.Println(greetings.Hello("World"))
    fmt.Println(quote.Go())
}

