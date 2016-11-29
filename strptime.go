package timefmt

import (
	"time"
	"regexp"
	"bytes"
	"errors"
	//"fmt"
	"strconv"
)

var input_regexes = map[rune]string {
	//| %a	| Weekday as locale’s abbreviated name.	| Mon|
	'a': "(?P<a>Mon|Tue|Wed|Thu|Fri|Sat|Sun)",
	//| %A	| Weekday as locale’s full name.	| Monday|
	'A': "(?P<A>Monday|Tuesday|Wednsday|Thursday|Friday|Saturday|Sunday)",
	//| %w	| Weekday as a decimal number, where 0 is Sunday and 6 is Saturday.	| 1|
	//'w': "(?P<w>[0-6])",
	//| %d	| Day of the month as a zero-padded decimal number.	| 30|
	//| %-d	| Day of the month as a decimal number. (Platform specific)	| 30|
	'd': "(?P<d>[0-9]{1,2})",
	//| %b	| Month as locale’s abbreviated name.	| Sep|
	'b': "(?P<b>Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec)",
	//| %B	| Month as locale’s full name.	| September|
	'B': "(?P<B>January|February|March|April|May|June|July|August|September|October|November|December)",
	//| %m	| Month as a zero-padded decimal number.	| 09|
	//| %-m	| Month as a decimal number. (Platform specific)	| 9|
	'm': "(?P<m>[0-9]{1,2})",
	//| %y	| Year without century as a zero-padded decimal number.	| 13|
	'y': "(?P<y>[0-9]{2})",
	//| %Y	| Year with century as a decimal number.	| 2013|
	'Y': "(?P<Y>[0-9]{4})",
	//| %H	| Hour (24-hour clock) as a zero-padded decimal number.	| 07|
	//| %-H	| Hour (24-hour clock) as a decimal number. (Platform specific)	| 7|
	'H': "(?P<H>[0-9]{1,2})",
	//| %I	| Hour (12-hour clock) as a zero-padded decimal number.	| 07|
	//| %-I	| Hour (12-hour clock) as a decimal number. (Platform specific)	| 7|
	'I': "(?P<I>[0-9]{1,2})",
	//| %p	| Locale’s equivalent of either AM or PM.	| AM|
	'p': "(?P<p>AM|PM)",
	//| %M	| Minute as a zero-padded decimal number.	| 06|
	//| %-M	| Minute as a decimal number. (Platform specific)	| 6|
	'M': "(?P<M>[0-9]{1,2})",
	//| %S	| Second as a zero-padded decimal number.	| 05|
	//| %-S	| Second as a decimal number. (Platform specific)	| 5|
	'S': "(?P<S>[0-9]{1,2})",
	//| %f	| Microsecond as a decimal number, zero-padded on the left.	| 000000|
	'f': "(?P<f>[0-9]{6})",
	//| %z	| UTC offset in the form +HHMM or -HHMM (empty string if the the object is naive).	| |
	'z': "(?P<z>+|-[0-9]{4})",
	//| %Z	| Time zone name (empty string if the object is naive).	| |
	'Z': "(?P<Z>[a-zA-Z/_]{3,})",
	//| %j	| Day of the year as a zero-padded decimal number.	| 273|
	//| %-j	| Day of the year as a decimal number. (Platform specific)	| 273|
	//'j': "(?P<j>[0-9]{1,3})",
	//| %U	| Week number of the year (Sunday as the first day of the week) as a zero padded decimal number. All days in a new year preceding the first Sunday are considered to be in week 0.	| 39|
	//'U': "(?P<U>[0-9]{1,2})",
	//| %W	| Week number of the year (Monday as the first day of the week) as a decimal number. All days in a new year preceding the first Monday are considered to be in week 0.	| 39|
	//'W': "(?P<W>[0-9]{1,2})",
	//| %c	| Locale’s appropriate date and time representation.	| Mon Sep 30 07:06:05 2013|
	'c': "(?P<a>Mon|Tue|Wed|Thu|Fri|Sat|Sun) (?P<b>Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) (?P<d>[0-9]{2}) (?P<H>[0-9]{2}):(?P<M>[0-9]{2}):(?P<S>[0-9]{2}) (?P<Y>[0-9]{4})",
	//| %x	| Locale’s appropriate date representation.	| 09/30/13|
	'x': "(?P<m>[0-9]{2})/(?P<d>[0-9]{2})/(?P<y>[0-9]{2})",
	//| %X	| Locale’s appropriate time representation.	| 07:06:05|
	'X': "(?P<H>[0-9]{2}):(?P<M>[0-9]{2}):(?P<S>[0-9]{2})",
	//| %%	| A literal '%' character.	| %|
	'%': "",
}

