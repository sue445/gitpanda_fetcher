package fetcher

import (
	"fmt"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
)

// truncateWithLine truncate string with line
func truncateWithLine(str string, maxLines int) string {
	if maxLines < 1 {
		return str
	}

	lines := strings.Split(str, "\n")

	if len(lines) > maxLines {
		lines = lines[:maxLines]
		return strings.Join(lines, "\n")
	}

	return str
}

// selectLine returns a specific line in the text
func selectLine(str string, line int) string {
	lines := strings.Split(str, "\n")

	line = clipNumber(line, 1, len(lines))

	return lines[line-1]
}

// selectLines returns a specific lines in the text
func selectLines(str string, startLine int, endLine int) string {
	lines := strings.Split(str, "\n")

	startLine = clipNumber(startLine, 1, len(lines))
	endLine = clipNumber(endLine, 1, len(lines))

	if startLine > endLine {
		startLine, endLine = endLine, startLine
	}

	return strings.Join(lines[startLine-1:endLine], "\n")
}

func clipNumber(number int, lower int, upper int) int {
	if number < lower {
		return lower
	}

	if number > upper {
		return upper
	}

	return number
}

// withDebugLogging executes blocks and outputs debug logs if necessary
func withDebugLogging[T any](label string, isDebugLogging bool, fn func() (*T, error)) (*T, error) {
	start := time.Now()

	ret, err := fn()

	if err != nil {
		return nil, errors.WithStack(err)
	}

	if isDebugLogging {
		duration := time.Since(start)
		fmt.Printf("[DEBUG] %s : duration=%s, ret=%+v\n", label, duration, ret)
	}

	return ret, nil
}
