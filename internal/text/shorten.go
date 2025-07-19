package text

// Return first n characters of text
func Shorten(text string, n int) string {
	if len(text) < n {
		return text
	}
	return text[:n]
}
