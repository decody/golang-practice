// 데이터타입 bool 2
package main

import (
  "fmt"
)

func main() {
  // 예제1 (논리연산자)
  fmt.Println("ex1: ", true && true)
  fmt.Println("ex1: ", true && false)
  fmt.Println("ex1: ", false && false)
  fmt.Println("ex1: ", true || true)
  fmt.Println("ex1: ", true || false)
  fmt.Println("ex1: ", false || false)
  fmt.Println("ex1: ", !true)
  fmt.Println("ex1: ", !false)

  // 예제2 (비교연산자)
  num1 := 15
  num2 := 37

  fmt.Println("ex2: ", num1 < num2)
  fmt.Println("ex2: ", num1 > num2)
  fmt.Println("ex2: ", num1 >= num2)
  fmt.Println("ex2: ", num1 <= num2)
  fmt.Println("ex2: ", num1 == num2)
  fmt.Println("ex2: ", num1 != num2)
}
