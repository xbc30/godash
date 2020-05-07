package godash

import (
	"fmt"
	"math"
	"reflect"
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

/*
 *@Method: ArrayChunk
 *@Description: 将数组（array）拆分成多个 size 长度的区块，并将这些区块组成一个新数组。 如果array 无法被分割成全部等长的区块，那么最后剩余的元素将组成一个区块
 *@Param: 1.arr []interface{} 2.size int
 *@Return: [][]interface{}
 */
func ArrayChunk(input []interface{}, size int) [][]interface{} {
	if size < 1 {
		panic("size: cannot be less than 1")
	}
	length := len(input)
	chunks := int(math.Ceil(float64(length) / float64(size)))
	output := make([][]interface{}, 0)
	for i, end := 0, 0; chunks > 0; chunks-- {
		end = (i + 1) * size
		if end > length {
			end = length
		}
		output = append(output, input[i*size:end])
		i++
	}
	return output
}

/*
 *@Method: ArrayIntChunk
 *@Description: 将数组（array）拆分成多个 size 长度的区块，并将这些区块组成一个新数组。 如果array 无法被分割成全部等长的区块，那么最后剩余的元素将组成一个区块
 *@Param: 1.arr []int 2.size int
 *@Return: [][]int
 */
func ArrayIntChunk(input []int, size int) [][]int {
	outInterface := ArrayChunk(ArrayIntToInterface(input), size)
	var output = make([][]int, len(outInterface))
	for index, item := range outInterface {
		output[index] = ArrayInterfaceToInt(item)
	}
	return output
}

/*
 *@Method: ArrayStringChunk
 *@Description: 将数组（array）拆分成多个 size 长度的区块，并将这些区块组成一个新数组。 如果array 无法被分割成全部等长的区块，那么最后剩余的元素将组成一个区块
 *@Param: 1.input []string 2.size int
 *@Return: [][]string
 */
func ArrayStringChunk(input []string, size int) [][]string {
	outInterface := ArrayChunk(ArrayStringToInterface(input), size)
	var output = make([][]string, len(outInterface))
	for index, item := range outInterface {
		output[index] = ArrayInterfaceToString(item)
	}
	return output
}

/*
 *@Method: ArrayDifference
 *@Description: 创建一个具有唯一array值的数组，每个值不包含在其他给定的数组中。（注：即创建一个新数组，这个数组中的值，为第一个数字（array 参数）排除了给定数组中的值。）该方法使用 SameValueZero做相等比较。结果值的顺序是由第一个数组中的顺序确定
 *@Param: 1.arr []interface{} 2.arrAnother []interface{}
 *@Return: []interface{}
 */
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

/*
 *@Method: ArrayIntDifference
 *@Description: 创建一个具有唯一array值的数组，每个值不包含在其他给定的数组中。（注：即创建一个新数组，这个数组中的值，为第一个数字（array 参数）排除了给定数组中的值。）该方法使用 SameValueZero做相等比较。结果值的顺序是由第一个数组中的顺序确定
 *@Param: 1.arr []int 2.arrAnother []int
 *@Return: []int
 */
func ArrayIntDifference(arr []int, arrAnother []int) []int {
	outInterface := ArrayDifference(ArrayIntToInterface(arr), ArrayIntToInterface(arrAnother))
	return ArrayInterfaceToInt(outInterface)
}

/*
 *@Method: ArrayStringDifference
 *@Description: 创建一个具有唯一array值的数组，每个值不包含在其他给定的数组中。（注：即创建一个新数组，这个数组中的值，为第一个数字（array 参数）排除了给定数组中的值。）该方法使用 SameValueZero做相等比较。结果值的顺序是由第一个数组中的顺序确定
 *@Param: 1.arr []string 2.arrAnother []string
 *@Return: []string
 */
func ArrayStringDifference(arr []string, arrAnother []string) []string {
	outInterface := ArrayDifference(ArrayStringToInterface(arr), ArrayStringToInterface(arrAnother))
	return ArrayInterfaceToString(outInterface)
}

/*
 *@Method: ArrayIntFill
 *@Description: 使用 value 值来填充（替换） array，从start位置开始, 到end位置结束（但不包含end位置）
 *@Param: 1.arr []int 2.val int 3.start int 4.end int
 *@Return: []int
 */
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

/*
 *@Method: ArrayFindIndex
 *@Description: 该方法类似_.find，区别是该方法返回第一个通过 predicate 判断为真值的元素的索引值（index），而不是元素本身。
 *@Param: 1.
 *@Return:
 */
func ArrayFindIndex(in, out, predicateFn interface{}) error {
	input := reflect.ValueOf(in)
	output := reflect.ValueOf(out)
	inputTypeElem := input.Type().Elem()
	if inputTypeElem != output.Elem().Type() {
		return fmt.Errorf("input slice (%s) and output (%s) should be of the same Type", inputTypeElem, output.Elem().Type())
	}

	predicate := reflect.ValueOf(predicateFn)
	if predicate.Type().NumOut() != 1 {
		return fmt.Errorf("predicate function should return only one return value - a boolean")
	}
	if predicateType := predicate.Type().Out(0).Kind(); predicateType != reflect.Bool {
		return fmt.Errorf("predicate function should return only a (boolean) and not a (%s)", predicateType)
	}
	if input.Kind() == reflect.Slice {
		{
			if inputTypeElem.Kind() != predicate.Type().In(0).Kind() {
				return fmt.Errorf(
					"predicate function's first argument has to be the type (%s) instead of (%s)",
					inputTypeElem,
					predicate.Type().In(0),
				)
			}
		}
		for i := 0; i < input.Len(); i++ {
			arg := input.Index(i)

			returnValues := predicate.Call([]reflect.Value{arg})
			predicatePassed := returnValues[0].Bool()

			if predicatePassed {
				output.Elem().Set(reflect.ValueOf(i))
				return nil
			}
		}
		return fmt.Errorf("element not found")
	}
	return fmt.Errorf("not implemented")
}

/*
 *@Method: ArrayFind
 *@Description: 该方法查找符合筛选条件的值
 *@Param: 1.
 *@Return:
 */
func ArrayFind(in, out, predicateFn interface{}) error {
	input := reflect.ValueOf(in)
	output := reflect.ValueOf(out)
	inputTypeElem := input.Type().Elem()
	if inputTypeElem != output.Elem().Type() {
		return fmt.Errorf("input slice (%s) and output (%s) should be of the same Type", inputTypeElem, output.Elem().Type())
	}

	predicate := reflect.ValueOf(predicateFn)
	if predicate.Type().NumOut() != 1 {
		return fmt.Errorf("predicate function should return only one return value - a boolean")
	}
	if predicateType := predicate.Type().Out(0).Kind(); predicateType != reflect.Bool {
		return fmt.Errorf("predicate function should return only a (boolean) and not a (%s)", predicateType)
	}
	if input.Kind() == reflect.Slice {
		{
			if inputTypeElem.Kind() != predicate.Type().In(0).Kind() {
				return fmt.Errorf(
					"predicate function's first argument has to be the type (%s) instead of (%s)",
					inputTypeElem,
					predicate.Type().In(0),
				)
			}
		}
		for i := 0; i < input.Len(); i++ {
			arg := input.Index(i)

			returnValues := predicate.Call([]reflect.Value{arg})
			predicatePassed := returnValues[0].Bool()

			if predicatePassed {
				output.Elem().Set(arg)
				return nil
			}
		}
		return fmt.Errorf("element not found")
	}
	return fmt.Errorf("not implemented")
}
