package utils

import (
	"fmt"
	"strings"
	timex "time"
)

func GetMoment() (string, int64) {
	utcTime := timex.Now().UTC()
	utcTimeMillis := utcTime.UnixMilli()
	localTime := timex.Now().Local()
	localTimestamp := localTime.Format(timex.RFC3339Nano)
	localTimestamp = strings.ReplaceAll(localTimestamp, "T", " ")
	localTimestamp = strings.ReplaceAll(localTimestamp, "+", " +")
	localTimestamp = localTimestamp[:strings.Index(localTimestamp, "+")-1]
	if strings.Contains(localTimestamp, ".") {
		remainder := localTimestamp[strings.Index(localTimestamp, ".")+1:]
		if len(remainder) > 3 {
			remainder = remainder[0:3]
		}
		localTimestamp = localTimestamp[:strings.Index(localTimestamp, ".")]
		localTimestamp = fmt.Sprintf("%s.%s", localTimestamp, remainder)
	}
	return localTimestamp, utcTimeMillis
}
