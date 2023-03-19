package util

import "regexp"

var int_m *regexp.Regexp = regexp.MustCompile(`^-?\d+$`)
var flt_m *regexp.Regexp = regexp.MustCompile(`^-?\d+.\d+$`)
var str_m *regexp.Regexp = regexp.MustCompile(`^\".*\"`)

func IsInteger(v string) bool {
	return int_m.MatchString(v)
}

func IsFloat(v string) bool {
	return flt_m.MatchString(v)
}

func IsString(v string) bool {
	return str_m.MatchString(v)
}
