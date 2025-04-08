package pkgloop
/*
dev a func Map(strs []string, mapFunc func(s string) string) []string
that transforms every element of slice `strs` by mapFunc and then returns a slice
*/

// BEGIN (write your solution here)
import (
  //"fmt"
)
func Map(strs []string, mapFunc func(s string) string) []string{
  lenInput := len(strs)
  resultSliceOfStrings := make([]string,0,lenInput)
  //fmt.Println("len = ", len(resultSliceOfStrings))
  for _, chad := range strs {
    //fmt.Println("Output is %s", mapFunc(chad) )
    resultSliceOfStrings = append(resultSliceOfStrings, mapFunc(chad))
  }

  return resultSliceOfStrings
}
// END