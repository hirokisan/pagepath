package pagepath

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
