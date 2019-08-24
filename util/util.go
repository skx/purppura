package util

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

//
// Str2Unix converts a "time string" into an absolute timestamp.
//
// For example "+5m", "+6H", "now", or "clear"
//
func Str2Unix(data string) (int64, error) {

	now := time.Now().Unix()

	if data == "clear" {
		return 0, nil
	}

	if data == "now" {
		return now, nil
	}

	re := regexp.MustCompile(`^\\+?([0-9]+)([hHmMsS])$`)
	out := re.FindStringSubmatch(data)

	res := now

	if len(out) > 1 {
		number, err := strconv.ParseInt(out[1], 10, 64)

		if err != nil {
			return 0, err
		}
		period := out[2]

		switch period {

		case "h":
			res = now + (number * 60 * 60)
		case "H":
			res = now + (number * 60 * 60)
		case "m":
			res = now + (number * 60)
		case "M":
			res = now + (number * 60)
		case "s":
			res = now + (number * 1)
		case "S":
			res = now + (number * 1)

		}

		return res, nil
	}

	return 0, errors.New("Failed to parse string: " + data)
}
