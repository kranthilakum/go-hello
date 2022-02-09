package main

import (
  "fmt"
  "log"
  "rsc.io/quote"
  "lakum.in/greetings"
)

func main() {
  // Set the Logger entry prefix.
  log.SetPrefix("greetings: ")

  // Set the flag to disable printing
  // the time, source file, and line number.
  log.SetFlags(0)

  greeting, error := greetings.Hello("")
  // If an error was returned, print it to the console
  // and exit the program.
  if error != nil {
    log.Fatal(error)
  }

  fmt.Println(greeting)
  fmt.Println("Hello, World! This is Go time!")
  fmt.Println(quote.Go())
}

