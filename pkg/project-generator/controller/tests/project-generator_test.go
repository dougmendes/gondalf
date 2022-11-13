package controller

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	. "github.com/dougmendes/gondalf/pkg/project-generator/controller"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Project Generator", func() {
	When("the user request a new project", func() {
		r := gin.Default()
		w := httptest.NewRecorder()
		BeforeEach(func() {
			r.POST("/project-generator", NewProject)
			req, _ := http.NewRequest("POST", "/project-generator", nil)
			r.ServeHTTP(w, req)
		})
		It("respond with pong", func() {

			mockResponse := `{"message":"new project"}`
			responseData, _ := ioutil.ReadAll(w.Body)
			Expect(string(responseData)).To(Equal(mockResponse))
			Expect(w.Code).To(Equal(200))
		})
	})
})
