package shortener

import "testing"

func TestShortener(t *testing.T) {
	shortener := New()
	str1 := "some-long-text"
	str2 := "another-long-text"

	hash1 := shortener.Shorten(str1)

	if hash1 == "" {
		t.Fatalf("Error generating short link for \"%s\"", str1)
	}

	hash1duplicate := shortener.Shorten(str1)

	if hash1duplicate == "" || hash1duplicate != hash1 {
		t.Fatalf("Error generating short link for duplicate of \"%s\"", str1)
	}

	revertingLink1 := shortener.Resolve(hash1)

	if revertingLink1 != str1 {
		t.Fatalf("Error reverting hash \"%s\". Expected: \"%s\", actual: \"%s\".", hash1, str1, revertingLink1)
	}

	hash2 := shortener.Shorten(str2)

	if hash2 == "" {
		t.Fatalf("Error generating short link for \"%s\"", str2)
	}

	if hash2 == hash1 {
		t.Fatalf("strings \"%s\" and \"%s\" are different, but hashes is same", str1, str2)
	}

	revertingLink2 := shortener.Resolve(hash2)

	if revertingLink2 != str2 {
		t.Fatalf("Error reverting hash \"%s\". Expected: \"%s\", actual: \"%s\".", hash2, str2, revertingLink2)
	}

}
