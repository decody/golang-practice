// for 2
package main

import "fmt"

func main() {
  // 예제1
  sum1 := 0
  for i := 0; i <= 100; i++ {
    sum1 += 1   // sum = sum + 1
  }
  fmt.Println("ex1: ", sum1)

  // 예제2
  sum2, i := 0, 0

  for i <= 100 {
    sum2 += i
    i++
    // j += i++  // Go에서 후치연산은 반환 값 없음
  }
  fmt.Println("ex2: ", sum2)

  // 예제3
  sum3, i := 0, 0
  for { // while 향태와 비슷
    if i >  100 {
      break
    }
    sum3 += i
    i++
  }
  fmt.Println("ex3: ", sum3)

  // 예제4
  for i, j := 0, 0; i <= 10; i, j = i+1, j+10 {
    fmt.Println("ex4: ", i, j)
  }

  // 에러 발생, 후치연산은 리턴 값으로 사용 못함
  /*
  for i, j := 0, 0; i <= 10; i++, j += 10 {
      fmt.Println("ex4: ", i, j)
  }
  */
}
