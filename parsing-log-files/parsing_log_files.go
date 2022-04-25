package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

type logTag string

var tags = [...]logTag{
	"[TRC]",
	"[DBG]",
	"[INF]",
	"[WRN]",
	"[ERR]",
	"[FTL]",
}

// IsValidLine checks if given text is from logs.
func IsValidLine(text string) bool {
	var valid bool
	for _, t := range tags {
		if strings.HasPrefix(text, string(t)) {
			valid = true
			break
		}
	}
	return valid
}

// SplitLogLine returns sections of a log line delimited by `<` to `>`
func SplitLogLine(text string) []string {
	return regexp.MustCompile(`<[\*~=-]*>`).Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	target := regexp.MustCompile(`(?i:".*password.*")`)
	cnt := 0
	for _, l := range lines {
		if target.MatchString(l) {
			cnt++
		}
	}
	return cnt
}

func RemoveEndOfLineText(text string) string {
	return regexp.MustCompile(`end-of-line\d+`).ReplaceAllString(text, "")
}

const user = "User "

func TagWithUserName(lines []string) []string {
	res := make([]string, len(lines))
	for i, l := range lines {
		if idx := strings.Index(l, user) + len(user); idx != -1 + len(user){
			usr := strings.Fields(l[idx:])[0]
			res[i] = fmt.Sprintf("[USR] %s %s", usr, l)
		} else {
			res[i] = l
		}
	}
	return res
}
