/*
  hi.go
  Greet using a module
  
  Sparisoma Viridi | https://github.com/dudung
  
  20221112 Copy [1] and modify it.
  
  execute:
  go mode edit -replace github.com/dudung/greetings=../greetings
  go mod tidy
  go run .
  
  refs
  1. url https://go.dev/doc/tutorial/call-module-code [20221112].
*/

package main

import "fmt"

import "github.com/dudung/greetings"

func main() {
  // Get a greeting message and print the message.
  msg := greetings.Hello("King")
  fmt.Println(msg)
}
