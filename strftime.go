package timefmt

import (
    "time"
    "fmt"
    "bytes"
    "errors"
)

//| %a	| Weekday as locale’s abbreviated name.	| Mon|
func cvt_output_a(t time.Time, flags bool) (string, error) {
    return shortDayNames[t.Weekday()], nil
}

//| %A	| Weekday as locale’s full name.	| Monday|
func cvt_output_A(t time.Time, flags bool) (string, error) {
    return longDayNames[t.Weekday()], nil
}

//| %w	| Weekday as a decimal number, where 0 is Sunday and 6 is Saturday.	| 1|
func cvt_output_w(t time.Time, flags bool) (string, error) {
    return fmt.Sprintf("%d", t.Weekday()), nil
}

//| %d	| Day of the month as a zero-padded decimal number.	| 30|
//| %-d	| Day of the month as a decimal number. (Platform specific)	| 30|
func cvt_output_d(t time.Time, flags bool) (string, error) {
    if flags {
        return fmt.Sprintf("%d", t.Day()), nil
    } else {
        return fmt.Sprintf("%02d", t.Day()), nil
    }
}

//| %b	| Month as locale’s abbreviated name.	| Sep|
func cvt_output_b(t time.Time, flags bool) (string, error) {
    return shortMonthNames[t.Month()], nil
}

//| %B	| Month as locale’s full name.	| September|
func cvt_output_B(t time.Time, flags bool) (string, error) {
    return longMonthNames[t.Month()], nil
}

//| %m	| Month as a zero-padded decimal number.	| 09|
//| %-m	| Month as a decimal number. (Platform specific)	| 9|
func cvt_output_m(t time.Time, flags bool) (string, error) {
    if flags {
        return fmt.Sprintf("%d", t.Month()), nil
    } else {
        return fmt.Sprintf("%02d", t.Month()), nil
    }
}

//| %y	| Year without century as a zero-padded decimal number.	| 13|
func cvt_output_y(t time.Time, flags bool) (string, error) {
    return fmt.Sprintf("%02d", t.Year()%100), nil
}

//| %Y	| Year with century as a decimal number.	| 2013|
func cvt_output_Y(t time.Time, flags bool) (string, error) {
    return fmt.Sprintf("%d", t.Year()), nil
}

//| %H	| Hour (24-hour clock) as a zero-padded decimal number.	| 07|
//| %-H	| Hour (24-hour clock) as a decimal number. (Platform specific)	| 7|
func cvt_output_H(t time.Time, flags bool) (string, error) {
    if flags {
        return fmt.Sprintf("%d", t.Hour()), nil
    } else {
        return fmt.Sprintf("%02d", t.Hour()), nil
    }
}

//| %I	| Hour (12-hour clock) as a zero-padded decimal number.	| 07|
//| %-I	| Hour (12-hour clock) as a decimal number. (Platform specific)	| 7|
func cvt_output_I(t time.Time, flags bool) (string, error) {
    if flags {
        return fmt.Sprintf("%d", t.Hour()%12), nil
    } else {
        return fmt.Sprintf("%02d", t.Hour()%12), nil
    }
}

//| %p	| Locale’s equivalent of either AM or PM.	| AM|
func cvt_output_p(t time.Time, flags bool) (string, error) {
    if t.Hour() > 12 {
        return "PM", nil
    } else {
        return "AM", nil
    }
}

//| %M	| Minute as a zero-padded decimal number.	| 06|
//| %-M	| Minute as a decimal number. (Platform specific)	| 6|
func cvt_output_M(t time.Time, flags bool) (string, error) {
    if flags {
        return fmt.Sprintf("%d", t.Minute()), nil
    } else {
        return fmt.Sprintf("%02d", t.Minute()), nil
    }
}

//| %S	| Second as a zero-padded decimal number.	| 05|
//| %-S	| Second as a decimal number. (Platform specific)	| 5|
func cvt_output_S(t time.Time, flags bool) (string, error) {
    if flags {
        return fmt.Sprintf("%d", t.Second()), nil
    } else {
        return fmt.Sprintf("%02d", t.Second()), nil
    }
}

