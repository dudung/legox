/*
  hi_rand.go
  Call code in a module (greetings_rand.go) from local folder (/internal)
  
  Sparisoma Viridi | https://github.com/dudung
  
  20221112 Copy [1] and modify it.
  
  execute:
  go mod edit -replace github.com/dudung/sego/internal/greetings_rand=../../internal/greetings_rand
  go mod tidy
  go run .
  
  refs
  1. url https://go.dev/doc/tutorial/random-greeting [20221112].
*/
package main

import (
  "fmt"
  "log"
  
  "github.com/dudung/sego/internal/greetings_rand"
)

func main() {
  log.SetPrefix("greetings: ")
  log.SetFlags(0)
  
  msg, err := greetings_rand.Hello("Human")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(msg)
  fmt.Println()
  
  msg, err = greetings_rand.Hello("")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(msg)
}
