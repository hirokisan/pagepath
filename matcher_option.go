package pagepath

// WithIgnoreTrailingSlash :
func WithIgnoreTrailingSlash() func(m *Matcher) {
	return func(m *Matcher) {
		m.ignoreTrailingSlash = true
	}
}

// WithPrepareFuncs :
func WithPrepareFuncs(fs ...func(p *Page)) func(m *Matcher) {
	return func(m *Matcher) {
		m.prepareFuncs = fs
	}
}
