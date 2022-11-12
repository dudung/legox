/*
  greetings2.go
  A module with error handler to be called by hi2.go
  
  Sparisoma Viridi | https://github.com/dudung
  
  20221112 Copy [1] and modifyt it.
  
  refs
  1. url https://go.dev/doc/tutorial/handle-errors [20221112].
*/

package greetings2

import (
  "errors"
  "fmt"
)

func Hello(name string) (string, error) {
  if name == "" {
    return "", errors.New("Empty name")
  }
  
  msg := fmt.Sprintf("Hi, %v. Welcome to Go!", name)
  return msg, nil
}
