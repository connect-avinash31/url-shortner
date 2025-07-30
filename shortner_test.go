package main

// this file will be used to test  url shortner functionlaity
import "testing"

func TestUrlShortner(t *testing.T) {
	// first create url shortner services
	testUrlShortner := NewUrlShortner()

	// now i will be end to end test for url shortner
	// first i will create a url
	originalUrl := "www.google.com/search/infra-cloud/careers"
	// now i will call ShortenValue function to shorten this url
	shortenedUrl, err := testUrlShortner.ShortenValue(originalUrl)
	if err != nil {
		t.Fatalf("failed to shorten URL: %v", err)
	}
	// now check if shortened url is not empty
	if len(shortenedUrl) == 0 {
		t.Fatal("shortened URL is empty")
	}
	// now i will call OriginalValue function to get original url from shortened url
	originalValue, err := testUrlShortner.OriginalValue(shortenedUrl)
	if err != nil {
		t.Fatalf("failed to get original URL: %v", err)
	}
	// now check if original value is same as original url
	if originalValue != originalUrl {
		t.Fatalf("original URL does not match: expected %s, got %s", originalUrl, originalValue)
	}
	// now again i will recall shortner url , it should give the same value else it will failt
	duplicateShortenerUrl, err := testUrlShortner.ShortenValue(originalUrl)
	if err != nil {
		t.Fatalf("failed to shorten URL again: %v", err)
	}
	// now check ealrier shortner url and new shortner url
	if duplicateShortenerUrl != shortenedUrl {
		t.Fatalf("duplicate shortened URL does not match: expected %s, got %s", shortenedUrl, duplicateShortenerUrl)
	}
	// now i will check matrics
	matrics, err := testUrlShortner.Metrics()
	if err != nil {
		t.Fatalf("failed to get metrics: %v", err)
	}
	// now check if matrics is not empty
	if len(matrics) == 0 {
		t.Fatal("metrics are empty")
	}
	// lenght should be one
	if len(matrics) != 1 {
		t.Fatalf("expected 1 metric, got %d", len(matrics))
	}

}
