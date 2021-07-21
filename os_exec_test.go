package goutils

import (
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestExec(t *testing.T) {
	out, err := ExecShell("go build")
	if err == nil {
		return
	}
	lines := strings.Split(out, "\n")
	for _, v := range lines {
		r := regexp.MustCompile(`.*\.go:[0-9]*:[0-9]*`)
		matchs := r.FindStringSubmatch(v)
		for _, s := range matchs {
			GetBetweenStr(s, "XXX", ":")
			strconv.Atoi(GetBetweenStr_FirstLast(s, ":", ":"))
		}
	}
}
