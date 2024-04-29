package solution

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var track = make([][]int, 26)

type Step struct {
	Index int
	Cost  int
}

func findRotateSteps(ring string, key string) int {
	resetArr(track, nil)

	for i := 0; i < len(key); i++ {
		if track[key[i]-'a'] == nil {
			track[key[i]-'a'] = make([]int, 0)
		}
	}

	// store the index of each char that can build the key in increasing manner
	var char int
	for i := 0; i < len(ring); i++ {
		char = int(ring[i] - 'a')
		track[char] = append(track[char], i)

	}

	getMinCost := func(currIdx int, targetIdx int) int {
		diff := abs(targetIdx - currIdx)

		return min(diff, len(ring)-diff)
	}

	// let dp[i][j] = the minimum steps to reach key[j] when the ring position is at i.
	// - so i is related to the current position at the ring
	// - j is position at the key
	//
	// initially we are at dp[0][0]
	// - if ring[0] == key[0] -> dp[0][0] = 1
	// - if not, then we need to search the position of key[0] in the rings, that is track[key[0]] which lists all the positions
	//   - for pos in track[key[0]]
	//        dp[currPos][pos] = getMinCost(currIdx, pos)
	dp := make([][]int, len(ring))
	for i := 0; i < len(ring); i++ {
		dp[i] = make([]int, len(key))
	}
	fillArr(dp, math.MaxInt32)

	for _, targetPos := range track[int(key[0]-'a')] {
		for _, currPos := range track[int(ring[0])-'a'] {
			initCost := getMinCost(0, currPos)
			dp[targetPos][0] = min(dp[targetPos][0], initCost+getMinCost(currPos, targetPos)+1)
		}
	}
	print2dArr(dp)

	for i := 1; i < len(key); i++ {
		for _, targetPos := range track[int(key[i]-'a')] {
			for _, prevPos := range track[int(key[i-1]-'a')] {
				dp[targetPos][i] = min(dp[targetPos][i], dp[prevPos][i-1]+getMinCost(prevPos, targetPos)+1)
			}
		}
	}
	print2dArr(dp)
	minStep := dp[0][len(key)-1]
	for i := 1; i < len(ring); i++ {
		minStep = min(minStep, dp[i][len(key)-1])
	}
	return minStep
}

func resetArr[T any](arr []T, val T) {
	for i := 0; i < len(arr); i++ {
		arr[i] = val
	}
}
func fillArr[T any](arr [][]T, val T) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			arr[i][j] = val
		}
	}
}
func print2dArr[T any](arr [][]T) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	fmt.Println("==============")
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func TestCase(t *testing.T) {
	assert.Equal(t, 4, findRotateSteps("godding", "gd"))
	assert.Equal(t, 7, findRotateSteps("godding", "gdi"))
	assert.Equal(t, 13, findRotateSteps("godding", "godding"))

	assert.Equal(t, 6, findRotateSteps("abcde", "ade"))

	assert.Equal(t, 19, findRotateSteps("nyngl", "yyynnnnnnlllggg"))
	assert.Equal(t, 2, findRotateSteps("eh", "h"))
}
