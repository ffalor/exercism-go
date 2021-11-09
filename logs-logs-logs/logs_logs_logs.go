package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {

	applications := map[rune]string{
		'❗': "recommendation",
		'🔍': "search",
		'☀': "weather",
	}

	for _, r := range log {
		value, ok := applications[r]
		if ok {
			return value
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	out := []rune(log)
	for i, r := range out {
		if r == oldRune {
			out[i] = newRune
		}
	}
	return string(out)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
