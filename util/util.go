package util

import (
	"strings"
	"time"
)

func GetTimestamp() (string, int64, string) {
	timeMillis := time.Now().UTC().UnixMilli()
	t := time.UnixMilli(timeMillis)
	t = t.Add(time.Duration(3) * time.Hour)
	timestamp := t.UTC().Format("2006-January-02 15:04:05")
	timestamp = strings.TrimSuffix(timestamp, "Z")
	return timestamp, timeMillis, t.Month().String()
}
