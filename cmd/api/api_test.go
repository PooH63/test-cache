package main

import (
	"bytes"
	"encoding/json"
	"github.com/RGRU/go-memorycache"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/test-cache/internal/api"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

var (
	router *gin.Engine
)

func TestMain(m *testing.M) {
	cStorage := memorycache.New(5*time.Minute, 5*time.Second)
	gin.SetMode("debug")

	router = gin.New()
	setupRoutes(router, cStorage)

	os.Exit(m.Run())
}

func TestOK(t *testing.T) {
	resp := sendRequest(t, "/?a=111", "GET", "", nil)

	assert.Equal(t, resp.Success, true)
	assert.Equal(t, resp.Result, api.STORED.String())

	resp = sendRequest(t, "/?a=111", "GET", "", nil)

	assert.Equal(t, resp.Success, true)
	assert.Equal(t, resp.Result, api.IDENTICAL.String())
	assert.Equal(t, resp.Stored, "111")
	assert.Equal(t, resp.Value, "111")

	resp = sendRequest(t, "/?a=222", "GET", "", nil)

	assert.Equal(t, resp.Success, true)
	assert.Equal(t, resp.Result, api.DIFFER.String())
	assert.Equal(t, resp.Stored, "222")
	assert.Equal(t, resp.Value, "111")
}

func TestError(t *testing.T) {
	resp := sendRequest(t, "/", "GET", "", nil)

	assert.Equal(t, resp.Success, false)
	assert.Equal(t, resp.Error, api.EMPTY_QUERY_PARAMS_ERROR)

	resp = sendRequest(t, "/?a=", "GET", "", nil)

	assert.Equal(t, resp.Success, false)
	assert.Equal(t, resp.Error, api.EMPTY_PARAM_VALUE_ERROR)
}

func performRequest(method, url string, body []byte, headers map[string]string) *httptest.ResponseRecorder {
	var reader io.Reader
	if body != nil {
		reader = bytes.NewReader(body)
	}
	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		panic(err)
	}
	req.Header.Set("x-client-ip", "127.0.0.1")

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func sendRequest(t *testing.T, url, method string, body interface{}, headers map[string]string) *api.Response {
	reqBody, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("web url:", url)
	t.Log("web body:", string(reqBody))
	httpResp := performRequest(method, url, reqBody, headers)

	t.Log("response body:", httpResp.Body.String())

	resp := &api.Response{}
	err = json.Unmarshal(httpResp.Body.Bytes(), resp)
	if err != nil {
		t.Fatal(err)
	}

	return resp
}
