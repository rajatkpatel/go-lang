package main

import (
	"testing"
)

var testCase = []struct {
	input  []interface{}
	output interface{}
}{
	{[]interface{}{1, 2}, 3},
	{[]interface{}{2.0, 3.0}, 5.0},
	{[]interface{}{-1, 2}, 1},
	{[]interface{}{1.0, 2}, "Error in Addition"},
	{[]interface{}{"hello", "world"}, "helloworld"},
	{[]interface{}{"hello", 1}, "Error in Addition"},
	{[]interface{}{byte(1), 2}, "Error in Addition"},
}

func TestAddition(t *testing.T) {
	for _, testValue := range testCase {
		if result := Addition(testValue.input[0], testValue.input[1]); result != testValue.output {
			t.Errorf("%v + %v should be equal to %v", testValue.input[0], testValue.input[1], testValue.output)
		}
	}

	main()
}
