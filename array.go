package godash

import (
	"math"
)

// @Method: ArrayIntChunk
// @Description: 将数组（array）拆分成多个 size 长度的区块，并将这些区块组成一个新数组。 如果array 无法被分割成全部等长的区块，那么最后剩余的元素将组成一个区块
// @param: 1.arr []int 2.size int
// @return: [][]int
func ArrayIntChunk(arr []int, size int) [][]int {
	chunkNum := int(math.Ceil(float64(len(arr)) / float64(size)))
	var retArr [][]int

	for i := 0; i < chunkNum; i++ {
		if i == (chunkNum - 1) {
			retArr = append(retArr, arr[i*size:len(arr)])
		} else {
			retArr = append(retArr, arr[i*size:i*size+size])
		}
	}
	return retArr
}

func ArrayStringChunk(arr []string, size int) [][]string {
	chunkNum := int(math.Ceil(float64(len(arr)) / float64(size)))
	var retArr [][]string

	for i := 0; i < chunkNum; i++ {
		if i == (chunkNum - 1) {
			retArr = append(retArr, arr[i*size:len(arr)])
		} else {
			retArr = append(retArr, arr[i*size:i*size+size])
		}
	}
	return retArr
}

// @Method: ArrayIntDifference
// @Description: 创建一个具有唯一array值的数组，每个值不包含在其他给定的数组中。（注：即创建一个新数组，这个数组中的值，为第一个数字（array 参数）排除了给定数组中的值。）该方法使用 SameValueZero做相等比较。结果值的顺序是由第一个数组中的顺序确定
// @param: 1.arr []int 2.arrAnother []int
// @return: []int
func ArrayIntDifference(arr []int, arrAnother []int) []int {
	var retArr []int

	for i := 0; i < len(arr); i++ {
		dif := arr[i]
		isDif := false
		for j := 0; j < len(arrAnother); j++ {
			if dif == arrAnother[j] {
				isDif = true
			}
		}
		if !isDif {
			retArr = append(retArr, dif)
		}
	}
	return retArr
}
