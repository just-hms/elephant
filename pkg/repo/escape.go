package repo

import "strings"

func escape(s string) string {
	return strings.ReplaceAll(s, "\n", "\\n")
}

func deEscape(s string) string {
	return strings.ReplaceAll(s, "\\n", "\n")
}
