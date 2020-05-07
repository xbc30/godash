package godash

import (
	"strconv"
	"strings"
)

/*
 *@Method: InterfaceToString
 *@Description: 类型转换interface{} => string
 *@Param: interface{}
 *@Return: string
 */
func InterfaceToString(input interface{}) string {
	str, _ := input.(string)
	return str
}

/*
 *@Method: InterfaceToBool
 *@Description: 类型转换interface{} => bool
 *@Param: interface{}
 *@Return: bool
 */
func InterfaceToBool(input interface{}) bool {
	strBool := InterfaceToString(input)
	return strings.ToLower(strBool) == "true"
}

/*
 *@Method: InterfaceToInt
 *@Description: 类型转换interface{} => int
 *@Param: interface{}
 *@Return: int
 */
func InterfaceToInt(input interface{}) int {
	strNumber := InterfaceToString(input)
	number, _ := strconv.Atoi(strNumber)
	return number
}

/*
 *@Method: ArrayInterfaceToInt
 *@Description: 类型转换[]interface{} => []int
 *@Param: []interface{}
 *@Return: []int
 */
func ArrayInterfaceToInt(input []interface{}) []int {
	var output = make([]int, len(input))
	for index, item := range input {
		item, _ := item.(int)
		output[index] = item
	}
	return output
}

/*
 *@Method: ArrayInterfaceToString
 *@Description: 类型转换[]interface{} => []string
 *@Param: []interface{}
 *@Return: []string
 */
func ArrayInterfaceToString(input []interface{}) []string {
	var output = make([]string, len(input))
	for index, item := range input {
		item, _ := item.(string)
		output[index] = item
	}
	return output
}

/*
 *@Method: ArrayIntToInterface
 *@Description: 类型转换[]int => []interface{}
 *@Param: []int
 *@Return: []interface{}
 */
func ArrayIntToInterface(input []int) (ret []interface{}) {
	filter := make(map[interface{}]int, len(input))
	ret = make([]interface{}, 0, len(input))
	if len(input) == 0 {
		return
	}
	for _, id := range input {
		if _, ok := filter[id]; ok {
			continue
		} else {
			filter[id] = 0
		}
		ret = append(ret, id)

	}
	return
}

/*
 *@Method: ArrayStringToInterface
 *@Description: 类型转换[]string => []interface{}
 *@Param: []string
 *@Return: []interface{}
 */
func ArrayStringToInterface(input []string) (ret []interface{}) {
	filter := make(map[interface{}]int, len(input))
	ret = make([]interface{}, 0, len(input))
	if len(input) == 0 {
		return
	}
	for _, id := range input {
		if _, ok := filter[id]; ok {
			continue
		} else {
			filter[id] = 0
		}
		ret = append(ret, id)
	}
	return
}
