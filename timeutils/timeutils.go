// Copyright 2016~2017 ecofast(无尽愿). All rights reserved.
// Use of this source code is governed by a BSD-style license.

// Package timeutils implements some useful date&time utility functions.
package timeutils

import (
	"fmt"
	"time"
)

const (
	// Units of time
	HoursPerDay = 24
	MinsPerHour = 60
	SecsPerMin  = 60
	MSecsPerSec = 1000
	MinsPerDay  = HoursPerDay * MinsPerHour
	SecsPerDay  = MinsPerDay * SecsPerMin
	MSecsPerDay = SecsPerDay * MSecsPerSec

	DaysPerWeek        = 7
	MonthsPerYear      = 12
	YearsPerDecade     = 10
	YearsPerCentury    = 100
	YearsPerMillennium = 1000

	// Days between 1/1/0001 and 12/31/1899
	DateDelta = 693594

	// Days between DateTime basis (12/31/1899) and Unix time_t basis (1/1/1970)
	UnixDateDelta = 25569
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

// DecodeDate decodes the integral (date) part of the given time.Time value
// into its corresponding year, month, and day.
func DecodeDate(dt time.Time) (year, month, day int) {
	y, m, d := dt.Date()
	return y, int(m), d
}

// DecodeTime decodes the fractional (time) part of the given TDateTime value
// into its corresponding hour, minute and second.
func DecodeTime(dt time.Time) (hour, min, sec int) {
	return dt.Clock()
}

func IsLeapYear(dt time.Time) bool {
	year := dt.Year()
	return (year%4 == 0) && ((year%100 != 0) || (year%400 == 0))
}

func IsPM(dt time.Time) bool {
	return dt.Hour() >= 12
}

func IsToday(dt time.Time) bool {
	return IsSameDay(dt, time.Now())
}

func IsSameDay(dt1, dt2 time.Time) bool {
	year1, month1, day1 := DecodeDate(dt1)
	year2, month2, day2 := DecodeDate(dt2)
	return (year1 == year2) && (month1 == month2) && (day1 == day2)
}

func Yesterday() time.Time {
	return time.Now().Add(-time.Hour * HoursPerDay)
}

func Tomorrow() time.Time {
	return time.Now().Add(time.Hour * HoursPerDay)
}
