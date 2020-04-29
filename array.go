package godash

import (
	"math"
)

func ArrayIntChunk(arr []int, chunk int) [][]int {
	perChunkLen := int(math.Ceil(float64(len(arr)) / float64(chunk)))
	var retArr [][]int

	for i := 0; i < chunk; i++ {
		if i == (chunk - 1) {
			retArr = append(retArr, arr[i*perChunkLen:len(arr)])
		} else {
			retArr = append(retArr, arr[i*perChunkLen:i*perChunkLen+perChunkLen])
		}
	}
	return retArr
}

func ArrayStringChunk(arr []string, chunk int) [][]string {
	perChunkLen := int(math.Ceil(float64(len(arr)) / float64(chunk)))
	var retArr [][]string

	for i := 0; i < chunk; i++ {
		if i == (chunk - 1) {
			retArr = append(retArr, arr[i*perChunkLen:len(arr)])
		} else {
			retArr = append(retArr, arr[i*perChunkLen:i*perChunkLen+perChunkLen])
		}
	}
	return retArr
}
