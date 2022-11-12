/*
  greetings_rand.go
  A module with error handler to be called by hi_rand.go
  
  Sparisoma Viridi | https://github.com/dudung
  
  20221112 Copy [1] and modify it.
  
  refs
  1. url https://go.dev/doc/tutorial/random-greeting [20221112].
*/

package greetings_rand

import (
  "errors"
  "fmt"
  "math/rand"
  "time"
)

func Hello(name string) (string, error) {
  if name == "" {
    return "", errors.New("Empty name")
  }
  
  msg := fmt.Sprintf(randomFormat(), name)
  return msg, nil
}


func init() {
  rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
  formats := [] string {
    "Hi, %v. Welcome to Go world!",
    "Great to see you, %v!",
    "Hail, %v! Well met!",
    "Schoen Sie, kennenzulernen, %v!",
    "Apa kabar, %v! Selamat datang di dunia Go.",
  }
  
  return formats[rand.Intn(len(formats))]
}
