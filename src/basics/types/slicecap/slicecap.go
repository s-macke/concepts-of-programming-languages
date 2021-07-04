package main

import "fmt"

func main() {
	var slice []int
	capBefore := 1
	for i := 0; i < 1e7; i++ {
		slice = append(slice, i)
		capAfter := cap(slice)
		if capAfter != capBefore {
			percent := 100 * float64(capAfter) / float64(capBefore)
			fmt.Printf("Cap increased from %d -> %d (by %f%%) at len %d\n",
				capBefore, capAfter, percent, len(slice))
			capBefore = capAfter
		}
	}
}
