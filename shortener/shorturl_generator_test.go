package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const UserId = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortLinkGenerator(t *testing.T) {
	initalLink_1 := "https://www.apple.com/ca/business/enterprise/"
	shortLink_1 := GenerateShortURL(initalLink_1, UserId)

	initialLink_2 := "https://about.google/"
	shortLink_2 := GenerateShortURL(initialLink_2, UserId)

	initialLink_3 := "https://www.microsoft.com/en-ca/about/"
	shortLink_3 := GenerateShortURL(initialLink_3, UserId)

	assert.Equal(t, "J9sgTr2T", shortLink_1)
	assert.Equal(t, "7Ki5z8PN", shortLink_2)
	assert.Equal(t, "PyjxzyzS", shortLink_3)
}
