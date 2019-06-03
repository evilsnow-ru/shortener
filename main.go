package shortener

import (
	"fmt"
	"hash/fnv"
	"strconv"
)

const (
	notFoundValue         = ""
	emptyHashValue        = ""
	maxUInt64Value        = 18446744073709551615
	maxUInt64CheckedValue = maxUInt64Value - 1
)

type Shortener interface {
	Shorten(url string) string
	Resolve(url string) string
}

type shortenerImpl struct {
	urlMap map[string]string
}

func makeHash(str string) (uint64, error) {
	hashBuilder := fnv.New64()
	_, err := hashBuilder.Write([]byte(str))

	if err != nil {
		return 0, err
	}

	hashValue := hashBuilder.Sum64()
	return hashValue, nil
}

func hashToChars(hashValue uint64) string {
	return strconv.FormatUint(hashValue, 36)
}

func (shortener *shortenerImpl) Shorten(url string) string {
	if url == "" {
		return emptyHashValue
	}

	hashValue, err := makeHash(url)

	if err != nil {
		fmt.Printf("Error generating short link for url \"%s\": %s.", url, err.Error())
		return emptyHashValue
	}

	shortLink := hashToChars(hashValue)

	savedLink, collision := shortener.urlMap[url]

	//Using open hash addressing if our url is different
	//from saved in map with same hash value
	if collision && savedLink != url {
		var tryCount uint64 = 0

		for collision {
			if tryCount == maxUInt64Value {
				break
			}

			if hashValue == maxUInt64CheckedValue {
				hashValue = 0
			} else {
				hashValue += 1
			}

			shortLink = hashToChars(hashValue)
			savedLink, collision = shortener.urlMap[url]

			if collision && savedLink == url {
				collision = false
			}

			tryCount++
		}

		if tryCount == maxUInt64Value {
			fmt.Println("No empty slots for storing new short link")
			return emptyHashValue
		}
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
