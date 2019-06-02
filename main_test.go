package shortener

import "testing"

func TestShortener(t *testing.T) {
	shortener := New()
	str1 := "some-long-text"
	str2 := "another-long-text"

	hash1 := shortener.Shorten(str1)

	if hash1 == "" {
		t.Fatal("Error generating short link for \"" + str1 + "\"")
	}

	hash2 := shortener.Shorten(str2)

	if hash2 == "" {
		t.Fatal("Error generating short link for \"" + str2 + "\"")
	}



}
