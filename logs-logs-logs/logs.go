package logs

import (
	"fmt"
	"log"
	"strings"
	"unicode/utf8"
)

// Message extracts the message from the provided log line.
func Message(line string) string {
	levelEndIdx := strings.IndexRune(line, ':') + 1
	if levelEndIdx == 0 {
		log.Panicf("Message: received a malformed log message -- %s", line)
	}

	return strings.TrimSpace(line[levelEndIdx:])
}

// MessageLen counts the amount of characters (runes) in the message of the log line.
func MessageLen(line string) int {
	return utf8.RuneCountInString(Message(line))
}

// LogLevel extracts the log level string from the provided log line.
func LogLevel(line string) string {
	endIdx := strings.IndexRune(line, ':')
	return strings.TrimFunc(strings.ToLower(line[:endIdx]), func(r rune) bool {
		return r == '[' || r == ']'
	})
}

// Reformat reformats the log line in the format `message (logLevel)`.
func Reformat(line string) string {
	return fmt.Sprintf("%s (%s)", Message(line), LogLevel(line))
}
