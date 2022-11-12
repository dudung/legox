package main

import "fmt"
  
import  "github.com/dudung/sego/pkg/repo"

func main() {
  msg := repo.Ping()
  fmt.Println(msg)
}
