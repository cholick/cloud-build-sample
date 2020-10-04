package controller_test

import (
	"io/ioutil"
	"net/http/httptest"

	"github.com/cholick/cloud-build-sample/pkg/controller"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Index", func() {
	It("writes success", func() {
		req := httptest.NewRequest("GET", "http://example.com/foo", nil)
		resp := httptest.NewRecorder()

		controller.Index(resp, req, nil)

		Expect(resp.Code).To(Equal(200))

		body, _ := ioutil.ReadAll(resp.Body)
		Expect(body).To(ContainSubstring("Success"))
		Expect(body).To(ContainSubstring("/foo"))
	})
})