type _DateTime struct {
	year int
	month time.Month
	day, hour, min, sec, nsec int
	loc *time.Location
	pm bool
}

var input_converters = map[rune]func(string, *_DateTime) error {
	//| %a	| Weekday as locale’s abbreviated name.	| Mon|
	//'a': func(val string, t *_DateTime) error {
	//	if nil == t {
	//		return errors.New("Invalid time parameter!")
	//	}
	//	return nil
	//},
	//| %A	| Weekday as locale’s full name.	| Monday|
	//'A': func(val string, t *_DateTime) error {
	//	if nil == t {
	//		return errors.New("Invalid time parameter!")
	//	}
	//	return nil
	//},
	//| %w	| Weekday as a decimal number, where 0 is Sunday and 6 is Saturday.	| 1|
	//'w': func(val string, t *_DateTime) error {
	//	if nil == t {
	//		return errors.New("Invalid time parameter!")
	//	}
	//	return nil
	//},
	//| %d	| Day of the month as a zero-padded decimal number.	| 30|
	//| %-d	| Day of the month as a decimal number. (Platform specific)	| 30|
	'd': func(val string, t *_DateTime) (e error) {
		if nil == t {
			return errors.New("Invalid time parameter!")
		}
		t.day, e = strconv.Atoi(val)
		return e
	},
	//| %b	| Month as locale’s abbreviated name.	| Sep|
	'b': func(val string, t *_DateTime) error {
		if nil == t {
			return errors.New("Invalid time parameter!")
		}
		for i, v := range shortMonthNames {
			if v == val {
				t.month = time.Month(i)
				return nil
			}
		}
		return errors.New("Month abbreviated name not match!")
	},
	//| %B	| Month as locale’s full name.	| September|
	'B': func(val string, t *_DateTime) error {
		if nil == t {
			return errors.New("Invalid time parameter!")
		}
		for i, v := range longMonthNames {
			if v == val {
				t.month = time.Month(i)
				return nil
			}
		}
		return errors.New("Month name not match!")
	},
	//| %m	| Month as a zero-padded decimal number.	| 09|
	//| %-m	| Month as a decimal number. (Platform specific)	| 9|
	'm': func(val string, t *_DateTime) (e error) {
		if nil == t {
			return errors.New("Invalid time parameter!")
		}
		if m, e := strconv.Atoi(val); nil == e {
			t.month = time.Month(m)
			return nil
		}else{
			return e
		}
	},
	//| %y	| Year without century as a zero-padded decimal number.	| 13|
	'y': func(val string, t *_DateTime) (e error) {
		if nil == t {
			return errors.New("Invalid time parameter!")
		}
		t.year, e = strconv.Atoi(val)
		if t.year < 70 {
			t.year += 2000
		}else{
			t.year += 1900
		}
		return e
	},
	//| %Y	| Year with century as a decimal number.	| 2013|
	'Y': func(val string, t *_DateTime) (e error) {
		if nil == t {
			return errors.New("Invalid time parameter!")
		}
		t.year, e = strconv.Atoi(val)
		return e
	},
	//| %H	| Hour (24-hour clock) as a zero-padded decimal number.	| 07|
	//| %-H	| Hour (24-hour clock) as a decimal number. (Platform specific)	| 7|
	'H': func(val string, t *_DateTime) (e error) {
		if nil == t {
			return errors.New("Invalid time parameter!")
		}
		t.hour, e = strconv.Atoi(val)
		return e
	},
	//| %I	| Hour (12-hour clock) as a zero-padded decimal number.	| 07|
	//| %-I	| Hour (12-hour clock) as a decimal number. (Platform specific)	| 7|
	'I': func(val string, t *_DateTime) (e error) {
		if nil == t {
			return errors.New("Invalid time parameter!")
		}
		t.hour, e = strconv.Atoi(val)
		return e
	},
	//| %p	| Locale’s equivalent of either AM or PM.	| AM|
	'p': func(val string, t *_DateTime) error {
		if nil == t {
			return errors.New("Invalid time parameter!")
		}
		switch val {
		case "PM":
			t.pm = true
		case "AM":
			t.pm = false
		}
		return nil
	},
	//| %M	| Minute as a zero-padded decimal number.	| 06|
	//| %-M	| Minute as a decimal number. (Platform specific)	| 6|
	'M': func(val string, t *_DateTime) (e error) {
		if nil == t {
			return errors.New("Invalid time parameter!")
		}
		t.min, e = strconv.Atoi(val)
		return e
	},
	//| %S	| Second as a zero-padded decimal number.	| 05|
	//| %-S	| Second as a decimal number. (Platform specific)	| 5|
	'S': func(val string, t *_DateTime) (e error) {
		if nil == t {
			return errors.New("Invalid time parameter!")
		}
		t.sec, e = strconv.Atoi(val)
		return e
	},
	//| %f	| Microsecond as a decimal number, zero-padded on the left.	| 000000|
	'f': func(val string, t *_DateTime) (e error) {
		if nil == t {
			return errors.New("Invalid time parameter!")
		}
		t.nsec, e = strconv.Atoi(val)
		t.nsec *= 1000
		return e
	},
	//| %z	| UTC offset in the form +HHMM or -HHMM (empty string if the the object is naive).	| |
	'z': func(val string, t *_DateTime) error {
		if nil == t {
			return errors.New("Invalid time parameter!")
		}
		return nil
	},
	//| %Z	| Time zone name (empty string if the object is naive).	| |
	'Z': func(val string, t *_DateTime) (e error) {
		if nil == t {
			return errors.New("Invalid time parameter!")
		}
		if t.loc,e = time.LoadLocation(val); nil != e {
			//fmt.Errorf("Failed: %s\n", e)
			return e
		}else{
			//fmt.Printf("Location: %s\n", t.loc)
			return nil
		}
	},
	//| %j	| Day of the year as a zero-padded decimal number.	| 273|
	//| %-j	| Day of the year as a decimal number. (Platform specific)	| 273|
	//'j': func(val string, t *_DateTime) error {
	//	if nil == t {
	//		return errors.New("Invalid time parameter!")
	//	}
	//	return nil
	//},
	//| %U	| Week number of the year (Sunday as the first day of the week) as a zero padded decimal number. All days in a new year preceding the first Sunday are considered to be in week 0.	| 39|
	//'U': func(val string, t *_DateTime) error {
	//	if nil == t {
	//		return errors.New("Invalid time parameter!")
	//	}
	//	return nil
	//},
	//| %W	| Week number of the year (Monday as the first day of the week) as a decimal number. All days in a new year preceding the first Monday are considered to be in week 0.	| 39|
	//'W': func(val string, t *_DateTime) error {
	//	if nil == t {
	//		return errors.New("Invalid time parameter!")
	//	}
	//	return nil
	//},
}

