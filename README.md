Package timefmt [ NOT FINISHED YET ]
====================================
A package of formatting and parsing datetime for golang which follows Python's directives in http://strftime.org/ .

[![Build Status](https://travis-ci.org/archsh/timefmt.svg?branch=master)](https://travis-ci.org/archsh/timefmt)

Index
-----
* Check
    `func Check(format string) error`
    Check if the given format string is valid.

* Strftime
    `func Strftime(t time.Time, format string) (string, error)`
    Return formatted time in string.

* Strptime
    `func Strptime(value string, format string) (time.Time, error)`
    Parse given string into time.

Example
-------
        package main
        import (
            "fmt"
            "time"
            "github.com/archsh/timefmt"
        )
        
        func main() {
            
        }


Stories
-------
See http://fuckinggodateformat.com/