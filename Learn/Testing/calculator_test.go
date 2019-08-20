package main

import (
	"testing"
)

func TestAddition(t *testing.T) {
	/*
		if addition(2) != 3 {
			t.Error("Expected output to be 3")
		}
	*/
	var addition_test = []struct {
		input    int
		expected int
	}{
		{2, 3},
		{-1, 0},
		{0, 1},
		{-5, -4},
		{99999, 100000},
	}

	for _, test := range addition_test {
		if output := addition(test.input); output != test.expected {
			t.Errorf("Test Failed: {%v} inputted, {%v} expected, recieved: {%v}", test.input, test.expected, output)
		}
	}

}

func TestSubtraction(t *testing.T) {
	if subtraction(2) != 1 {
		t.Error("Expected output to be 1")
	}
}

func TestMain(t *testing.T) {
	main()
}
