/*
Implementation of Python's strftime and strptime in Go
Example:
    str, err := timefmt.Strftime(time.Now(), "%Y/%m/%d") // 2016/09/22
Directives:
    %a - Locale’s abbreviated weekday name
    %A - Locale’s full weekday name
    %b - Locale’s abbreviated month name
    %B - Locale’s full month name
    %c - Locale’s appropriate date and time representation
    %d - Day of the month as a decimal number [01,31]
    %H - Hour (24-hour clock) as a decimal number [00,23]
    %I - Hour (12-hour clock) as a decimal number [01,12]
    %j - Day of year
    %m - Month as a decimal number [01,12]
    %M - Minute as a decimal number [00,59]
    %p - Locale’s equivalent of either AM or PM
    %S - Second as a decimal number [00,61]
    %U - Week number of the year
    %w - Weekday as a decimal number
    %W - Week number of the year
    %x - Locale’s appropriate date representation
    %X - Locale’s appropriate time representation
    %y - Year without century as a decimal number [00,99]
    %Y - Year with century as a decimal number
    %Z - Time zone name (no characters if no time zone exists)
Note that %c returns RFC1123 which is a bit different from what Python does
*/
package timefmt

import (
	"time"
)

var longDayNames = []string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

var shortDayNames = []string{
	"Sun",
	"Mon",
	"Tue",
	"Wed",
	"Thu",
	"Fri",
	"Sat",
}

var shortMonthNames = []string{
	"---",
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec",
}

var longMonthNames = []string{
	"---",
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

func Strftime(t time.Time, format string) (string, error) {

	return "", nil
}

func Strptime(value string, format string) (time.Time, error) {

	return time.Time{}, nil
}

func Check(format string) error {
	return nil
}