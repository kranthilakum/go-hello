package main

import "fmt"
import "rsc.io/quote"
import "lakum.in/greetings"

func main() {
  greeting := greetings.Hello("Go")
  fmt.Println(greeting)
  fmt.Println("Hello, World! This is Go time!")
  fmt.Println(quote.Go())
}

