package unit

import "testing"
import "time"

import "../../models"
import "../../objects"
import "../../services"

// import "../../repositories"
import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)
import "fmt"

type dbMock struct {
	mock.Mock
}

func (serviceTest *dbMock) GetByShortcode(shortCode string) models.UrlShortCode {
	fmt.Println("Mocked GetByShortcode function")
	// fmt.Printf("Value passed in: %t\n", shortCode)
	resultMock := models.UrlShortCode{}
	resultMock.Uri = "https://www.google.com"
	resultMock.Shortcode = "asd123"
	resultMock.Counter = 5
	resultMock.PublishAt = time.Now()
	resultMock.LastUsedAt = time.Now()
	// return resultMock
	// args := serviceTest.Called(shortCode)
	return resultMock
}
func (serviceTest *dbMock) UpdateShortcodeData(urlShortCode models.UrlShortCode) {
	// doing nothing
	serviceTest.Called(urlShortCode)
}
func (serviceTest *dbMock) CreateShortcode(requestData objects.URLRequestShortRequest) {
	serviceTest.Called(requestData)
	// doing nothing
}

func TestGetUrlShortUrl(t *testing.T) {
	dbMockData := dbMock{}

	resultMock := models.UrlShortCode{}
	resultMock.Uri = "https://www.google.com"
	resultMock.Shortcode = "asd123"
	resultMock.Counter = 5
	resultMock.PublishAt = time.Now()
	resultMock.LastUsedAt = time.Now()
	dbMockData.On("GetByShortcode", "asd123").Return(resultMock)
	dbMockData.On("UpdateShortcodeData", "asd123")

	myService := services.UrlShortService{&dbMockData}
	coba := myService.GetUrlShortUrl("asd123")
	fmt.Println("coba")
	fmt.Println(coba)
	assert.Equal(t, coba.Shortcode, "asd123", "It should be same")
}