func buildRegexp(format string) (*regexp.Regexp, error) {
	buf := bytes.Buffer{}
	length := len(format)
	for i, j := 0, 0; i < length; i+=j+1 {
		c := format[i]
		j = 0
		//fmt.Printf("c > %c \n",c)
		if c != 0x25 || (i+1) >= length { // "%" -> 0x25
			buf.WriteByte(c)
			continue
		}
		if format[i+1] == 0x2d { // "-" -> 0x2d
			j += 1
		}
		if (i+j+1) < length {
			//fmt.Printf("format[i+j+1]> %c \n",format[i+j+1])
			if pattern, ok := input_regexes[rune(format[i+j+1])]; ok {

				buf.WriteString(pattern)
			}else{
				return nil, errors.New("Unknown Code:"+format[i+j:i+j+1])
			}
			j += 1
		}else{
			buf.WriteByte(format[i+1])
		}
	}
	return regexp.Compile(buf.String())
}

func Strptime(value string, format string) (time.Time, error) {
	re, e := buildRegexp(format)
	if nil != e {
		return time.Time{}, e
	}
	dt := &_DateTime{}
	dt.loc, _ = time.LoadLocation("UTC")
	match := re.FindStringSubmatch(value)
	if len(match) > 0 {
		for i, name := range re.SubexpNames() {
			if i != 0 {
				//fmt.Printf("Matched:(%s): %s \n", name, match[i])
				c := rune(name[0])
				if cvt_func, ok := input_converters[c]; ok {
					if e = cvt_func(match[i], dt); e != nil{
						//fmt.Errorf("Call '%s' function failed: %s \n", name, e)
						return time.Time{}, e
					}
				}else{
					return time.Time{}, errors.New("Invalid directive:"+string(c))
				}
			}
		}
	}else{
		return time.Time{}, errors.New("Can not match string with given format!")
	}

	if dt.pm && dt.hour < 12 {
		dt.hour += 12
	}

	return time.Date(dt.year, dt.month, dt.day, dt.hour, dt.min, dt.sec, dt.nsec, dt.loc), nil
}
