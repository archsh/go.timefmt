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
	"bytes"
	"errors"
	//"fmt"
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

type OutputConverter func(time.Time, bool)(string, error) /** flags: true means with '-' */

var ontput_converters = map[rune]OutputConverter {
	//| %a	| Weekday as locale’s abbreviated name.	| Mon|
	'a': cvt_output_a,
	//| %A	| Weekday as locale’s full name.	| Monday|
	'A': cvt_output_A,
	//| %w	| Weekday as a decimal number, where 0 is Sunday and 6 is Saturday.	| 1|
	'w': cvt_output_w,
	//| %d	| Day of the month as a zero-padded decimal number.	| 30|
	//| %-d	| Day of the month as a decimal number. (Platform specific)	| 30|
	'd': cvt_output_d,
	//| %b	| Month as locale’s abbreviated name.	| Sep|
	'b': cvt_output_b,
	//| %B	| Month as locale’s full name.	| September|
	'B': cvt_output_B,
	//| %m	| Month as a zero-padded decimal number.	| 09|
	//| %-m	| Month as a decimal number. (Platform specific)	| 9|
	'm': cvt_output_m,
	//| %y	| Year without century as a zero-padded decimal number.	| 13|
	'y': cvt_output_y,
	//| %Y	| Year with century as a decimal number.	| 2013|
	'Y': cvt_output_Y,
	//| %H	| Hour (24-hour clock) as a zero-padded decimal number.	| 07|
	//| %-H	| Hour (24-hour clock) as a decimal number. (Platform specific)	| 7|
	'H': cvt_output_H,
	//| %I	| Hour (12-hour clock) as a zero-padded decimal number.	| 07|
	//| %-I	| Hour (12-hour clock) as a decimal number. (Platform specific)	| 7|
	'I': cvt_output_I,
	//| %p	| Locale’s equivalent of either AM or PM.	| AM|
	'p': cvt_output_p,
	//| %M	| Minute as a zero-padded decimal number.	| 06|
	//| %-M	| Minute as a decimal number. (Platform specific)	| 6|
	'M': cvt_output_M,
	//| %S	| Second as a zero-padded decimal number.	| 05|
	//| %-S	| Second as a decimal number. (Platform specific)	| 5|
	'S': cvt_output_S,
	//| %f	| Microsecond as a decimal number, zero-padded on the left.	| 000000|
	'f': cvt_output_f,
	//| %z	| UTC offset in the form +HHMM or -HHMM (empty string if the the object is naive).	| |
	'z': cvt_output_z,
	//| %Z	| Time zone name (empty string if the object is naive).	| |
	'Z': cvt_output_Z,
	//| %j	| Day of the year as a zero-padded decimal number.	| 273|
	//| %-j	| Day of the year as a decimal number. (Platform specific)	| 273|
	'j': cvt_output_j,
	//| %U	| Week number of the year (Sunday as the first day of the week) as a zero padded decimal number. All days in a new year preceding the first Sunday are considered to be in week 0.	| 39|
	'U': cvt_output_U,
	//| %W	| Week number of the year (Monday as the first day of the week) as a decimal number. All days in a new year preceding the first Monday are considered to be in week 0.	| 39|
	'W': cvt_output_W,
	//| %c	| Locale’s appropriate date and time representation.	| Mon Sep 30 07:06:05 2013|
	'c': cvt_output_c,
	//| %x	| Locale’s appropriate date representation.	| 09/30/13|
	'x': cvt_output_x,
	//| %X	| Locale’s appropriate time representation.	| 07:06:05|
	'X': cvt_output_X,
	//| %%	| A literal '%' character.	| %|
	'%': cvt_output_percent,
}

func Strftime(t time.Time, format string) (string, error) {
	buf := bytes.Buffer{}
	length := len(format)
	for i, j := 0, 0; i < length; i+=j+1 {
		c := format[i]
		j = 0
		flag := false
		//fmt.Printf("c > %c \n",c)
		if c != 0x25 || (i+1) >= length { // "%" -> 0x25
			buf.WriteByte(c)
			continue
		}
		if format[i+1] == 0x2d { // "-" -> 0x2d
			flag = true
			j += 1
		}
		if (i+j+1) < length {
			//fmt.Printf("format[i+j+1]> %c \n",format[i+j+1])
			if cvt_func, ok := ontput_converters[rune(format[i+j+1])]; ok {
				s, e := cvt_func(t, flag)
				if e != nil {
					return "", e
				}
				buf.WriteString(s)
			}else{
				return "", errors.New("Unknown Code:"+format[i+j:i+j+1])
			}
			j += 1
		}else{
			buf.WriteByte(format[i+1])
		}
	}
	return buf.String(), nil
}

func Strptime(value string, format string) (time.Time, error) {

	return time.Time{}, nil
}