package fetcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_normalizeJobTrace(t *testing.T) {
	raw := ReadTestData("testdata/job_trace.txt")
	got := normalizeJobTrace(raw)

	expected := ReadTestData("testdata/job_trace_plain.txt")
	assert.Equal(t, expected, got)
}
