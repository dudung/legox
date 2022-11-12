/*
  greetings.go
  A module to be called from hi.go
  
  Sparisoma Viridi | https://github.com/dudung
  
  20221112 Copy [1] and modify it.
  
  refs
  1. url https://go.dev/doc/tutorial/create-module [20221112].
*/

package greetings

import "fmt"

// Hello return a greeting from the named person.
func Hello(name string) string {
  // Return a greeting that embeds the name in a message.
  msg := fmt.Sprintf("Hi, %v. Welcome to Go world!", name)
  return msg
}
