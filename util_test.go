package fetcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTruncateWithLine(t *testing.T) {
	type args struct {
		str      string
		maxLines int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "maxLines < 1",
			args: args{
				str:      "a\nb\nc\n",
				maxLines: 0,
			},
			want: "a\nb\nc\n",
		},
		{
			name: "lines <= maxLines",
			args: args{
				str:      "a\nb\n",
				maxLines: 3,
			},
			want: "a\nb\n",
		},
		{
			name: "lines > maxLines",
			args: args{
				str:      "a\nb\nc\nd\n",
				maxLines: 3,
			},
			want: "a\nb\nc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := truncateWithLine(tt.args.str, tt.args.maxLines)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSelectLine(t *testing.T) {
	str := `111
222
333
`
	type args struct {
		str  string
		line int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Select 2nd line",
			args: args{
				str:  str,
				line: 2,
			},
			want: "222",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selectLine(tt.args.str, tt.args.line); got != tt.want {
				t.Errorf("selectLine() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestSelectLines(t *testing.T) {
	str := `111
222
333
444
555
`

	type args struct {
		str       string
		startLine int
		endLine   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Select lines",
			args: args{
				str:       str,
				startLine: 2,
				endLine:   4,
			},
			want: "222\n333\n444",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := selectLines(tt.args.str, tt.args.startLine, tt.args.endLine)

			assert.Equal(t, tt.want, got)
		})
	}
}
