package oauth

import (
	"crypto"
	_ "crypto/sha1"
	"fmt"
	"io"
	"sort"
	"strings"
)

var HASH_METHOD_MAP = map[crypto.Hash]string{
	crypto.SHA1:   "SHA1",
	crypto.SHA256: "SHA256",
}

// Creates a new Consumer instance, with a HMAC-SHA1 signer
func NewConsumer(consumerKey string, consumerSecret string, requestMethod string, requestURL string) *Consumer {
	consumer := &Consumer{
		consumerKey:    consumerKey,
		consumerSecret: consumerSecret,
		requestMethod:  requestMethod,
		requestURL:     requestURL,

		AdditionalParams: make(map[string]string),
	}

	return consumer
}

type Consumer struct {
	AdditionalParams map[string]string

	// The rest of this class is configured via the NewConsumer function.
	consumerKey string

	consumerSecret string

	requestMethod string

	requestURL string

	debug bool
}

func escape(s string) string {
	t := make([]byte, 0, 3*len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if isEscapable(c) {
			t = append(t, '%')
			t = append(t, "0123456789ABCDEF"[c>>4])
			t = append(t, "0123456789ABCDEF"[c&15])
		} else {
			t = append(t, s[i])
		}
	}
	return string(t)
}

func isEscapable(b byte) bool {
	return !('A' <= b && b <= 'Z' || 'a' <= b && b <= 'z' || '0' <= b && b <= '9' || b == '-' || b == '.' || b == '_' || b == '~')
}

func (c *Consumer) Debug(enabled bool) {
	c.debug = enabled
}

func (c *Consumer) GetRequest(url string) (string, string, io.Reader, error) {
	params := c.baseParams(c.AdditionalParams)

	if c.debug {
		fmt.Println("params:", params)
	}

	result := ""

	for pos, key := range params.Keys() {
		for innerPos, value := range params.Get(key) {
			if pos+innerPos != 0 {
				result += "&"
			}
			result += fmt.Sprintf("%s=%s", key, value)
		}
	}

	if c.debug {
		fmt.Println("req: ", result)
	}

	if c.requestMethod == "GET" {
		return url, result, nil, nil
	} else if c.requestMethod == "POST" {
		return url, "", strings.NewReader(result), nil
	} else {
		return "", "", nil, fmt.Errorf("Not supported method %s", c.requestMethod)
	}
}

func (c *Consumer) baseParams(additionalParams map[string]string) *OrderedParams {
	params := NewOrderedParams()

	for key, value := range additionalParams {
		params.Add(key, value)
	}

	return params
}

func (c *Consumer) requestString(method string, url string, params *OrderedParams) string {
	result := method + "&" + escape(url)
	for pos, key := range params.Keys() {
		for innerPos, value := range params.Get(key) {
			if pos+innerPos == 0 {
				result += "&"
			} else {
				result += escape("&")
			}
			result += escape(fmt.Sprintf("%s=%s", key, value))
		}
	}
	return result
}

//
// String Sorting helpers
//

type ByValue []string

func (a ByValue) Len() int {
	return len(a)
}

func (a ByValue) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByValue) Less(i, j int) bool {
	return a[i] < a[j]
}

//
// ORDERED PARAMS
//

type OrderedParams struct {
	allParams   map[string][]string
	keyOrdering []string
}

func NewOrderedParams() *OrderedParams {
	return &OrderedParams{
		allParams:   make(map[string][]string),
		keyOrdering: make([]string, 0),
	}
}

func (o *OrderedParams) Get(key string) []string {
	sort.Sort(ByValue(o.allParams[key]))
	return o.allParams[key]
}

func (o *OrderedParams) Keys() []string {
	sort.Sort(o)
	return o.keyOrdering
}

func (o *OrderedParams) Add(key, value string) {
	o.AddUnescaped(key, escape(value))
}

func (o *OrderedParams) AddUnescaped(key, value string) {
	if _, exists := o.allParams[key]; !exists {
		o.keyOrdering = append(o.keyOrdering, key)
		o.allParams[key] = make([]string, 1)
		o.allParams[key][0] = value
	} else {
		o.allParams[key] = append(o.allParams[key], value)
	}
}

func (o *OrderedParams) Len() int {
	return len(o.keyOrdering)
}

func (o *OrderedParams) Less(i int, j int) bool {
	return o.keyOrdering[i] < o.keyOrdering[j]
}

func (o *OrderedParams) Swap(i int, j int) {
	o.keyOrdering[i], o.keyOrdering[j] = o.keyOrdering[j], o.keyOrdering[i]
}
