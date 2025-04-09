package pkgloop

import (
	"fmt"
	"slices"
	"sort"
)

func UniqueSortedUserIDs(userIDs []int64) []int64 {

	if len(userIDs) <= 1 {
		return userIDs
	}
	fmt.Println("input data >> ", userIDs)
	var noRepeatedNums []int64
	noRepeatedNums = append(noRepeatedNums, userIDs[0])

	for _, val := range userIDs {
		if !slices.Contains(noRepeatedNums, val) {
			noRepeatedNums = append(noRepeatedNums, val)
		}
	}

	sort.Slice(noRepeatedNums, func(i, j int) bool {

		fmt.Println("comaring >> ", noRepeatedNums[i], noRepeatedNums[j])
		return noRepeatedNums[i] < noRepeatedNums[j]
	})
	fmt.Println("result slice is >> ", noRepeatedNums)
	return noRepeatedNums
}
