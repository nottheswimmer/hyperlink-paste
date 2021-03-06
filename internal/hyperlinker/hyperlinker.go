package hyperlinker

import (
	"golang.org/x/net/html"
	"io"
	"net/http"
	"strings"
)

// GetURLHyperlinkHTML takes a URL and returns HTML with a hyperlink that includes a suitable title
func GetURLHyperlinkHTML(clipURL string) []byte {
	name := getURLTitle(clipURL)
	return []byte("<a href=\"" + clipURL + "\">" + html.EscapeString(name) + "</a>")
}

func getURLTitle(clipURL string) string {
	// If there is a redirect, treat that as an error.
	// e.g. "DC-2717" is more meaningful than "Log in with Atlassian account"
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}
	resp, err := client.Get(clipURL)
	name := getLastPartOfURL(clipURL)
	if err == nil && resp != nil {
		defer resp.Body.Close()
		if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
			if title, ok := getHTMLTitle(resp.Body); ok {
				name = title
			}
		}
	}
	return name
}

func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

func getLastPartOfURL(text string) string {
	ss := strings.Split(text, "/")
	return ss[len(ss)-1]
}


func traverse(n *html.Node) (string, bool) {
	if isTitleElement(n) {
		return n.FirstChild.Data, true
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result, ok := traverse(c)
		if ok {
			return result, ok
		}
	}

	return "", false
}

// getHTMLTitle gets the HTML title for a response
func getHTMLTitle(r io.Reader) (string, bool) {
	doc, err := html.Parse(r)
	if err != nil {
		panic("Fail to parse html")
	}
	return traverse(doc)
}

