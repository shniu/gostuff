package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/shniu/cache/cache"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test http handler: https://blog.questionable.services/article/testing-http-handlers-go/

var s *Server
var c cache.Cache

func init() {
	c = cache.NewInMemoryCache()
	s = &Server{cache: c}
}

func Test_handleCache_Given_KeyAndValue_Then_PutSucceed(t *testing.T) {

	key := "uid:12345"
	request, e := http.NewRequest(http.MethodPut, "/api/cache/"+key, bytes.NewBuffer([]byte("unit test")))
	assert.Nil(t, e,
		fmt.Sprintf("New request err should be nil, but got %v", e))

	resp := getTestHandlerResp(t, request, s.handleCache)

	assertCode(t, resp, 0)
	assertMessage(t, resp, "succeed")

	if v, e := c.Get(key); e == nil {
		fmt.Println("v: ", string(v))
		assert.Equal(t, "unit test", string(v))
	}

}

func Test_handleCache_When_GetNotExistKey_Then_GotNon(t *testing.T) {
	req, e := http.NewRequest(http.MethodGet, "/api/cache/non-exists-key", nil)
	assert.Nil(t, e, "New request err should be nil")

	resp := getTestHandlerResp(t, req, s.handleCache)

	assertCode(t, resp, 404)
	assertData(t, resp, "")
}

func Test_handleCache_When_GetExistKey_Then_GotTheValue(t *testing.T) {
	key := "uid:12345"
	_, _ = http.NewRequest(http.MethodPut, "/api/cache/"+key, bytes.NewBuffer([]byte("unit test")))
	req, _ := http.NewRequest(http.MethodGet, "/api/cache/"+key, nil)
	resp := getTestHandlerResp(t, req, s.handleCache)
	assertCode(t, resp, 0)
	assertData(t, resp, "unit test")
}

func Test_handleCache_When_DelExistKey_Then_GotSucceed(t *testing.T) {
	key := "k1"
	_, _ = http.NewRequest(http.MethodPut, "/api/cache/"+key, bytes.NewBuffer([]byte("opus!!!")))
	req, _ := http.NewRequest(http.MethodDelete, "/api/cache/"+key, nil)
	resp := getTestHandlerResp(t, req, s.handleCache)
	assertCode(t, resp, 0)
	_, e := s.cache.Get(key)
	assert.NotNil(t, e)
	assert.Equal(t, "the key does not exist", e.Error())
}

func assertMessage(t *testing.T, resp map[string]interface{}, exceptedMsg string) {
	msg, ok := resp["message"]
	assert.True(t, ok)
	assert.Equal(t, exceptedMsg, msg)
}

func assertCode(t *testing.T, resp map[string]interface{}, exceptedCode int) {
	code, ok := resp["code"]
	assert.True(t, ok)
	icode := int(code.(float64))
	assert.Equal(t, exceptedCode, icode)
}

func getTestHandlerResp(t *testing.T, request *http.Request, handlerFunc func(http.ResponseWriter, *http.Request)) map[string]interface{} {
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlerFunc)
	handler.ServeHTTP(rr, request)

	assert.Equal(t, http.StatusOK, rr.Code,
		fmt.Sprintf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK))

	var resp map[string]interface{}
	e := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, e, fmt.Sprintf("Unmarshal error should be nil, but got %v", e))

	return resp
}

func assertData(t *testing.T, resp map[string]interface{}, exceptedData string) {
	data, ok := resp["data"]
	assert.True(t, ok)
	assert.Equal(t, exceptedData, data)
}
