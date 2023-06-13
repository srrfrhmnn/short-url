package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionRetrieval(t *testing.T) {
	initialURL := "https://srrfrhmn.com/"
	userUUID := "1234567890"
	shortURL := "123456"

	//Persist data mapping
	SaveUrlMapping(shortURL, initialURL, userUUID)

	//Retrieve initial URL
	retrievedUrl := RetrieveInitialUrl(shortURL)

	// Check if the retrieved URL is the same as the initial URL
	assert.Equal(t, initialURL, retrievedUrl)
}
