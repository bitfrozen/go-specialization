package main

import "testing"

//goland:noinspection SpellCheckingInspection
func TestFindian(t *testing.T) {
	type test struct {
		input string
		want  bool
	}

	tests := []test{
		{input: "ian", want: true},
		{input: "Ian", want: true},
		{input: "iuiygaygn", want: true},
		{input: "I d skd a efju N", want: true},
		{input: "ihhhhhn", want: false},
		{input: "ina", want: false},
		{input: "xian", want: false},
	}

	for _, testCase := range tests {
		got := checkString(testCase.input)
		if testCase.want != got {
			t.Errorf("input: %s, expected: %t, got: %t", testCase.input, testCase.want, got)
		}
	}
}
