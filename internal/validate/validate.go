package validate

import (
	"regexp"
	"strconv"
)

func Date(flow ...string) bool {
	var ok = true
	re := regexp.MustCompile("(((19|20)\\d\\d)/(0?[1-9]|1[012])/0?[1-9]|[12][0-9]|3[01])")

	for _, s := range flow {
		if s != "" {
			ok = re.MatchString(s)
		}
	}

	return ok
}

func Price(flow  ...string) error {
	var err error
	for _, s := range flow {
		if s != "" {
			_, err = strconv.ParseFloat(s, 32)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
