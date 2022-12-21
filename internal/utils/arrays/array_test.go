package arrays_test

import (
	"fmt"
	"github.com/fdaines/go-architect-lib/internal/utils/arrays"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicateStr(t *testing.T) {
	var tests = []struct {
		input []string
		want  []string
	}{
		{[]string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{[]string{"a", "b", "c", "b", "f", "a"}, []string{"a", "b", "c", "f"}},
		{[]string{}, []string{}},
		{[]string{"a", "a", "a"}, []string{"a"}},
	}

	for i, tt := range tests {
		testCase := fmt.Sprintf("input case: %d", i)
		t.Run(testCase, func(t *testing.T) {
			answer := arrays.RemoveDuplicatedStrings(tt.input)
			assert.ElementsMatch(t, tt.want, answer)
		})
	}
}
