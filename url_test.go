package gourl

import "testing"

var urls = []string{"https://www.google.de"}

func Test_isValid(t *testing.T) {

	for uri, err := range urls {
		if IsValidUrl(urls[uri]) == false {
			t.Errorf("Non valid url %s", err)
		}
	}
}

func Test_testUrl(t *testing.T) {
	for uri, err := range urls {
		if TestUrl(urls[uri]).StatusCode != 200 {
			t.Error(err)
		}
	}
}
