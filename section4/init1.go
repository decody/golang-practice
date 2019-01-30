// go 초기화 함수 1
package main

import (
  "fmt"
)

// 초기화 함수 main보다 먼저 실행
func init() {
  fmt.Println("Init Method Start!")
}

func main() {
  // init: 패키지 로드 시에 가장 먼저 호출
  // 가장 먼저 초기화되는 작업 작성시 유용함

  fmt.Println("Main Method Start!")
}
