package protos

import "regexp"

const (
	maxAliasLenght = 255
	maxAliasCount  = 200
)

var regexAlias = regexp.MustCompile(`^(agent|group|system|user)\/[0-9a-z-_.]{1,255}\/[0-9a-z-_.\/]{1,255}$`)

func IsValidAlias(input string) bool {
	if !regexAlias.MatchString(input) {
		return false
	}

	if len(input) > maxAliasLenght {
		return false
	}

	return len(input) > 0
}

func IsValidAliases(input map[string]SecretAlias) bool {
	if len(input) == 0 {
		return true
	}

	if len(input) > maxAliasCount {
		return false
	}

	return true
}
