package protos

import (
	"regexp"
)

const (
	maxTagLenght = 255
	maxTagCount  = 200
)

var regexTag = regexp.MustCompile("^[0-9a-zA-Z-_./:]*$")

func IsValidTag(input string) bool {
	if !regexTag.MatchString(input) {
		return false
	}

	if len(input) > maxTagLenght {
		return false
	}

	return len(input) > 0
}

func IsValidTags(input []string) bool {
	if len(input) == 0 {
		return true
	}

	if len(input) > maxTagCount {
		return false
	}

	return true
}
