package main

import (
	"fmt"
	"math"
)

func ArrayChunk(arr []int, chunk int) [3][3]int {
	perChunkLen := int(math.Ceil(float64(len(arr)) / float64(chunk)))
	var retArr [3][3]int

	for index, item := range arr {
		chunkIndex := int(math.Floor(float64(index) / float64(perChunkLen)))
		var chunkItemIndex int
		if chunkIndex != 0 {
			chunkItemIndex = index % perChunkLen
		} else {
			chunkItemIndex = index
		}
		fmt.Printf("chunkIndex: %d chunkItemIndex: %d \n", chunkIndex, chunkItemIndex)
		retArr[chunkIndex][chunkItemIndex] = item
	}
	return retArr
}

func main() {
	b := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(ArrayChunk(b, 3))
	ArrayChunk(b, 3)
}
