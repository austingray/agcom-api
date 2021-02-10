package user

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/austingray/agcom-api/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	d := database.Default()
	r.Use(func(c *gin.Context) {
		c.Set("d", d)
		c.Next()
	})

	r.POST("/api/v1/user/register", Register)
	return r
}

func decodeJSONString(s string) map[string]interface{} {
	byt := []byte(s)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	return dat
}

func TestSetupRouter(t *testing.T) {
	router := setupRouter()
	assert.NotEqual(t, router, nil)
}

func doMockHTTPRequest(data url.Values) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	router := setupRouter()
	req, _ := http.NewRequest("POST", "/api/v1/user/register", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w, req)
	return w
}

func TestRegister(t *testing.T) {
	t.Run("bad request if invalid email", func(t *testing.T) {
		data := url.Values{}
		data.Set("email", "test@invalid-email")
		w := doMockHTTPRequest(data)
		assert.Equal(t, 400, w.Code)
		body := decodeJSONString(w.Body.String())
		assert.Equal(t, body["error"], "Key: 'User.Email' Error:Field validation for 'Email' failed on the 'email' tag")
	})

	t.Run("ok with valid email and password", func(t *testing.T) {
		data := url.Values{}
		data.Set("email", "test@valid-email.com")
		data.Add("password", "test-pass-1234")
		w := doMockHTTPRequest(data)
		assert.Equal(t, 200, w.Code)
	})
}
