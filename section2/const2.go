// 상수 2
package main

import "fmt"

func main() {
  // 여러 개 선언
  const a, b int = 10, 99
  const c, d, e = true, 0.84, "test"
  const (
    x, y int16  = 50, 90
    i, k        = "Data", 7776
  )

  fmt.Println(a, b, c, d, e)
  fmt.Println(x, y, i, k)
}
