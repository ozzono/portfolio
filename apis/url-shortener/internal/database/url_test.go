package database

import (
	"fmt"
	"testing"

	"url-shortener/internal/models"
	"url-shortener/utils"

	"github.com/stretchr/testify/require"
)

func TestURL(t *testing.T) {

	testURL := &models.URL{
		Source: fmt.Sprintf("https://%s.%s", utils.RString(5, 7), utils.RString(2, 3)),
	}

	client, err := NewClient()
	require.NoError(t, err, "NewClient")

	testURL, err = client.AddURL(testURL, false)
	require.NoError(t, err, "failed to store test url")

	foundURL, found, err := client.FindURLBySource(testURL, false)
	require.NoError(t, err, "client.FindURLBySource")
	require.True(t, found, "source url not found")

	shortURL, found, err := client.FindURLByShortened(testURL, false)
	require.NoError(t, err, "client.FindURLByShortened")
	require.True(t, found, "url not found by shortened reference")

	require.Equal(t, shortURL, foundURL, "shortURL != foundURL")

	incURL, err := client.IncrementURL(foundURL, false)
	require.NoError(t, err, "client.IncrementURL")
	require.Equal(t, foundURL.Count+1, incURL.Count, "failed to increment url counter")

	err = client.DelURL(foundURL)
	require.NoError(t, err, "client.DelURL")

	_, found, err = client.FindURLBySource(foundURL, false)
	require.NoError(t, err, "client.FindURLBySource")
	require.True(t, !found, "found URL that should not exist")

	incURL.Log("found and incremented url", true)
}
