package com

import (
	"math"
	"time"
)

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

func GetWeekNo(args ... int64) int {
	var t int64
	if len(args) > 0 {
		t = args[0]
	} else {
		t = time.Now().Unix()
	}
	dayNo := GetDayNo(t)
	return int(math.Ceil(float64(dayNo)/7))
}