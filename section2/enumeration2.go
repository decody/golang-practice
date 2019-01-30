// 열거형 2
package main

import "fmt"

func main() {
  // iota는 0부터 시작해서 x 10씩 증가
  const (
    A = iota * 10
    B
    C
  )

  fmt.Println(A, B, C)

  // iota는 0부터 시작해서 + 1씩 증가
  const (
    Jan = iota + 1
    Feb
    Mar
    Apr
    May
    Jun
  )

  fmt.Println(Jan)
  fmt.Println(Feb)
  fmt.Println(Mar)
  fmt.Println(Apr)
  fmt.Println(May)
  fmt.Println(Jun)
}
