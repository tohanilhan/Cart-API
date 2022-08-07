package utils

import (
	"testing"
	"time"
)

func TestGetTimestamp(t *testing.T) {

	timeMillis := time.Now().UTC().UnixMilli()
	timex := time.UnixMilli(timeMillis)
	timex = timex.Add(time.Duration(3) * time.Hour)
	timestamp := timex.UTC().Format("2006-January-02 15:04:05")

	tests := []struct {
		name  string
		want  string
		want1 int64
		want2 string
	}{
		// Add test cases.

		{
			name:  "Get timestamp",
			want:  timestamp,
			want1: time.Now().UTC().UnixMilli(),
			want2: timex.Month().String(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := GetTimestamp()
			if got != tt.want {
				t.Errorf("GetTimestamp() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetTimestamp() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("GetTimestamp() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
