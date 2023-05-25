package pagepath

import (
	"fmt"
	"net/url"
)

// Element : Components of page path
type Element string

const (
	// ElementScheme : scheme
	// e.g. http, https, ...
	ElementScheme = Element("scheme")
	// ElementHost : host
	// e.g. example.com, www.example.com, ...
	ElementHost = Element("host")
	// ElementPath : path
	// e.g. /aaa/bbb, /aaa/bbb/, ...
	ElementPath = Element("path")
	// ElementQuery : query
	// e.g. ?a=aa&b=bb, ...
	ElementQuery = Element("query")
	// ElementFragment : fragment
	// e.g. #xxx, ...
	ElementFragment = Element("fragment")
)

// Elements : Reflects the sort order
func Elements() []Element {
	return []Element{
		ElementScheme,
		ElementHost,
		ElementPath,
		ElementQuery,
		ElementFragment,
	}
}

// Page : Page that holds the Elements
type Page struct {
	Scheme   string
	Host     string
	Path     string
	Query    string
	Fragment string
}

// ComparePageWithElements : For a given Elements, compares whether it matches the given Page
func (p *Page) ComparePageWithElements(
	obj Page,
	elements ...Element,
) bool {
	for _, element := range elements {
		if p.Extract(element) != obj.Extract(element) {
			return false
		}
	}
	return true
}

func (p *Page) String() string {
	if p == nil {
		return ""
	}
	u := url.URL{
		Scheme:   p.Scheme,
		Host:     p.Host,
		Path:     p.Path,
		RawQuery: p.Query,
		Fragment: p.Fragment,
	}
	return u.String()
}

// Extract : Extracts the specified element
func (p *Page) Extract(element Element) string {
	switch element {
	case ElementScheme:
		return p.Scheme
	case ElementHost:
		return p.Host
	case ElementPath:
		return p.Path
	case ElementQuery:
		return p.Query
	case ElementFragment:
		return p.Fragment
	default:
		panic(fmt.Errorf("%s is not supported element", element))
	}
}

// New :
func New(pagePath string) (Page, error) {
	u, err := url.Parse(pagePath)
	if err != nil {
		return Page{}, fmt.Errorf("parse: %w", err)
	}
	return Page{
		Scheme:   u.Scheme,
		Host:     u.Host,
		Path:     u.Path,
		Query:    u.RawQuery,
		Fragment: u.Fragment,
	}, nil
}

// Extract : Extracts the specified element from page path
func Extract(
	pagePath string,
	element Element,
) (string, error) {
	page, err := New(pagePath)
	if err != nil {
		return "", err
	}
	return page.Extract(element), nil
}
