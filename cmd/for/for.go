package pkgloop

// BEGIN (write your solution here)
import (
  "fmt"
  "strings"
)
func MapSimple(strs []string) []string{
  for _, chad := range strs {
    fmt.Println("input is ", strings.ToUpper(chad))
  }

  fmt.Println("_________________")
  return nil
}
// END