// Package loc provides utility functions to count lines of code in files
package loc

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"path/filepath"
)

// CountLinesOfCode returns the number of lines of code for a certain file in a specific package directory
//
// It returns an error if there is something unexpected with the provided file
func CountLinesOfCode(packageDir, srcFile string) (int, error) {
	f, err := os.Open(filepath.Join(packageDir, srcFile))
	if err != nil {
		return 0, err
	}
	defer f.Close()

	return countLines(f)
}

func countLines(r io.Reader) (int, error) {
	var count int
	const lineBreak = '\n'

	buf := make([]byte, bufio.MaxScanTokenSize)

	for {
		bufferSize, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}

		var buffPosition int
		for {
			i := bytes.IndexByte(buf[buffPosition:], lineBreak)
			if i == -1 || bufferSize == buffPosition {
				break
			}
			buffPosition += i + 1
			count++
		}
		if err == io.EOF {
			break
		}
	}

	if count == 0 {
		return 0, nil
	}

	return count + 1, nil
}
