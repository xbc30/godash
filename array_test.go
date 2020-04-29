package godash

import (
	"reflect"
	"testing"
)

func TestArrayIntChunk(t *testing.T) {
	//	测试用例
	cases := []struct {
		name       string
		inputArr   []int
		inputChunk int
		expect     [][]int
	}{
		{"ArrayIntChunk test 1", []int{2, 7, 11, 15, 9}, 3, [][]int{{2, 7}, {11, 15}, {9}}},
		{"ArrayIntChunk test 2", []int{3, 2, 4, 6}, 2, [][]int{{3, 2}, {4, 6}}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := ArrayIntChunk(c.inputArr, c.inputChunk)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputArr)
			}
		})
	}
}
