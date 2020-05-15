package godash

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ArrayChunk单元测试用例
func TestArrayChunk(t *testing.T) {
	t.Run("ArrayChunk Test", func(t *testing.T) {
		input := []int{2, 7, 11, 15, 9}

		outputChunk := ArrayChunk(ArrayIntToInterface(input), 2)
		assert.Equal(t, [][]interface{}{{2, 7}, {11, 15}, {9}}, outputChunk)

	})
}

// ArrayIntChunk基准测试用例
func BenchmarkArrayChunk(b *testing.B) {
	input := []int{1, 2, 3, 4, 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ArrayChunk(ArrayIntToInterface(input), 2)
	}
}

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
	input := []int{1, 2, 3, 4, 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ArrayIntChunk(input, 3)
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

func TestArrayFind(t *testing.T) {
	t.Run("should filter elements that fail predicate", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8}
		var output int

		err := ArrayFind(input, &output, func(a int) bool {
			return a == 1 // starts with
		})
		expected := 1
		assert.NoError(t, err)
		assert.Equal(t, expected, output)
	})

	t.Run("should struct types", func(t *testing.T) {
		type person struct {
			age int
		}
		input := []person{
			{30},
			{20},
			{40},
			{10},
		}
		var output person

		err := ArrayFind(input, &output, func(p person) bool {
			return p.age > 20
		})
		expected := person{30}

		assert.NoError(t, err)
		assert.Equal(t, expected, output)
	})

	t.Run("should validate predicate's arg", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8}
		var output int

		err := ArrayFind(input, &output, func(a string) bool {
			return a == ""
		})

		assert.EqualError(t, err, "predicate function's first argument has to be the type (int) instead of (string)")
	})

	t.Run("should validate predicate's return type", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8}
		var output int

		{
			err := ArrayFind(input, &output, func(a int) int {
				return a
			})
			assert.EqualError(t, err, "predicate function should return only a (boolean) and not a (int)")
		}
		{
			err := ArrayFind(input, &output, func(int) (int, bool) {
				return 1, true
			})
			assert.EqualError(t, err, "predicate function should return only one return value - a boolean")
		}
	})

	t.Run("should validate output's type", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8}
		var output string

		err := ArrayFind(input, &output, func(a int) bool {
			return a == 0
		})

		assert.EqualError(t, err, "input slice (int) and output (string) should be of the same Type")
	})

	t.Run("should return error if element not found", func(t *testing.T) {
		in := []int{1, 2, 3}
		{
			var out int

			err := ArrayFind(in, &out, func(x int) bool { return x == 4 })

			assert.EqualError(t, err, "element not found")
		}
		{
			var out int
			err := ArrayFind(in, &out, func(x int) bool { return x == 1 })
			assert.NoError(t, err)
		}
	})

}
