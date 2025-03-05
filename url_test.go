package gourl

import (
	"testing"
)

// Valid urls to test connection and url validation separately.
var validUrls = []string{"https://www.google.de", "https://www.heise.de"}

// Unvalid urls to test the parsing of urls.
var unvalUrls = []string{"www.google.de", "www.heise.de"}

func TestIsValidUrl(t *testing.T) {

	for uri, err := range validUrls {
		if IsValidUrl(validUrls[uri]) == false {
			t.Errorf("Non valid url %s", err)
		}
	}
}

func TestTestUrl(t *testing.T) {

	for uri, err := range validUrls {
		if TestUrl(validUrls[uri]).StatusCode != 200 {
			t.Error(err)
		}
	}
}

func TestParseToUrl(t *testing.T) {
	for uri, _ := range unvalUrls {
		// Parse url if not in correct format.
		// Than test connection.
		TestUrl(ParseToUrl(unvalUrls[uri]).String())
	}
}
