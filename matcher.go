package pagepath

import "strings"

// Matcher : Check the identity between Pages.
type Matcher struct {
	ignoreTrailingSlash bool

	prepareFuncs []func(p *Page)
}

// Compare : For a given Elements, compares whether it matches the pages
func (m *Matcher) Compare(
	target Page,
	comparison Page,
	elements ...Element,
) bool {
	target = m.prepare(target)
	comparison = m.prepare(comparison)

	for _, element := range elements {
		if target.Extract(element) != comparison.Extract(element) {
			return false
		}
	}
	return true
}

// prepare : Apply settings to page before compare
func (m *Matcher) prepare(p Page) Page {
	if m.ignoreTrailingSlash {
		p.Path = strings.TrimSuffix(p.Path, "/")
	}
	for _, prepareFunc := range m.prepareFuncs {
		prepareFunc(&p)
	}
	return p
}

// NewMatcher : Generate Matcher
func NewMatcher(options ...func(m *Matcher)) *Matcher {
	m := &Matcher{}
	for _, opt := range options {
		opt(m)
	}
	return m
}
