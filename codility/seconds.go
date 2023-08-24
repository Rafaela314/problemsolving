package codility

import "fmt"

func GetTime(seconds int) string {

	// min => 60
	// hour => 3600
	// day => 86400

	if seconds == 0 {
		return "now"
	}

	if seconds < 60 {
		return fmt.Sprintf("%d seconds", seconds)
	}

	if seconds >= 60 && seconds < 3600 {

		min := seconds / 60
		sec := seconds - (min * 60)

		return fmt.Sprintf("%d minutes and %d seconds", min, sec)
	}

	if seconds >= 3600 && seconds < 86400 {

		hr := seconds / 3600
		min := (seconds - hr*3600) / 60
		sec := seconds - (hr * 3600) - (min * 60)

		return fmt.Sprintf("%d hour %d minutes and %d seconds", hr, min, sec)
	}

	day := seconds / 86400
	hr := (seconds - (day * 86400)) / 3600
	min := (seconds - (day * 86400) - (hr * 3600)) / 60
	sec := seconds - (day * 86400) - (hr * 3600) - (min * 60)

	return fmt.Sprintf("%d days %d hour %d minutes and %d seconds", day, hr, min, sec)
}
