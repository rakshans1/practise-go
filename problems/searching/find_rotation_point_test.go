/*
Problem:
- Given a list of words, return an index of a rotation point.
Example:
- Input:
       []string{
	       "ptolemaic",
	       "retrograde",
	       "supplant",
	       "undulate",
	       "xenoepist",
	       "asymptote", // <-- rotates here!
	       "babka",
	       "banoffee",
	       "engender",
	       "karpatka",
	       "othellolagkage",
       },
  Output: 5
Approach:
- Use a binary search approach to search for word, to either go right or go
  left because they are in order alphabetically.
- Keep track of the lower and upper bounds and so that when they are next
  to each other, the floor is the last item while the ceiling is the rotation
  point.
Solution:
- Keep track of the lower and upper bounds on the rotation points.
- While is lower bound is less than the upper bound, check if the guessed word
  (the one in the middle of the range), comes after the first word.
- If so, go right by making the floor index be the guessed index.
- Otherwise, go left by making the ceiling index be the guessed index.
- When the floor index and ceiling index have converged, the ceiling is
  the rotation point.
Cost:
- O(logn) time, O(1) space.
*/

package problems

import (
	"practise-go/common"
	"testing"
)

func TestFindRotationPoint(t *testing.T) {
	tests := []struct {
		in       []string
		expected int
	}{
		{
			[]string{}, 0,
		},
		{
			[]string{
				"xenoepist",
				"asymptote", // <-- rotates here!
			},
			1,
		},
		{
			[]string{
				"xenoepist",
				"asymptote", // <-- rotates here!
				"babka",
				"banoffee",
				"engender",
				"karpatka",
				"othellolagkage",
			},
			1,
		},
		{
			[]string{
				"asymptote",
				"babka",
				"banoffee",
				"engender",
				"karpatka",
				"othellolagkage",
			},
			5, // returns the last index if there's no rotation
		},
		{
			[]string{
				"ptolemaic",
				"retrograde",
				"supplant",
				"undulate",
				"xenoepist",
			},
			4, // returns the last index if there's no rotation
		},
		{
			[]string{
				"ptolemaic",
				"retrograde",
				"supplant",
				"undulate",
				"xenoepist",
				"asymptote", // <-- rotates here!
				"babka",
				"banoffee",
				"engender",
				"karpatka",
				"othellolagkage",
			},
			5,
		},
	}

	for _, tt := range tests {
		result := findRotationPoint(tt.in)
		common.Equal(t, tt.expected, result)
	}
}

func findRotationPoint(words []string) int {
	if len(words) <= 1 {
		return 0
	}
	firstWord := words[0]
	floorIndex := 0
	ceilingIndex := len(words) - 1
	for floorIndex < ceilingIndex {
		guessIndex := floorIndex + (ceilingIndex-floorIndex)/2

		if words[guessIndex] >= firstWord {
			floorIndex = guessIndex
		} else {
			ceilingIndex = guessIndex
		}

		if floorIndex+1 == ceilingIndex {
			return ceilingIndex
		}
	}
	return -1
}
