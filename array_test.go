package godash

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ArrayChunk单元测试用例
// func TestArrayChunk(t *testing.T) {
// 	t.Run("support int types", func(t *testing.T) {
// 		in := []int{2, 7, 11, 15, 9}
// 		out := make([]interface{}, 0)
// 		chunk := 2

// 		err := ArrayChunk(in, &out, chunk)

// 		expected := [][]int{{2, 7}, {11, 15}, {9}}
// 		assert.NoError(t, err)
// 		assert.Equal(t, expected, out)
// 	})
// }
func TestArrayChunk(t *testing.T) {
	t.Run("ArrayChunk Test", func(t *testing.T) {
		in := []int{2, 7, 11, 15, 9}

		tArrayChunk := ArrayChunk(IntToInterface(in), 2)
		assert.Equal(t, [][]interface{}{{2, 7}, {11, 15}, {9}}, tArrayChunk)
	})
}

// func TestArrayChunk(t *testing.T) {
// 	cases := []struct {
// 		name       string
// 		inputArr   []int
// 		inputChunk int
// 		expect     [][]int
// 	}{
// 		{"ArrayChunk test 1", []int{2, 7, 11, 15, 9}, 2, [][]int{{2, 7}, {11, 15}, {9}}},
// 		{"ArrayChunk test 2", []int{3, 2, 4, 6}, 3, [][]int{{3, 2, 4}, {6}}},
// 	}

// 	//	开始测试

// 	for _, c := range cases {
// 		t.Run(c.name, func(t *testing.T) {
// 			var inputArray = make([]interface{}, len(c.inputArr))
// 			for index, item := range c.inputArr {
// 				inputArray[index] = item
// 			}
// 			ret := ArrayChunk(inputArray, c.inputChunk)
// 			if !reflect.DeepEqual(ret, c.expect) {
// 				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
// 					c.expect, ret, c.inputArr)
// 			}
// 		})
// 	}
// }

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

// ArrayIntFill单元测试用例
func TestArrayIntFill(t *testing.T) {
	cases := []struct {
		name       string
		inputArr   []int
		inputVal   int
		inputStart int
		inputEnd   int
		expect     []int
	}{
		{"ArrayIntFill test 1", []int{1, 2, 3}, 4, 1, 2, []int{1, 4, 3}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := ArrayIntFill(c.inputArr, c.inputVal, c.inputStart, c.inputEnd)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputArr)
			}
		})
	}
}

// ArrayIntFill基准测试用例
func BenchmarkArrayIntFill(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5}
	val := 4
	start := 1
	end := 2

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ArrayIntFill(arr, val, start, end)
	}
}
