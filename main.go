package main

import (
	"bytes"
	"strings"
)

func main() {

}

func mergeAlternately(word1 string, word2 string) string {
	len1, len2 := len(word1), len(word2)
	merged := ""

	lenuse := 0

	lenuse = len1

	if len1 < len2 {
		lenuse = len2
	}

	for i := 0; i < lenuse; i++ {
		c1, c2 := "", ""
		if i+1 <= len1 {
			c1 = string(word1[i])
		}

		if i+1 <= len2 {
			c2 = string(word2[i])
		}
		merged += c1 + c2
	}

	return merged
}

func mergeAlternately2(word1 string, word2 string) string {
	var ans string
	for len(word1) > 0 && len(word2) > 0 {
		ans += string(word1[0]) + string(word2[0])
		word1 = word1[1:]
		word2 = word2[1:]
	}
	if len(word1) > 0 {
		ans += word1
	}
	if len(word2) > 0 {
		ans += word2
	}

	return ans
}

func mergeAlternatelyBytesBuffer(word1 string, word2 string) string {
	var b bytes.Buffer
	i, j := 0, 0
	for i < len(word1) || j < len(word2) {
		if i < len(word1) {
			b.WriteByte(word1[i])
			i++
		}
		if j < len(word2) {
			b.WriteByte(word2[j])
			j++
		}
	}
	return b.String()
}

func mergeAlternatelyStringBuilder(word1 string, word2 string) string {
	len1, len2 := len(word1), len(word2)
	maxLength := len1 + len2

	var builder strings.Builder
	builder.Grow(maxLength)

	i, j := 0, 0
	for i < len1 || j < len2 {
		if i < len1 {
			builder.WriteByte(word1[i])
			i++
		}
		if j < len2 {
			builder.WriteByte(word2[j])
			j++
		}
	}

	return builder.String()
}

func mergeAlternatelyStringBuilder2(word1 string, word2 string) string {
	len1, len2 := len(word1), len(word2)

	var builder strings.Builder
	builder.Grow(len1 + len2)

	i, j := 0, 0
	for i < len1 || j < len2 {
		if i < len1 {
			builder.WriteByte(word1[i])
			i++
		}
		if j < len2 {
			builder.WriteByte(word2[j])
			j++
		}
	}

	return builder.String()
}

func mergeAlternatelyCopy(word1, word2 string) string {
	stringlen := len(word1) + len(word2)

	byeteMerged := make([]byte, stringlen)
	byteLength := 0

	for len(word1) > 0 && len(word2) > 0 {
		byteLength += copy(byeteMerged[byteLength:], string(word1[0]))
		byteLength += copy(byeteMerged[byteLength:], string(word2[0]))

		word1 = word1[1:]
		word2 = word2[1:]
	}

	byteLength += copy(byeteMerged[byteLength:], word1)

	copy(byeteMerged[byteLength:], word2)

	return string(byeteMerged)
}

func mergeAlternatelyRecursive(word1 string, word2 string) string {
	if len(word1) == 0 {
		return word2
	}

	if len(word2) == 0 {
		return word1
	}

	if len(word1) == 1 && len(word2) == 1 {
		return word1 + word2
	}

	return string(word1[0]) + string(word2[0]) + mergeAlternately(word1[1:], word2[1:])
}
