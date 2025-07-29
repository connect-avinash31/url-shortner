package main

import (
	"fmt"
	"strings"
)

// So will be writing a Shortner Service

type ShortnerService interface {
	ShortenValue(originalValue string) (string, error)
	OriginalValue(shortenedValue string) (string, error)
	Metrics(shortenedValue string) (map[string]int, error)
}

type Mapper struct {
	// this mapper will conatins 2 maps one for original to shortened and one for shortened to original
	originalToShortened map[string]string
	shortenedToOriginal map[string]string
}

type UrlShortner struct {
	// this will conatins 2 things one is webSiteMapperMap and other is Hasher
	webSiteMapperMap map[string]*Mapper
	hasher           Hasher
}

// NewUrlShortner creates a new UrlShortner instacnce with hasher of it's own and webSiteMapper
func NewUrlShortner(hasher Hasher) *UrlShortner {
	return &UrlShortner{
		webSiteMapperMap: make(map[string]*Mapper),
		hasher:           hasher,
	}
}

func (urlShtnr *UrlShortner) ShortenValue(originalValue string) (string, error) {
	// now first i need to find which is start website , and which is extended key
	// so my website will be www.udemy.com/courses/ai-course so i will split using /
	// and use first array element as website and rest as shortning path
	parts := strings.Split(originalValue, "/")
	if len(parts) < 2 {
		return "", fmt.Errorf("invalid URL format : %s", originalValue)
	}
	// now first parte is website
	website := parts[0]
	// now i will check if mapper present for this website or not
	websiteMapper, exists := urlShtnr.webSiteMapperMap[website]
	if !exists {
		// then create new Mapper for this website
		websiteMapper = &Mapper{
			originalToShortened: make(map[string]string),
			shortenedToOriginal: make(map[string]string),
		}
		// set this mapper to webSiteMapperMap
		urlShtnr.webSiteMapperMap[website] = websiteMapper
	}
	// now we will check if original value present in original Mapper if it's present
	// then we can use it and return shortned value
	if shortenedValue, exists := websiteMapper.originalToShortened[originalValue]; exists {
		return shortenedValue, nil
	}
	// if not exists then now call the hasher and create the hasher
	shortenedValue, err := urlShtnr.hasher.Hash(originalValue)
	// if hasher throws an error then we will return the error
	if err != nil {
		return "", fmt.Errorf("error hashing original value: %w", err)
	}
	// now w can set current shortenv value to Mappers
	websiteMapper.originalToShortened[originalValue] = shortenedValue
	websiteMapper.shortenedToOriginal[shortenedValue] = originalValue
	// while replayiong back we will create whole url with shortner and send back
	shortnedUrl := website + "/" + shortenedValue
	return shortnedUrl, nil

}

type Hasher interface {
	Hash(value string) (string, error)
}

type URLHasher struct {
	// hashign alogorithm
}
