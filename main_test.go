package main

import (
	"fmt"
	"testing"
)

func TestMergeAlternately(t *testing.T) {
	var tests = []struct {
		word1  string
		word2  string
		merged string
	}{
		{"abc", "pqr", "apbqcr"},
		{"ab", "pqrs", "apbqrs"},
		{"abcd", "pq", "apbqcd"},
	}

	for _, test := range tests {
		testname := fmt.Sprintf(`testing with words "%s", "%s"`, test.word1, test.word2)

		t.Run(testname, func(t *testing.T) {
			merged := mergeAlternately(test.word1, test.word2)

			if merged != test.merged {
				t.Errorf(`testing mergeAlternately("%s","%s"), expected "%s", but got "%s"`, test.word1, test.word2, test.merged, merged)
			}
		})
	}
}
