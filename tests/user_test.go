package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"jd_workout_golang/internal/router"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUser(t *testing.T) {
	// 註冊 router
	r := gin.Default()
	router.RegisterUser(r.Group("/api"))

	// 建立 response & request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/user", nil)

	r.ServeHTTP(w, req)

	// assert response
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "user", w.Body.String())

	if w.Code != 200 {
		t.Errorf("Response code is %v", w.Code)
	}
}
