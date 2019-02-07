package helpers

import (
	// "fmt"
	"regexp"
	"math/rand"
	"net/url"
)

func UrlValidator(urldata string) bool {
    _, err := url.ParseRequestURI(urldata)
    if err != nil {
        return false
    }
    return true
}

// regex matcher = ^[0-9a-zA-Z_]{6}$
// TODO : move letterBytes to constant
const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
func BuildRandomString(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    // fmt.Println(b)
    return string(b)
}

func ValidatorShortCode(data string) bool {
	validString := regexp.MustCompile(`^[0-9a-zA-Z_]{6}$`)
	return validString.MatchString(data)
}