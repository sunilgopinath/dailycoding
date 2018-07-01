package shorten

import (
	"github.com/sunilgopinath/dailycoding/rand"
)

var short = make(map[string]string)
var urls = make(map[string]string)

// URL takes in URL and returns a length 6 alphanumeric string
func URL(u string) string {
	if val, ok := urls[u]; ok {
		return val
	}
	s := generateRandom(6, short)
	short[u] = s
	return s
}

// Expand takes in a shortened string and returns a URL
func Expand(s string, m map[string]string) string {
	if val, ok := m[s]; ok {
		return val
	}
	return ""
}

func generateRandom(l int, m map[string]string) string {
	exists := true
	s := rand.String(l)
	for exists {
		if _, ok := m[s]; !ok {
			exists = false
		} else {
			s = rand.String(l)
		}
	}
	return s
}
