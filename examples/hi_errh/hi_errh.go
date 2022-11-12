/*
  hi_errh.go
  Call code in a module (greetings_errh.go) from local folder (/internal)
  
  Sparisoma Viridi | https://github.com/dudung
  
  20221112 Copy [1] and modify it.
  
  execute:
  go mod edit -replace github.com/dudung/sego/internal/greetings2=../../internal/greetings2
  go mod tidy
  go run .
  
  refs
  1. url https://go.dev/doc/tutorial/handle-errors [20221112].
*/
package main

import (
  "fmt"
  "log"
  
  "github.com/dudung/sego/internal/greetings_errh"
)

func main() {
  log.SetPrefix("greetings: ")
  log.SetFlags(0)
  
  msg, err := greetings_errh.Hello("Human")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(msg)
  fmt.Println()
  
  msg, err = greetings_errh.Hello("")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(msg)
}
