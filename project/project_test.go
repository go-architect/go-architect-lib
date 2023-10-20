package project_test

import (
	"fmt"
	"github.com/go-architect/go-architect-lib/internal/utils"
	"testing"
)

func TestRoundFloat(t *testing.T) {
	var tests = []struct {
		input     float64
		precision uint
		want      float64
	}{
		{45.67859, 2, 45.68},
		{45, 2, 45},
		{0.123456, 4, 0.1235},
		{12.98765, 0, 13},
		{0.00000000, 2, 0},
	}

	for _, tt := range tests {
		testCase := fmt.Sprintf("input: %v", tt.input)
		t.Run(testCase, func(t *testing.T) {
			answer := utils.RoundFloat(tt.input, tt.precision)
			if answer != tt.want {
				t.Errorf("got %v, want %v", answer, tt.want)
			}
		})
	}
}
