package json_errors

import "strings"

const newline = "\\n"

var replacer = strings.NewReplacer(
	"\"", "\\\"",
	"\t", "\\t",
	"\r\n", newline,
	"\r", newline,
	"\n", newline,
)

// escapeJSON escapes the characters that are reserved in JSON.
func escapeJSON(s string) string {
	return replacer.Replace(s)
}
