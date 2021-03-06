package hyperlinker_test

import (
	"golang.org/x/net/html"
	"testing"
	"github.com/nottheswimmer/hyperlink-paste/internal/hyperlinker"

	"github.com/stretchr/testify/assert"
)

func TestGetURLHyperlinkHTML(t *testing.T) {
	cases := []struct {
		url   string
		title string
	}{
		{"https://github.com/", "GitHub: Where the world builds software · GitHub"},
		{"https://xxx.atlassian.net/browse/DC-2712", "DC-2712"},
	}

	// TODO: Mock requests
	for _, testCase := range cases {
		t.Run(testCase.url, func(t *testing.T) {
			assert.Equal(t, "<a href=\"" + testCase.url + "\">" + html.EscapeString(testCase.title) + "</a>",
				string(hyperlinker.GetURLHyperlinkHTML(testCase.url)))
		})
	}
}
