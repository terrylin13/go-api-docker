package index

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHandle(t *testing.T) {
	r := gin.Default()
	r.GET("/v1/", Handle)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/v1/", nil)
	r.ServeHTTP(w, req)
	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error("respon body err:", err)
	}
	t.Log("Status code 是否为200:", resp.StatusCode == http.StatusOK)
	t.Log("resp body:", string(body))
}
