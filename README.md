# pagepath

`pagepath` is a Golang library for handling and matching page paths.

## Usage

Get element from page path

```go
import (
	"fmt"
	"log"

	"github.com/hirokisan/pagepath"
)

func main() {
	pagePath := "https://pkg.go.dev/net/url?xxx=bbb#Parse"

	host, err := pagepath.Extract(pagePath, pagepath.ElementHost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(host) // pkg.go.dev
}
```

Compare pages, for given elements

```go
import (
	"fmt"
	"log"

	"github.com/hirokisan/pagepath"
)

func main() {
	matcher := pagepath.NewMatcher()

	page, err := pagepath.New("https://pkg.go.dev/net/url?xxx=bbb#Parse")
	if err != nil {
		log.Fatal(err)
	}

	comparisonPage, err := pagepath.New("https://pkg.go.dev/net/url?xxx=bbb")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(matcher.Compare(page, comparisonPage, pagepath.ElementPath)) // true
}
```
