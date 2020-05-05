package godash

import (
	"math"
)

// @Method: ArrayIntChunk
// @Description: 将数组（array）拆分成多个 size 长度的区块，并将这些区块组成一个新数组。 如果array 无法被分割成全部等长的区块，那么最后剩余的元素将组成一个区块
// @Param: 1.arr []int 2.size int
// @Return: [][]int
// TODO: 测试不通过（04.30）
// func ArrayChunk(in, out interface{}, size int) error {
// 	input := reflect.ValueOf(in)
// 	output := reflect.ValueOf(out)

// 	if input.Kind() == reflect.Slice || input.Kind() == reflect.Array {
// 		chunkNum := int(math.Ceil(float64(input.Len()) / float64(size)))
// 		result := reflect.MakeSlice(output.Elem().Type(), 0, chunkNum)

// 		for i := 0; i < chunkNum; i++ {
// 			if i == (chunkNum - 1) {
// 				chunkArr := reflect.MakeSlice(output.Elem().Type(), 0, 0)
// 				chunkArr = reflect.AppendSlice(chunkArr, input.Slice(i*size, input.Len()))
// 				result = reflect.AppendSlice(result, chunkArr)
// 			} else {
// 				chunkArr := reflect.MakeSlice(output.Elem().Type(), 0, 0)
// 				chunkArr = reflect.AppendSlice(chunkArr, input.Slice(i*size, i*size+size))
// 				result = reflect.AppendSlice(result, chunkArr)
// 			}
// 		}
// 		output.Elem().Set(result)

// 		return nil
// 	} else {
// 		return fmt.Errorf("ArrayChunk input type cannot be (%s)", input.Kind())
// 	}
// }

// @Method: ArrayChunk
// @Description: 将数组（array）拆分成多个 size 长度的区块，并将这些区块组成一个新数组。 如果array 无法被分割成全部等长的区块，那么最后剩余的元素将组成一个区块
// @Param: 1.arr []int 2.size int
// @Return: [][]int
func ArrayChunk(s []interface{}, size int) [][]interface{} {
	if size < 1 {
		panic("size: cannot be less than 1")
	}
	length := len(s)
	chunks := int(math.Ceil(float64(length) / float64(size)))
	var n [][]interface{}
	for i, end := 0, 0; chunks > 0; chunks-- {
		end = (i + 1) * size
		if end > length {
			end = length
		}
		n = append(n, s[i*size:end])
		i++
	}
	return n
}

// @Method: ArrayIntChunk
// @Description: 将数组（array）拆分成多个 size 长度的区块，并将这些区块组成一个新数组。 如果array 无法被分割成全部等长的区块，那么最后剩余的元素将组成一个区块
// @Param: 1.arr []int 2.size int
// @Return: [][]int
func ArrayIntChunk(input []int, size int) [][]int {
	outInterface := ArrayChunk(IntToInterface(input), size)
	var output = make([][]int, len(outInterface))
	for index, item := range outInterface {
		output[index] = ArrayInterfaceToInt(item)
	}
	return output
}

// @Method: ArrayStringChunk
// @Description: 将数组（array）拆分成多个 size 长度的区块，并将这些区块组成一个新数组。 如果array 无法被分割成全部等长的区块，那么最后剩余的元素将组成一个区块
// @Param: 1.arr []string 2.size int
// @Return: [][]string
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

// @Method: ArrayDifference
// @Description: 创建一个具有唯一array值的数组，每个值不包含在其他给定的数组中。（注：即创建一个新数组，这个数组中的值，为第一个数字（array 参数）排除了给定数组中的值。）该方法使用 SameValueZero做相等比较。结果值的顺序是由第一个数组中的顺序确定
// @Param: 1.arr []interface{} 2.arrAnother []interface{}
// @Return: []interface{}
func ArrayDifference(arr []interface{}, arrAnother []interface{}) []interface{} {
	var retArr []interface{}

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

// @Method: ArrayIntDifference
// @Description: 创建一个具有唯一array值的数组，每个值不包含在其他给定的数组中。（注：即创建一个新数组，这个数组中的值，为第一个数字（array 参数）排除了给定数组中的值。）该方法使用 SameValueZero做相等比较。结果值的顺序是由第一个数组中的顺序确定
// @Param: 1.arr []int 2.arrAnother []int
// @Return: []int
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

// @method: ArrayIntFill
// @Description: 使用 value 值来填充（替换） array，从start位置开始, 到end位置结束（但不包含end位置）
// @Param: 1.arr []int 2.val int 3.start int 4.end int
// @Return: []int
func ArrayIntFill(arr []int, val int, start int, end int) []int {
	var retArr []int
	for i := 0; i < len(arr); i++ {
		if i >= start && i < end {
			retArr = append(retArr, val)
		} else {
			retArr = append(retArr, arr[i])
		}
	}
	return retArr
}
