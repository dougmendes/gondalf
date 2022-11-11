package tests

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	. "github.com/dougmendes/gondalf/pkg/instances/controller"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ping", func() {
	When("the user send a ping", func() {
		r := gin.Default()
		w := httptest.NewRecorder()
		BeforeEach(func() {
			r.GET("/ping", Pong)
			req, _ := http.NewRequest("GET", "/ping", nil)
			r.ServeHTTP(w, req)
		})
		It("respond with pong", func() {

			mockResponse := `{"message":"pong!!"}`
			responseData, _ := ioutil.ReadAll(w.Body)
			Expect(string(responseData)).To(Equal(mockResponse))
			Expect(w.Code).To(Equal(200))
		})
	})
})
