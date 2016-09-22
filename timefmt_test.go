package timefmt

import (
	"testing"
	"time"
)

var formats = []string{
	"",
	"",
	"",
	"",
	"",
	"",
}
var invalid_formats = []string{
	"",
	"",
	"",
	"",
}

func TestCheck(t *testing.T) {
	for _, f := range formats {
		if nil != Check(f) {
			t.Errorf("'%s' Should be valid!\n", f)
		}
	}
	for _, f := range invalid_formats {
		if nil == Check(f) {
			t.Errorf("'%s' Should be invalid!\n", f)
		}
	}
}

func TestStrftime(t *testing.T) {
	var validate = func (tm time.Time, format string, result string) {
		if s, e := Strftime(tm, format); e != nil || s != result {
			t.Errorf("Strftime(/%v/, '%s') should return '%s' but not (%s) (%s)\n", tm, format, result, e, s)
		}
	}
	tm := time.Unix(1474524266, 321)
	validate(tm, "%Y-%m-%dT%H%M%S","/2016-09-22T14:04:26")
}

func TestStrptime(t *testing.T) {

	var validate = func (val string, format string, result time.Time) {
		if tm, e := Strptime(val, format); e != nil || tm != result {
			t.Errorf("Strptime('%s', '%s') should return /%v/ but not (%s) (%v)\n", val, format, result, e, tm)
		}
	}
	tm := time.Unix(1474524266, 321)
	validate("2016-09-22T14:04:26.000000321", "%Y-%m-%dT%H%M%S.%f", tm)

}

func BenchmarkStrftime(b *testing.B) {

}

func BenchmarkStrptime(b *testing.B) {

}