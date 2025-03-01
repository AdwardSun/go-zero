package internal

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogInterceptor(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	req, err := http.NewRequest(http.MethodGet, svr.URL, nil)
	assert.Nil(t, err)
	req, handler := LogInterceptor(req)
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	handler(resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestLogInterceptorServerError(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	req, err := http.NewRequest(http.MethodGet, svr.URL, nil)
	assert.Nil(t, err)
	req, handler := LogInterceptor(req)
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	handler(resp)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}
