package stringsizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransform(t *testing.T) {
	baseChars = "abc"
	base = float64(len(baseChars))
	// Table driven tests, see https://medium.com/@sebdah/go-best-practices-testing-3448165a0e18
	tests := []struct {
		num    int
		result string
	}{
		{0, "a"},
		{1, "b"},
		{2, "c"},
		{3, "ba"},
		{4, "bb"},
		{5, "bc"},
		{243, "baaaaa"},
	}

	for _, test := range tests {
		result := Transform(test.num)
		assert.Equal(t, test.result, result)
	}
}

func TestTransform2(t *testing.T) {
	alreadyHave := make(map[string]struct{})
	for i := 0; i < 300; i++ {
		s := Transform(i)
		if _, ok := alreadyHave[s]; !ok {
			alreadyHave[s] = struct{}{}
		} else {
			t.Errorf("already have one: %d, %s", i, s)
		}
	}
}
