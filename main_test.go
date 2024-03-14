package main

import (
	"fmt"
	"testing"
)

var tests = []struct {
	word1  string
	word2  string
	merged string
}{
	{"abc", "pqr", "apbqcr"},
	{"ab", "pqrs", "apbqrs"},
	{"abcd", "pq", "apbqcd"},
}

func TestMergeAlternately(t *testing.T) {
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

func TestMergeAlternately2(t *testing.T) {
	for _, test := range tests {
		testname := fmt.Sprintf(`testing with words "%s", "%s"`, test.word1, test.word2)

		t.Run(testname, func(t *testing.T) {
			merged := mergeAlternately2(test.word1, test.word2)

			if merged != test.merged {
				t.Errorf(`testing mergeAlternately2("%s","%s"), expected "%s", but got "%s"`, test.word1, test.word2, test.merged, merged)
			}
		})
	}
}

func TestMergeAlternatelyBytesBuffer(t *testing.T) {
	for _, test := range tests {
		testname := fmt.Sprintf(`testing with words "%s", "%s"`, test.word1, test.word2)

		t.Run(testname, func(t *testing.T) {
			merged := mergeAlternatelyBytesBuffer(test.word1, test.word2)

			if merged != test.merged {
				t.Errorf(`testing mergeAlternatelyBytesBuffer("%s","%s"), expected "%s", but got "%s"`, test.word1, test.word2, test.merged, merged)
			}
		})
	}
}

func TestMergeAlternatelyStringBuilder(t *testing.T) {

	for _, test := range tests {
		testname := fmt.Sprintf(`testing with words "%s", "%s"`, test.word1, test.word2)

		t.Run(testname, func(t *testing.T) {
			merged := mergeAlternatelyStringBuilder(test.word1, test.word2)

			if merged != test.merged {
				t.Errorf(`testing mergeAlternatelyStringBuilder("%s","%s"), expected "%s", but got "%s"`, test.word1, test.word2, test.merged, merged)
			}
		})
	}
}

func TestMergeAlternatelyCopy(t *testing.T) {

	for _, test := range tests {
		testname := fmt.Sprintf(`testing with words "%s", "%s"`, test.word1, test.word2)

		t.Run(testname, func(t *testing.T) {
			merged := mergeAlternatelyCopy(test.word1, test.word2)

			if merged != test.merged {
				t.Errorf(`testing mergeAlternatelyCopy("%s","%s"), expected "%s", but got "%s"`, test.word1, test.word2, test.merged, merged)
			}
		})
	}
}

func TestMergeAlternatelyRecursive(t *testing.T) {

	for _, test := range tests {
		testname := fmt.Sprintf(`testing with words "%s", "%s"`, test.word1, test.word2)

		t.Run(testname, func(t *testing.T) {
			merged := mergeAlternatelyRecursive(test.word1, test.word2)

			if merged != test.merged {
				t.Errorf(`testing mergeAlternatelyRecursive("%s","%s"), expected "%s", but got "%s"`, test.word1, test.word2, test.merged, merged)
			}
		})
	}
}

func BenchmarkMergeAlternately(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeAlternately("abc", "def31")
	}
}

func BenchmarkMergeAlternately2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeAlternately2("abc", "def31")
	}
}

func BenchmarkMergeAlternatelyBytesBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeAlternatelyBytesBuffer("abc", "def31")
	}
}

func BenchmarkMergeAlternatelyStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeAlternatelyStringBuilder("abc", "def31")
	}
}

func BenchmarkMergeAlternatelyStringBuilder2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeAlternatelyStringBuilder2("abc", "def31")
	}
}

func BenchmarkMergeAlternatelyCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeAlternatelyCopy("abc", "def31")
	}
}

func BenchmarkMergeAlternatelyRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeAlternatelyRecursive("abc", "def31")
	}
}
