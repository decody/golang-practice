// Switch 2
package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  // 예제1
  rand.Seed(time.Now().UnixNano())  // 난수 발생

  switch i := rand.Intn(100); {
  case i >= 50 && i < 100:
    fmt.Println("i -> ", i, "50이상 100미만")
  case i >= 25 && i < 50:
    fmt.Println("i -> ", i, "25이상 100미만")
  default:
    fmt.Println("i -> ", i, "기본 값")
  }
}