package godash

import (
	"reflect"
	"testing"
)

func TestStringCapitalize(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect string
	}{
		{"StringCapitalize", "STRING", "String"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := StringCapitalize(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.input)
			}
		})
	}
}

func TestStringEndsWith(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		suffix string
		end    int
		expect bool
	}{
		{"StringEndsWith", "hello", "hell", 1, true},
		{"StringEndsWith", "hello", "l", 1, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := StringEndsWith(c.input, c.suffix, c.end)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.input)
			}
		})
	}
}
