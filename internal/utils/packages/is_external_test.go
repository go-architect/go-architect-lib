package packages_test

import (
	"fmt"
	"github.com/go-architect/go-architect-lib/internal/utils/packages"
	"testing"
)

func TestIsExternal(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"foobar/x", false},
		{"fooBar", false},
		{"golang.org/x/foobar", false},
		{"golang.org/x", false},
		{"foobar/bar", false},
		{"blablabla/xxxxxx", false},
		{"blablabla.com/xxxxxx", true},
	}

	for _, tt := range tests {
		testCase := fmt.Sprintf("input: %s", tt.input)
		t.Run(testCase, func(t *testing.T) {
			ans := packages.IsExternalPackage(tt.input, "foobar")
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
