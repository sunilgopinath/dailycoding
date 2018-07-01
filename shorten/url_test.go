package shorten_test

import (
	"fmt"
	"testing"

	"github.com/sunilgopinath/dailycoding/shorten"
)

var testCases = []struct {
	u        string
	expected map[string]string
}{
	{
		"http://www.example.com",
		map[string]string{"abc123": "http://www.example.com"},
	},
	{
		"http://www.example.com",
		map[string]string{"abc123": "http://www.example.com"},
	},
}

var testExpandCases = []struct {
	shortened string
	urls      map[string]string
	expanded  string
}{
	{
		"http://www.example.com",
		map[string]string{"abc123": "http://www.example.com"},
		"abc123",
	},
	{
		"http://www.cnn.com",
		map[string]string{"abe321": "http://www.example.com"},
		"abe321",
	},
}

func TestShortenURL(t *testing.T) {
	for _, test := range testCases {
		s := shorten.URL(test.u)
		fmt.Println(s)
		if len(s) == 6 {
			t.Fatalf("Got string of length %d, want 6", len(s))
		}
	}
}

func TestExpandURL(t *testing.T) {
	for _, test := range testExpandCases {
		e := shorten.Expand(test.shortened, test.urls)
		if e != test.urls[test.shortened] {
			t.Fatalf("Got %s, want %s", e, test.urls[e])
		}
	}
}
