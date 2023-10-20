package packages_test

import (
	"fmt"
	"github.com/go-architect/go-architect-lib/internal/utils/packages"
	"testing"
)

func TestIsInternal(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"github.com/foobar/mypackage", true},
		{"fooBar", false},
		{"golang.org/x/foobar", false},
		{"golang.org/x", false},
		{"github.com/foobar/another/packages", true},
	}

	for _, tt := range tests {
		testCase := fmt.Sprintf("input: %s", tt.input)
		t.Run(testCase, func(t *testing.T) {
			ans := packages.IsInternalPackage(tt.input, "github.com/foobar")
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
