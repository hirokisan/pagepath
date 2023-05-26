package pagepath

// WithIgnoreTrailingSlash :
func WithIgnoreTrailingSlash() func(m *Matcher) {
	return func(m *Matcher) {
		m.ignoreTrailingSlash = true
	}
}
