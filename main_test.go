package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func performRequest(c *IndexClient, method, path string) *httptest.ResponseRecorder {

	req := httptest.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	c.ServeHTTP(w, req)
	return w
}

func TestHttpServerCtx_BuildApplication(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	c := NewIndexClient()

	w := performRequest(c, http.MethodGet, "/err/404")
	assert.Equal(t, http.StatusNotFound, w.Code)
	w = performRequest(c, http.MethodGet, "/err/403")
	assert.Equal(t, http.StatusForbidden, w.Code)
	w = performRequest(c, http.MethodGet, "/err/500")
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	w = performRequest(c, http.MethodGet, "/err/502")
	assert.Equal(t, http.StatusBadGateway, w.Code)
	w = performRequest(c, http.MethodGet, "/ip")
	assert.Equal(t, http.StatusOK, w.Code)
	fmt.Println("current_ip", w.Body.String())
	w = performRequest(c, http.MethodGet, "/ip")
	assert.Equal(t, http.StatusOK, w.Code)
	w = performRequest(c, http.MethodGet, "/1s")
	assert.Equal(t, http.StatusOK, w.Code)
	//w = performRequest(c, http.MethodGet, "/panic")
	//assert.Equal(t, http.StatusInternalServerError, w.Code)
}
