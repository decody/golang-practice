// go 초기화 함수 3
// init 함수를 이용하여 다른 패키지 내의 init 함수가 있으면 그것도 실행됨
// go 라이브러리들이 이런 의존성(dependancy)를 가지면서 설계
// https://golang.org/doc/effective_go.html#initialization
package main

import (
  "fmt"
  "section4/lib"
)

var num int32

// 변수 초기화
func init() {
  num = 30
  fmt.Println("Main init start!")
}

func main() {
  fmt.Println("10보다 큰 수?: ", lib.CheckNum(num) )
}
