package user

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/api/v1/user/register", Register)
	return r
}

func TestRouter(t *testing.T) {
	router := setupRouter()
	assert.NotEqual(t, router, nil)
}

func TestRegister(t *testing.T) {
	// setup
	router := setupRouter()
	w := httptest.NewRecorder()

	// prepare
	data := url.Values{}
	data.Set("email", "test@email.com")
	data.Add("password", "test-pass-1234")

	// serve
	req, _ := http.NewRequest("POST", "/api/v1/user/register", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w, req)

	// test
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"email\":\"test@email.com\"}", w.Body.String())
}
