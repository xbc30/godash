package godash

import (
	"strconv"
	"strings"
)

func InterfaceToString(src interface{}) string {
	str, _ := src.(string)
	return str
}

func InterfaceToBool(src interface{}) bool {
	strBool := InterfaceToString(src)
	return strings.ToLower(strBool) == "true"
}

func InterfaceToInt(src interface{}) int {
	strNumber := InterfaceToString(src)
	number, _ := strconv.Atoi(strNumber)
	return number
}

func ArrayInterfaceToInt(src []interface{}) []int {
	var output = make([]int, len(src))
	for index, item := range src {
		item, _ := item.(int)
		output[index] = item
	}
	return output
}

//IntToInterface  数据转换
func IntToInterface(ids []int) (ret []interface{}) {
	filter := make(map[interface{}]int, len(ids))
	ret = make([]interface{}, 0, len(ids))
	if len(ids) == 0 {
		return
	}
	for _, id := range ids {
		if _, ok := filter[id]; ok {
			continue
		} else {
			filter[id] = 0
		}
		ret = append(ret, id)

	}
	return
}

//StringsToInterface  数据转换
func StringsToInterface(ids []string) (ret []interface{}) {
	filter := make(map[interface{}]int, len(ids))
	ret = make([]interface{}, 0, len(ids))
	if len(ids) == 0 {
		return
	}
	for _, id := range ids {
		if _, ok := filter[id]; ok {
			continue
		} else {
			filter[id] = 0
		}
		ret = append(ret, id)

	}
	return
}
