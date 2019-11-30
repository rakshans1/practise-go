/*
Problem:
- Given a sentence (string), return its word count map.
Example:
- Input: "Cliff finished his cake and paid the bill. Bill finished his cake at the edge of the cliff."
  Output: map[string]int{"cliff": 1, "Cliff": 1, "finished": 2, "his": 2, "cake": 2, "and": 1, "paid": 1, "the": 3, "bill": 1, "Bill": 1, "at": 1, "edge": 1, "of": 1}
Approach:
- First get rid of special characters, then use a hashmap to keep counts of words
  as we iterate through the string.
Solution:
- Get rid of special characters using regex then split the sentence by space.
- Choose to count the uppercase word only if it is uppercase in the original
  string. It's a reasonable approach, but not perfect one since proper nouns
  like "Bill" and "bill" point to the same meaning and are placed in a
  different place of a sentence - one in the front, other in the middle.
- Use a hashmap to keep counts of words.
Cost:
- O(n) time, O(n) space.
*/

package problems

import (
	"practise-go/common"
	"regexp"
	"strings"
	"testing"
)

func TestBuildWordCount(t *testing.T) {
	tests := []struct {
		in       string
		expected map[string]int
	}{
		{"", map[string]int{}},
		{" ", map[string]int{}},
		{".,#", map[string]int{}},
		{"1, 2, 3, 4", map[string]int{}},
		{"Bill.bill.", map[string]int{"bill": 1, "Bill": 1}},
		{
			"Cliff finished his cake and paid the bill. Bill finished his cake at the edge of the cliff.",
			map[string]int{"cliff": 1, "Cliff": 1, "finished": 2, "his": 2, "cake": 2, "and": 1, "paid": 1, "the": 3, "bill": 1, "Bill": 1, "at": 1, "edge": 1, "of": 1},
		},
	}

	for _, tt := range tests {
		result := buildWordCount(tt.in)
		common.Equal(t, tt.expected, result)
	}
}

func buildWordCount(in string) map[string]int {

	m := map[string]int{}

	s := regexp.MustCompile(`[^a-zA-Z]+`).ReplaceAllString(in, " ")
	if s == " " {
		return m
	}

	words := strings.Split(s, " ")

	for _, v := range words[:len(words)-1] {

		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	return m
}
