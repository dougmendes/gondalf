package controller_test

import (
	"net/http"

	"github.com/dougmendes/gondalf/pkg/instances/controller"
	"github.com/dougmendes/gondalf/pkg/instances/model"
	mock_usecases "github.com/dougmendes/gondalf/tests/mock/instances"
	"github.com/dougmendes/gondalf/tests/stubs"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Instances", func() {
	var pattern string
	var router *stubs.RouterStub
	var w *stubs.ResponseWriterStub
	var url string
	var mockCtrl *gomock.Controller

	Context("Instances Attributes", func() {
		BeforeEach(func() {
			pattern = "/instances/{region}"
		})

		When("Instances are found", func() {
			BeforeEach(func() {
				mockCtrl = gomock.NewController(GinkgoT())
				solver := mock_usecases.NewMockInstanceSolver(mockCtrl)
				solver.EXPECT().GetAllInstances("XBKQwthuUqd").Return(&model.InstanceList{}, nil)
				controller := controller.NewInstanceController(solver)
				router = &stubs.RouterStub{}
				controller.AddRoutes(router)
				w = &stubs.ResponseWriterStub{}
				url = pattern
			})
			It("Should respond status code 200", func() {
				router.Exec(url, w, &http.Request{})
				Expect(200).To(Equal(w.StatusCode))
			})
		})
	})
})
