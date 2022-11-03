package tests

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/dougmendes/gondalf/pkg/authentication/controller"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}
func TestRegister(t *testing.T) {
	mockResponse := `{"message":"not implement yet"}`
	r := SetUpRouter()
	r.POST("/register", Register)
	req, _ := http.NewRequest("POST", "/register", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)

}
