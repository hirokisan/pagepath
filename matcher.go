package pagepath

// Matcher : Check the identity between Pages.
type Matcher struct{}

// Compare : For a given Elements, compares whether it matches the pages
func (m *Matcher) Compare(
	target Page,
	comparison Page,
	elements ...Element,
) bool {
	for _, element := range elements {
		if target.Extract(element) != comparison.Extract(element) {
			return false
		}
	}
	return true
}

// NewMatcher : Generate Matcher
func NewMatcher() *Matcher {
	return &Matcher{}
}
