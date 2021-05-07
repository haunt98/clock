package clock

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormatDate(t *testing.T) {
	tests := []struct {
		name string
		when time.Time
		want string
	}{
		{
			name: "2020-04-01",
			when: time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local),
			want: "2020-04-01",
		},
		{
			name: "2020-01-18",
			when: time.Date(2020, 1, 18, 0, 0, 0, 0, time.Local),
			want: "2020-01-18",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := FormatDate(tc.when)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestFormatDateTime(t *testing.T) {
	tests := []struct {
		name string
		when time.Time
		want string
	}{
		{
			name: "2020-04-01 23:59:0",
			when: time.Date(2020, 4, 1, 23, 59, 0, 0, time.Local),
			want: "2020-04-01 23:59:00",
		},
		{
			name: "2020-01-18 0:12:24",
			when: time.Date(2020, 1, 18, 0, 12, 24, 0, time.Local),
			want: "2020-01-18 00:12:24",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := FormatDateTime(tc.when)
			assert.Equal(t, tc.want, got)
		})
	}
}
