// 열거형 3
package main

import "fmt"

func main() {
  // _ 은 skip의 의미
  const (
    _ = iota
    A
    _
    C
    D
  )

  const (
    _ = iota + 0.75 * 2
    DEFAULT
    SILVER
    _
    PLATINUM
  )

  fmt.Println("DEFAULT: ", DEFAULT, "SILVER: ",  SILVER, "GOLD: ", GOLD, "PLATINUM: ", PLATINUM)
}
