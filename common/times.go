package common

import "time"

var timeBase = time.Date(2019, 9, 1, 0, 0, 0, 0, time.Now().Location()).Unix()

func GetDayNo(args ...int64) int {
	var t int64
	if len(args) > 0 {
		t = args[0]
	} else {
		t = time.Now().Unix()
	}
	return int((t-timeBase)/86400 + 1)
}