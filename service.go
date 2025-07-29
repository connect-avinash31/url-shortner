package main

import (
	"fmt"
)

func main() {
	// first will be writing an url shortner and then running an Server to handle it
	urlShortner := NewUrlShortnerWithDefaultHasher()
	// now we can call ShortenValue and OriginalValue methods on urlShortner
	// exmaple
	shortenedValue, err := urlShortner.ShortenValue("www.udemy.com/courses/ai-course")
	if err != nil {
		println("Error shortening value:", err.Error())
		return
	}
	fmt.Println("Shortened Value:", shortenedValue)
	originalValue, err := urlShortner.OriginalValue(shortenedValue)
	if err != nil {
		println("Error getting original value:", err.Error())
		return
	}
	fmt.Println("Original Value:", originalValue)
}
