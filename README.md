# timefmt

A package of formatting and parsing datetime for golang which follows Python's directives in http://strftime.org/ .
*_It should work, but performance maybe should be improved later._*

[![Build Status](https://travis-ci.org/archsh/timefmt.svg?branch=master)](https://travis-ci.org/archsh/timefmt)

## Index

### Strftime
`func Strftime(t time.Time, format string) (string, error)`
Return formatted time in string.

### Strptime
`func Strptime(value string, format string) (time.Time, error)`
Parse given string into time.

## Example

```go
package main
import (
    "fmt"
    "time"
    "github.com/archsh/timefmt"
)

func main() {
    tm := time.Now()
    s, e := timefmt.Strftime(tm, "%Y-%m-%dT%H:%M:%S")//,"2016-09-22T06:04:26")
    fmt.Printf("%s <%s>\n",s,e)
    s, e = timefmt.Strftime(tm, "%y-%m-%dT%H:%M:%S")//,"16-09-22T06:04:26")
    fmt.Printf("%s <%s>\n",s,e)
    s, e = timefmt.Strftime(tm, "%Y-%m-%dT%I:%M:%S")//,"2016-09-22T06:04:26")
    fmt.Printf("%s <%s>\n",s,e)
    s, e = timefmt.Strftime(tm, "%Y-%m-%dT %p %I:%M:%S")//,"2016-09-22T AM 06:04:26")
    fmt.Printf("%s <%s>\n",s,e)
    s, e = timefmt.Strftime(tm, "%Y-%b-%dT%H:%M:%S")//,"2016-Sep-22T06:04:26")
    fmt.Printf("%s <%s>\n",s,e)
    s, e = timefmt.Strftime(tm, "%Y-%B-%dT%H:%M:%S")//,"2016-September-22T06:04:26")
    fmt.Printf("%s <%s>\n",s,e)
    s, e = timefmt.Strftime(tm, "%Y-%b-%dT%H:%-M:%S")//,"2016-Sep-22T06:4:26")
    fmt.Printf("%s <%s>\n",s,e)
    s, e = timefmt.Strftime(tm, "%c")//, "Thu Sep 22 06:04:26 2016")
    fmt.Printf("%s <%s>\n",s,e)
    s, e = timefmt.Strftime(tm, "%x")//, "09/22/16")
    fmt.Printf("%s <%s>\n",s,e)
    s, e = timefmt.Strftime(tm, "%X")//, "06:04:26")
    fmt.Printf("%s <%s>\n",s,e)
    s, e = timefmt.Strftime(tm, "%Y-%m-%dT%H:%M:%S %z")//,"2016-09-22T06:04:26 +0000")
    fmt.Printf("%s <%s>\n",s,e)
}

```
## Directives
-----

| Code | Meaning | Example |
|------|---------|---------|
| %a	| Weekday as locale’s abbreviated name.	| Mon| 
| %A	| Weekday as locale’s full name.	| Monday| 
| %w	| Weekday as a decimal number, where 0 is Sunday and 6 is Saturday.	| 1| 
| %d	| Day of the month as a zero-padded decimal number.	| 30| 
| %-d	| Day of the month as a decimal number. (Platform specific)	| 30| 
| %b	| Month as locale’s abbreviated name.	| Sep| 
| %B	| Month as locale’s full name.	| September| 
| %m	| Month as a zero-padded decimal number.	| 09| 
| %-m	| Month as a decimal number. (Platform specific)	| 9| 
| %y	| Year without century as a zero-padded decimal number.	| 13| 
| %Y	| Year with century as a decimal number.	| 2013| 
| %H	| Hour (24-hour clock) as a zero-padded decimal number.	| 07| 
| %-H	| Hour (24-hour clock) as a decimal number. (Platform specific)	| 7| 
| %I	| Hour (12-hour clock) as a zero-padded decimal number.	| 07| 
| %-I	| Hour (12-hour clock) as a decimal number. (Platform specific)	| 7| 
| %p	| Locale’s equivalent of either AM or PM.	| AM| 
| %M	| Minute as a zero-padded decimal number.	| 06| 
| %-M	| Minute as a decimal number. (Platform specific)	| 6| 
| %S	| Second as a zero-padded decimal number.	| 05| 
| %-S	| Second as a decimal number. (Platform specific)	| 5| 
| %f	| Microsecond as a decimal number, zero-padded on the left.	| 000000| 
| %z	| UTC offset in the form +HHMM or -HHMM (empty string if the the object is naive).	| | 
| %Z	| Time zone name (empty string if the object is naive).	| | 
| %j	| Day of the year as a zero-padded decimal number.	| 273| 
| %-j	| Day of the year as a decimal number. (Platform specific)	| 273| 
| %U	| Week number of the year (Sunday as the first day of the week) as a zero padded decimal number. All days in a new year preceding the first Sunday are considered to be in week 0.	| 39| 
| %W	| Week number of the year (Monday as the first day of the week) as a decimal number. All days in a new year preceding the first Monday are considered to be in week 0.	| 39| 
| %c	| Locale’s appropriate date and time representation.	| Mon Sep 30 07:06:05 2013| 
| %x	| Locale’s appropriate date representation.	| 09/30/13| 
| %X	| Locale’s appropriate time representation.	| 07:06:05| 
| %%	| A literal '%' character.	| %| 

## Note

### Not supported codes for Strptime()
The following codes was not supported because it does not make sense:
- `%a`
- `%A` 
- `%w` 
- `%j` 
- `%-j` 
- `%U` 
- `%W`

### Not ready yet for Strptime()
- `%z`