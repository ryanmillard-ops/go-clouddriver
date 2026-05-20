package core_test

import (
	"io"
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SecurityGroups", func() {
	Describe("#ListSecurityGroups", func() {
		BeforeEach(func() {
			setup()
			uri = svr.URL + "/securityGroups"
			createRequest(http.MethodGet)
		})

		AfterEach(func() {
			teardown()
		})

		JustBeforeEach(func() {
			doRequest()
		})

		It("returns an empty JSON object", func() {
			Expect(res.StatusCode).To(Equal(http.StatusOK))

			b, readErr := io.ReadAll(res.Body)
			Expect(readErr).To(BeNil())
			Expect(string(b)).To(MatchJSON("{}"))
		})
	})
})
