package main

import (
	"fmt"
	"sort"
	"strings"
)

// So will be writing a Shortner Service

type ShortnerService interface {
	ShortenValue(originalValue string) (string, error)
	OriginalValue(shortenedValue string) (string, error)
	Metrics() (map[string]int, error)
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
		return website + "/" + shortenedValue, nil
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

func (urlShtnr *UrlShortner) OriginalValue(shortenedValue string) (string, error) {
	// first we need to find which website this shortened value belongs to
	parts := strings.Split(shortenedValue, "/")
	// if url not devided into 2 parts then it's not shortned url
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid shortened URL format : %s", shortenedValue)
	}
	website := parts[0]
	// now we will check if mapper present for this website or not
	websiteMapper, exists := urlShtnr.webSiteMapperMap[website]
	if !exists {
		return "", fmt.Errorf("no mapping found for website: %s", website)
	}
	// now we will check if shortened value present in shortened Mapper if it's present
	if originalValue, exists := websiteMapper.shortenedToOriginal[parts[1]]; exists {
		return originalValue, nil
	}
	return "", fmt.Errorf("shortened value not found: %s", shortenedValue)
}

type Metrics struct {
	website string
	count   int
}

func (urlShtnr *UrlShortner) Metrics() (map[string]int, error) {
	// define array of metrics
	metrics := make([]Metrics, len(urlShtnr.webSiteMapperMap))
	// now adding values in the metrics
	i:=0
	for website, mapper := range urlShtnr.webSiteMapperMap {
		metrics[i] = Metrics{
			website: website,
			count:   len(mapper.originalToShortened),
		}
		i++
	}
	// now sorting teh metrics by count in descending order
	sort.Slice(metrics, func(i, j int) bool {
		return metrics[i].count > metrics[j].count
	})
	result := make(map[string]int)
	// now we will create a map of metrics
	var output []Metrics
	if len(metrics) > 3 {
		output = metrics[:3]
	} else {
		output = metrics
	}
	for _, metric := range output {
		result[metric.website] = metric.count
	}
	return result, nil
}

type Hasher interface {
	Hash(value string) (string, error)
}

type URLHasher struct {
}

func (hasher URLHasher) Hash(value string) (string, error) {
	// here i will use normal hashing mechnism where i will take sum of all between
	// characters and then convert it to string
	var hash uint64 = 0
	for i := 0; i < len(value); i++ {
		hash = hash*31 + uint64(value[i])
	}

	return fmt.Sprintf("%d", hash), nil
}

// NewUrlShortner creates a new UrlShortner instacnce with hasher of it's own and webSiteMapper
func NewUrlShortner() ShortnerService {
	if shortnerService == nil {
		shortnerService = &UrlShortner{
			webSiteMapperMap: make(map[string]*Mapper),
			hasher:           URLHasher{},
		}
	}
	return shortnerService
}
