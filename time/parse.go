package time

import (
	"time"
)

func UnixToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

func StringToTimeStamp(target string, format string) (int64, error) {
	t, err := time.Parse(format, target)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}
