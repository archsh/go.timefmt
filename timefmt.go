package timefmt

import "time"

func Strftime(t time.Time, format string) (string, error) {

	return "", nil
}

func Strptime(value string, format string) (time.Time, error) {

	return time.Time{}, nil
}

func Check(format string) error {
	return nil
}