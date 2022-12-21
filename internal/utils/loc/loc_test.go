package loc

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCountLoc(t *testing.T) {
	t.Run("Return lines of code", func(t *testing.T) {
		content := "hello world\nand hello\n go and\n more\nblablabla\nbye"
		myReader := strings.NewReader(content)

		expected := 6

		loc, err := countLines(myReader)

		assert.Nil(t, err)
		assert.Equal(t, expected, loc)
	})

	t.Run("Return lines of code for empty file", func(t *testing.T) {
		content := ""
		myReader := strings.NewReader(content)

		expected := 0

		loc, err := countLines(myReader)

		assert.Nil(t, err)
		assert.Equal(t, expected, loc)
	})
}
