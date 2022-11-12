/*
  quote.go
  Call code in an external package
  
  Sprisoma Viridi | https://github.com/dudung
  
  20221112 Copy [1] and modify it.
  
  execute:
  go mod tidy
  go run .
  
  refs
  1. url https://go.dev/doc/tutorial/getting-started [20221112].
*/

package main

import "fmt"

import "rsc.io/quote"

func main() {
  fmt.Println(quote.Glass())
  fmt.Println(quote.Go())
  fmt.Println(quote.Hello())
  fmt.Println(quote.Opt())
}
