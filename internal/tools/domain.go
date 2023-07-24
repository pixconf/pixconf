package tools

import "strings"

func DomainNameGuesses(name string) []string {
	var names []string
	var nameSplited string

	names = append(names, name)
	nameSplited = name

	for strings.Contains(nameSplited, ".") {
		result := strings.SplitN(nameSplited, ".", 2)

		nameSplited = result[1]

		if strings.Contains(result[1], ".") && len(result[1]) > 1 {
			names = append(names, result[1])
		}
	}

	return names
}
