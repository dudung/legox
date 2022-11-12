/*
  ping.go
  Call code from  module from a repository (repo.io)
  and local folder (lofo.io) placed in pkg folder
  
  Sprisoma Viridi | https://github.com/dudung
  
  execute:
  go mod edit -replace github.com/dudung/sego/pkg/lofo=../../pkg/lofo
  go mod tidy
  go run .
  
  
  20221112 Create it.
*/
package main

import (
  "fmt"
  
  "github.com/dudung/sego/pkg/repo"
  "github.com/dudung/sego/pkg/lofo"
)

func main() {
  rmsg := repo.Ping()
  fmt.Println(rmsg)
  
  lmsg := lofo.Ping()
  fmt.Println(lmsg)
}
