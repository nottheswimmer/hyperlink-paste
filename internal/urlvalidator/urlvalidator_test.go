package urlvalidator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWrite(t *testing.T) {
	cases := []struct {
		url   string
		valid bool
	}{
		{"https://github.com/", true},
		{"http://github.com/", true},
		{"//github.com/", true},
		{"", false},
		{"github.com", false},
		{"x://github.com", false},
		{"hello how are you", false},
		{"https://", false},
	}

	for _, testCase := range cases {
		t.Run(testCase.url, func(t *testing.T) {
			if testCase.valid {
				assert.True(t, IsValidURL(testCase.url))
			} else {
				assert.False(t, IsValidURL(testCase.url))
			}
		})
	}
}
