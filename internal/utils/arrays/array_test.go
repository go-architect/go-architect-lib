package arrays_test

import (
	"fmt"
	"github.com/fdaines/go-architect-lib/internal/utils/arrays"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicatedStrings(t *testing.T) {
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

func TestIndexOf(t *testing.T) {
	var tests = []struct {
		input []string
		want  int
	}{
		{[]string{"a", "b", "c"}, 0},
		{[]string{"b", "a", "c", "b", "f", "a"}, 1},
		{[]string{"x", "y"}, -1},
		{[]string{"a", "a", "a"}, 0},
	}

	for i, tt := range tests {
		testCase := fmt.Sprintf("input case: %d", i)
		t.Run(testCase, func(t *testing.T) {
			answer := arrays.IndexOf(tt.input, "a")
			assert.Equal(t, tt.want, answer)
		})
	}
}

func TestContains(t *testing.T) {
	var tests = []struct {
		input []string
		want  bool
	}{
		{[]string{"a", "b", "c"}, true},
		{[]string{"b", "a", "c", "b", "f", "a"}, true},
		{[]string{"x", "y"}, false},
		{[]string{}, false},
		{[]string{"a", "a", "a"}, true},
	}

	for i, tt := range tests {
		testCase := fmt.Sprintf("input case: %d", i)
		t.Run(testCase, func(t *testing.T) {
			answer := arrays.Contains(tt.input, "a")
			assert.Equal(t, tt.want, answer)
		})
	}
}

func TestDequeue(t *testing.T) {
	var tests = []struct {
		input []string
		want  string
		err   error
	}{
		{input: []string{"a", "b", "c"}, want: "a", err: nil},
		{[]string{"b", "a", "c", "b", "f", "a"}, "b", nil},
		{input: []string{}, want: "", err: fmt.Errorf("queue is empty2")},
	}

	for i, tt := range tests {
		testCase := fmt.Sprintf("input case: %d", i)
		t.Run(testCase, func(t *testing.T) {
			answer, err := arrays.Dequeue(tt.input)
			assert.Equal(t, tt.want, answer)
			if tt.err == nil {
				assert.Nil(t, err)
			} else {
				assert.Error(t, err, tt.err.Error())
			}
		})
	}
}
