package main

import (
	"github.com/atotto/clipboard"
	"github.com/go-vgo/robotgo"
	"github.com/nottheswimmer/hyperlink-paste/internal/htmlclip"
	"github.com/nottheswimmer/hyperlink-paste/internal/hyperlinker"
	"github.com/nottheswimmer/hyperlink-paste/internal/urlvalidator"
)

func main() {

	clipURL, _ := clipboard.ReadAll()

	if urlvalidator.IsValidURL(clipURL) {
		urlHTML := hyperlinker.GetURLHyperlinkHTML(clipURL)
		htmlclip.ClipAsHTML(urlHTML)
	}
	paste()
}

func paste() string {
	return robotgo.KeyTap("v", "command")
}
