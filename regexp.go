package main

import "regexp"

// RegDelete DeleteRegExp
func RegDelete(data, rule string) string {
	reg := regexp.MustCompile(rule)
	result := reg.ReplaceAllString(data, "")
	return result
}

// RegMatch MatchRegExp
func RegMatch(data, rule string) string {
	reg := regexp.MustCompile(rule)
	result := reg.FindString(data)
	return result
}

// RegMatchArray RegMatchArray
func RegMatchArray(data, rule string) []string {
	reg := regexp.MustCompile(rule)
	result := reg.FindAllString(data, -1)
	return result
}

// RegReplace ReplaceRegExp
func RegReplace(source, after, rule string) string {
	reg := regexp.MustCompile(rule)
	result := reg.ReplaceAllString(source, after)
	return result
}

// ParseBool String to Bool
func ParseBool(str string) bool {
	switch str {
	case "1", "t", "T", "true", "TRUE", "True":
		return true
	case "0", "f", "F", "false", "FALSE", "False":
		return false
	}
	return false
}
