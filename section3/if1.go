package main

import "fmt"

func main() {
  // 제어문 (조건문)
  // if문: 반드시 Boolean 검사 -> 1, 0 (사용불가, 자동 형 변환 불가)
  // 소괄호 사용불가

  var a int = 20
  b := 20

  // 예제 1
  if a >= 15 {
    fmt.Println("15이상")
  }

  // 예제 2
  if b >= 25 {
    fmt.Println("25이상")
  }

  // 에러발생 1
  /*
  if b >= 25
  {
    fmt.Println("25이상")
  }
  */

  // 예러발생 2
  /*
  if b >=25
    fmt.Println("25이상")
  */
  // 에러발생 3
  /*
  if c := 1; c {
    fmt.Println("True")
  }
  */

  // 예제 3
  if c := 40; c >= 35 {
    fmt.Println("35이상")
  }

  //c += 2- //에러 발생
}
