package packages_test

import (
	"github.com/fdaines/go-architect-lib/internal/utils/packages"
	"github.com/stretchr/testify/assert"
	"go/build"
	"testing"
)

func TestGetCodeFiles(t *testing.T) {
	t.Run("Return list of code files from packages", func(t *testing.T) {
		pkg := &build.Package{
			GoFiles:     []string{"foo", "bar"},
			CgoFiles:    []string{"abc1", "abc2", "abc3"},
			TestGoFiles: []string{"test1", "test2", "test3", "test4"},
		}
		expected := []string{"abc1", "abc2", "abc3", "foo", "bar", "test1", "test2", "test3", "test4"}

		codeFiles := packages.GetCodeFiles(pkg)
		assert.Equal(t, 9, len(codeFiles))
		assert.ElementsMatch(t, expected, codeFiles)
	})

	t.Run("Return empty list of code files from packages", func(t *testing.T) {
		pkg := &build.Package{}

		codeFiles := packages.GetCodeFiles(pkg)
		assert.Equal(t, 0, len(codeFiles))
	})
}
