package htmlclip_test

import (
	"github.com/atotto/clipboard"
	"testing"
	"github.com/nottheswimmer/hyperlink-paste/internal/htmlclip"

	"github.com/stretchr/testify/assert"
)

func TestClipAsHTML(t *testing.T) {
	tests := []string{"", "<a href=\"https://github.com\">Github</a>"};

	// TODO: Check if it actually got copied as HTML. Currently only actually checks if it's copied as a string.
	for _, testString := range tests {
		t.Run(testString, func(t *testing.T) {
			htmlclip.ClipAsHTML([]byte(testString))
			text, _ := clipboard.ReadAll()
			assert.Equal(t, testString, text);
		})
	}
}