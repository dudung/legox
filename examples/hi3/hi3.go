/*
  hi3.go
  Call code in a module (greetings2.go) from local folder (/internal)
  
  Sparisoma Viridi | https://github.com/dudung
  
  20221112 Copy [1] and modifyt it.
  
  execute:
  go mod edit -replace github.com/dudung/sego/internal/greetings3=../../internal/greetings3
  go mod tidy
  go run .
  
  refs
  1. url https://go.dev/doc/tutorial/random-greeting [20221112].
*/
package main

import (
  "fmt"
  "log"
  
  "github.com/dudung/sego/internal/greetings3"
)

func main() {
  log.SetPrefix("greetings: ")
  log.SetFlags(0)
  
  msg, err := greetings3.Hello("Human")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(msg)
  fmt.Println()
  
  msg, err = greetings3.Hello("")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(msg)
}
