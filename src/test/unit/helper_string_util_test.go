package unit

import "testing"
import "../../helpers"
import "github.com/stretchr/testify/assert"


func TestUrlValidator(t *testing.T) {
	valid := helpers.UrlValidator("https://www.google.com")
	invalid := helpers.UrlValidator("google")
	
	assert.Equal(t, valid, true, "It should be true")
	assert.Equal(t, invalid, false, "It should be false")
}

func TestBuildRandomString(t *testing.T) {
	randomString := helpers.BuildRandomString(6)
	randomStringAgain := helpers.BuildRandomString(9)

	assert.Equal(t, len(randomString), 6, "It should be 6 char long")
	assert.Equal(t, len(randomStringAgain), 9, "It should be 9 char long")
}

func TestValidatorShortCode(t *testing.T) {
	valid1 := helpers.ValidatorShortCode("asd123")
	valid2 := helpers.ValidatorShortCode("123asd")

	invalid1 := helpers.ValidatorShortCode("asd12?")
	invalid2 := helpers.ValidatorShortCode("asd12")
	invalid3 := helpers.ValidatorShortCode("")
	
	assert.Equal(t, valid1, true, "It should be true")
	assert.Equal(t, valid2, true, "It should be true")

	assert.Equal(t, invalid1, false, "It should be false")
	assert.Equal(t, invalid2, false, "It should be false")
	assert.Equal(t, invalid3, false, "It should be false")
}