package main

import (
	"testing"
	"wordle/util"
)

func TestCheckGuess(t *testing.T) {
	// Defining the columns of the table
	var tests = []struct {
		name  string
		guess string
		word  string
		want  []string
	}{
		{"All correct", "ASDFG", "ASDFG", []string{"correct", "correct", "correct", "correct", "correct"}},
		{"1st absent", "XSDFG", "ASDFG", []string{"absent", "correct", "correct", "correct", "correct"}},
		{"2nd absent", "AXDFG", "ASDFG", []string{"correct", "absent", "correct", "correct", "correct"}},
		{"3rd absent", "ASXFG", "ASDFG", []string{"correct", "correct", "absent", "correct", "correct"}},
		{"4th absent", "ASDXG", "ASDFG", []string{"correct", "correct", "correct", "absent", "correct"}},
		{"5th absent", "ASDFX", "ASDFG", []string{"correct", "correct", "correct", "correct", "absent"}},
		{"All absent", "XXXXX", "ASDFG", []string{"absent", "absent", "absent", "absent", "absent"}},
		{"1,2 present 3 absent 4,5 correct", "SAXFG", "ASDFG", []string{"present", "present", "absent", "correct", "correct"}},
	}

	// The execution loop
	for _, _test := range tests {
		t.Run(_test.name, func(t *testing.T) {
			checkedGuess := checkGuess(_test.guess, _test.word)
			areEqual := util.SlicesEqual(checkedGuess, _test.want)

			if !areEqual {
				t.Errorf("got %s, want %s", checkedGuess, _test.want)
			}
		})
	}
}
