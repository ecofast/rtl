// Copyright 2016~2017 ecofast(无尽愿). All rights reserved.
// Use of this source code is governed by a BSD-style license.

// Package timeutils implements some useful date&time utility functions.
package timeutils

import (
	"fmt"
	"time"
)

func MilliSecondsBetween(now, then time.Time) int64 {
	return int64(now.Sub(then) / time.Millisecond)
}

func SecondsBetween(now, then time.Time) int64 {
	return MilliSecondsBetween(now, then) / 1000
}

func DateTimeToStr(dt time.Time) string {
	year, month, day := dt.Date()
	hour, min, sec := dt.Clock()
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", year, int(month), day, hour, min, sec)
}

func DateToStr(dt time.Time) string {
	year, month, day := dt.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, int(month), day)
}

func TimeToStr(dt time.Time) string {
	hour, min, sec := dt.Clock()
	return fmt.Sprintf("%02d:%02d:%02d", hour, min, sec)
}
