package calculator_test

import (
	"calculator"
	"testing"
)

type testCase struct {
	a, b float64
	want float64
	name string
}

func TestAdd(t *testing.T) {
	testCases := []testCase{
		{a: 0, b: 1, want: 1, name: "two low positive numbers which sum to a positive number"},
		{a: 10, b: 10, want: 20, name: "two positive numbers which sum to a positive number"},
		{a: 0, b: 0, want: 0, name: "two zeros which sum to a zero"},
		{a: 2, b: -1, want: 0, name: "one positive number and one negative which sum to a positive number"},
		{a: -10, b: 10, want: 0},
	}
	t.Parallel()
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f, name %v", tc.want, got, tc.name)
		}
	}
}

func TestSubtract(t *testing.T) {
	testCases := []testCase{
		{a: 0, b: 1, want: -1},
		{a: 10, b: 10, want: 0},
		{a: 0, b: 0, want: 0},
		{a: 2, b: -1, want: 3},
		{a: -10, b: -10, want: 0},
		{a: 15, b: 5, want: 10},
	}
	t.Parallel()
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	testCases := []testCase{
		{a: 0, b: 1, want: 0},
		{a: 10, b: 10, want: 100},
		{a: 0, b: 0, want: 0},
		{a: 2, b: -1, want: -2},
		{a: -10, b: -10, want: 100},
		{a: 15, b: 5, want: 75},
	}
	t.Parallel()
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}
