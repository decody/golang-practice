// Go 특징 1
package main

import "fmt"

func main() {
  // Go: 모호하거나 애매한 문법 금지
  // 후치(후위) 연산자 허용 (i++), 전치(전위) 연산자 허용 비허용 (++i)
  // 증감연산 반환 값 없음
  // 포인트(Pointer 허용, 연산 비허용)
  // 주석 (/, /* */) 사요

  // 예제1
  sum, i := 0, 0

  for i <= 100 {
    // sum += i++ //예외 발생
    sum += i
    i++
    // ++i 예외 발생 (전위 증감)
  }
  fmt.Println("ex1: ", sum) 
}
