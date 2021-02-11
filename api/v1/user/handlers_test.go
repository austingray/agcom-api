package user

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/austingray/agcom-api/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	t.Run("bad request if invalid email", func(t *testing.T) {
		data := url.Values{}
		data.Set("email", "test@invalid-email")
		w, _ := doHTTPTest(data)
		body := decodeJSONString(w.Body.String())

		// 400 response code
		assert.Equal(t, 400, w.Code)

		// correct error message
		assert.Equal(t, body["error"], "Key: 'User.Email' Error:Field validation for 'Email' failed on the 'email' tag")
	})

	t.Run("ok with valid email", func(t *testing.T) {
		data := url.Values{}
		data.Set("email", "test@valid-email.com")
		data.Add("password", "test-pass-1234")
		w, _ := doHTTPTest(data)
		assert.Equal(t, 200, w.Code)
	})

	t.Run("creates the user", func(t *testing.T) {
		// generate a random email
		rand.Seed(time.Now().UnixNano())
		email := randSeq(10) + "@unique-email.com"

		// do request
		data := url.Values{}
		data.Set("email", email)
		data.Add("password", "test-pass-1234")
		w, d := doHTTPTest(data)

		// find and assert equal
		user := d.GetUserByEmail(email)
		body := decodeJSONString(w.Body.String())
		respUser := body["user"].(map[string]interface{})
		assert.Equal(t, user.Email, respUser["email"])
	})
}

func doHTTPTest(data url.Values) (*httptest.ResponseRecorder, *database.Database) {
	w := httptest.NewRecorder()
	r := gin.Default()
	d := database.Test()

	r.Use(func(c *gin.Context) {
		c.Set("d", d)
		c.Next()
	})

	// register routes
	r.POST("/api/v1/user/register", Register)

	req, _ := http.NewRequest("POST", "/api/v1/user/register", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	r.ServeHTTP(w, req)
	return w, d
}

func decodeJSONString(s string) map[string]interface{} {
	byt := []byte(s)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	return dat
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
