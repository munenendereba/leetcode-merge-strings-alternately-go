package main

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
