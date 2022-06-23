package bob

import (
	"regexp"
	"strings"
)

// Hey returns a response from Bob depending on what you say to him.
func Hey(remark string) string {
	if remark = strings.TrimSpace(remark); remark == "" {
		return "Fine. Be that way!"
	}

	reAlpha := regexp.MustCompile("[a-zA-Z]")
	hasAlpha := reAlpha.FindString(remark) != ""
	isQuestion := strings.HasSuffix(remark, "?")
	isShouting := strings.ToUpper(remark) == remark && hasAlpha

	if isShouting && isQuestion {
		return "Calm down, I know what I'm doing!"
	} else if isShouting {
		return "Whoa, chill out!"
	} else if isQuestion {
		return "Sure."
	}
	return "Whatever."
}
