package strcase

import (
	"strings"
)

func SnakeCase(s string) string {
	return delimiterCase(s, "_", false)
}

func UpperSnakeCase(s string) string {
	return delimiterCase(s, "_", true)
}

func delimiterCase(s string, delimiter string, upperCase bool) string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return s
	}
	var ns string
	sl := strings.Split(s, delimiter)

	for _, v := range sl {

		nsl := len(ns)

		if len(v) == 0 {
			continue
		}

		if nsl == 0 {
			if upperCase {
				ns = strings.ToUpper(v[:1]) + v[1:]
			} else {
				ns = v
			}
		}

		if nsl > 0 {
			ns += strings.ToUpper(v[:1]) + v[1:]
		}
	}
	return ns
}
