package bing_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/wonsikin/dictionary/src"
)

func TestQueryDictionary(t *testing.T) {
	r := src.Router()
	server := httptest.NewServer(r)
	defer server.Close()

	url := server.URL + "/wd/bing/falcon"

	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("send request for %s fail: %s", url, err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
