package support

import (
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

// GetPageTitle returns the <title> of a given page
func GetPageTitle(pageURL string) (string, error) {
	r, err := http.Get(pageURL)
	if err != nil {
		return "", err
	}

	title := ""

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" {
			title = n.FirstChild.Data
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	p, _ := html.Parse(r.Body)
	f(p)

	if title == "" {
		return "", errors.New("cannot find page title")
	}

	return title, nil
}

// ValidateURL validates an URL
func ValidateURL(pageURL string) (string, error) {
	if !govalidator.IsURL(pageURL) {
		return "", errors.New("argument is not a valid URL")
	}

	if !strings.HasPrefix(pageURL, "http://") && !strings.HasPrefix(pageURL, "https://") {
		pageURL = fmt.Sprintf("http://%s", pageURL)
	}

	return pageURL, nil
}
