package grep

import (
	"regexp"
	"bufio"
	"strings"
)

func SearchString(re, str string) (o []string, err error) {
	regex, err := regexp.Compile(re)
	if err != nil {
		return
	}

	r := bufio.NewReader(strings.NewReader(str))
	if err != nil {
		return
	}

	buf := make([]byte, 1024)
	for {
		buf, _, err = r.ReadLine()
		if err != nil {
			return
		}

		s := string(buf)
		if regex.MatchString(s) {
			o = append(o, string(buf))
		}
	}

	return
}
