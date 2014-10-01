package template

import (
	"testing"
)

func TestChangeScheme(t *testing.T) {
	urls := []string{
		"tcp://localhost:1230/",
		"http://127.0.0.1:194/",
	}
	newScheme := []string{
		"http",
		"tcp",
	}
	expected := []string{
		"http://localhost:1230/",
		"tcp://127.0.0.1:194/",
	}

	for i := range urls {
		result := changeScheme(newScheme[i], urls[i])
		if result != expected[i] {
			t.Errorf("Expected %s, got %s", expected[i], result)
		}
	}
}
