package godash

import (
	"reflect"
	"testing"
)

// ArrayIntChunk单元测试用例
func TestArrayIntChunk(t *testing.T) {
	cases := []struct {
		name       string
		inputArr   []int
		inputChunk int
		expect     [][]int
	}{
		{"ArrayIntChunk test 1", []int{2, 7, 11, 15, 9}, 2, [][]int{{2, 7}, {11, 15}, {9}}},
		{"ArrayIntChunk test 2", []int{3, 2, 4, 6}, 3, [][]int{{3, 2, 4}, {6}}},
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

// ArrayIntChunk基准测试用例
func BenchmarkArrayIntChunk(b *testing.B) {
	a := []int{1, 2, 3, 4, 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ArrayIntChunk(a, 3)
	}
}

// ArrayIntDifference单元测试用例
func TestArrayIntDifference(t *testing.T) {
	cases := []struct {
		name            string
		inputArr        []int
		inputArrAnother []int
		expect          []int
	}{
		{"ArrayIntDifference test 1", []int{3, 2, 1}, []int{4, 2}, []int{3, 1}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := ArrayIntDifference(c.inputArr, c.inputArrAnother)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputArr)
			}
		})
	}
}

// ArrayIntDifference基准测试用例
func BenchmarkArrayIntDifference(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5}
	arrAnother := []int{1, 3, 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ArrayIntDifference(arr, arrAnother)
	}
}
