package alog

import (
	"math/rand"
	"strings"
	"time"
)

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandomString(n ...int) string {
	var intLen int
	if len(n) > 0 {
		intLen = n[0]
	} else {
		intLen = 32
	}
	b := make([]byte, intLen)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := intLen-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return strings.ToLower(string(b))
}