//| %f	| Microsecond as a decimal number, zero-padded on the left.	| 000000|
func cvt_output_f(t time.Time, flags bool) (string, error) {
    return fmt.Sprintf("%06d", t.Nanosecond()/1000), nil
}

//| %z	| UTC offset in the form +HHMM or -HHMM (empty string if the the object is naive).	| |
func cvt_output_z(t time.Time, flags bool) (string, error) {
    _, o := t.Zone()
    var pfx string
    if o >= 0 {
        pfx = "+"
    } else {
        pfx = "-"
        o = 0 - o
    }
    return fmt.Sprintf("%s%02d%02d", pfx, o/3600, o%60), nil
}

//| %Z	| Time zone name (empty string if the object is naive).	| |
func cvt_output_Z(t time.Time, flags bool) (string, error) {
    s, _ := t.Zone()
    return s, nil
}

//| %j	| Day of the year as a zero-padded decimal number.	| 273|
//| %-j	| Day of the year as a decimal number. (Platform specific)	| 273|
func cvt_output_j(t time.Time, flags bool) (string, error) {
    if flags {
        return fmt.Sprintf("%d", t.YearDay()), nil
    } else {
        return fmt.Sprintf("%03d", t.YearDay()), nil
    }
}

//| %U	| Week number of the year (Sunday as the first day of the week) as a zero padded decimal number. All days in a new year preceding the first Sunday are considered to be in week 0.	| 39|
func cvt_output_U(t time.Time, flags bool) (string, error) {
    _, w := t.ISOWeek() //TODO: Need update.
    return fmt.Sprintf("%02d", w), nil
}

//| %W	| Week number of the year (Monday as the first day of the week) as a decimal number. All days in a new year preceding the first Monday are considered to be in week 0.	| 39|
func cvt_output_W(t time.Time, flags bool) (string, error) {
    _, w := t.ISOWeek()
    return fmt.Sprintf("%02d", w), nil
}

//| %c	| Locale’s appropriate date and time representation.	| Mon Sep 30 07:06:05 2013|
func cvt_output_c(t time.Time, flags bool) (string, error) {
    a, _ := cvt_output_a(t, flags)
    b, _ := cvt_output_b(t, flags)
    d, _ := cvt_output_d(t, flags)
    H, _ := cvt_output_H(t, flags)
    M, _ := cvt_output_M(t, flags)
    S, _ := cvt_output_S(t, flags)
    Y, _ := cvt_output_Y(t, flags)
    return fmt.Sprintf("%s %s %s %s:%s:%s %s", a, b, d, H, M, S, Y), nil
}

//| %x	| Locale’s appropriate date representation.	| 09/30/13|
func cvt_output_x(t time.Time, flags bool) (string, error) {
    m, _ := cvt_output_m(t, flags)
    d, _ := cvt_output_d(t, flags)
    y, _ := cvt_output_y(t, flags)
    return fmt.Sprintf("%s/%s/%s", m, d, y), nil
}

//| %X	| Locale’s appropriate time representation.	| 07:06:05|
func cvt_output_X(t time.Time, flags bool) (string, error) {
    H, _ := cvt_output_H(t, flags)
    M, _ := cvt_output_M(t, flags)
    S, _ := cvt_output_S(t, flags)
    return fmt.Sprintf("%s:%s:%s", H, M, S), nil
}

//| %%	| A literal '%' character.	| %|
func cvt_output_percent(t time.Time, flags bool) (string, error) {
    return "%", nil
}

var ontput_converters = map[rune]func(time.Time, bool) (string, error){
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
    for i, j := 0, 0; i < length; i += j + 1 {
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
        if (i + j + 1) < length {
            //fmt.Printf("format[i+j+1]> %c \n",format[i+j+1])
            if cvt_func, ok := ontput_converters[rune(format[i+j+1])]; ok {
                s, e := cvt_func(t, flag)
                if e != nil {
                    return "", e
                }
                buf.WriteString(s)
            } else {
                return "", errors.New("Unknown Code:" + format[i+j:i+j+1])
            }
            j += 1
        } else {
            buf.WriteByte(format[i+1])
        }
    }
    return buf.String(), nil
}
