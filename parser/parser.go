package parser

import (
	"strings"
)

// ParseRequest parses the request string. To use, remove the `!socrates` prefix
func ParseRequest(request string) (command string, arguments []string) {
	parts := strings.Split(request, " ")
	if len(parts) == 0 {
		return "", nil
	}
	if len(parts) == 1 {
		return parts[0], nil
	}
	var args []string
	for _, arg := range parts[1:] {
		if arg != "" {
			args = append(args, arg)
		}
	}
	return parts[0], args
}
