package shortener

import (
	"errors"
	"fmt"
	"hash/fnv"
	"strconv"
)

const (
	notFoundValue = ""
	emptyHashValue = ""
)

type Shortener interface {
	Shorten(url string) string
	Resolve(url string) string
}

type shortenerImpl struct {
	urlMap map[string]string
}

func makeHash(str string) (string, error) {
	if str == "" {
		return emptyHashValue, errors.New("empty string not supported")
	}

	hashBuilder := fnv.New64()
	_, err := hashBuilder.Write([]byte(str))

	if err != nil {
		return emptyHashValue, err
	}

	hashValue := hashBuilder.Sum64()
	return strconv.FormatUint(hashValue, 36), nil
}

func (shortener *shortenerImpl) Shorten(url string) string {
	shortLink, err := makeHash(url)

	if err != nil {
		fmt.Printf("Error generating short link for url \"%s\": %s.", url, err.Error())
		return emptyHashValue
	}

	shortener.urlMap[shortLink] = url
	return shortLink
}

func (shortener *shortenerImpl) Resolve(url string) string {
	originalUrl, present := shortener.urlMap[url]

	if present {
		return originalUrl
	} else {
		return notFoundValue
	}

}

func New() Shortener {
	return &shortenerImpl{
		urlMap: make(map[string]string),
	}
}