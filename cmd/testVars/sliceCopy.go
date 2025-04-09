package testvars

import "fmt"

func PrintCopiedSlice() {
    nums := []int{1, 2, 3, 4, 5}

    numsCp := make([]int, 1)

    copy(numsCp, nums)

    fmt.Println(numsCp) // [1 2 3 4 5]
}