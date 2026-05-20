package core_test

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/homedepot/go-clouddriver/internal/api/core"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Features", func() {
	Describe("#ListStages", func() {
		BeforeEach(func() {
			setup()
			uri = svr.URL + "/features/stages"
			createRequest(http.MethodGet)
		})

		AfterEach(func() {
			teardown()
		})

		JustBeforeEach(func() {
			doRequest()
		})

		It("returns a list of stages with all enabled set to false", func() {
			Expect(res.StatusCode).To(Equal(http.StatusOK))

			b, readErr := io.ReadAll(res.Body)
			Expect(readErr).To(BeNil())

			var stages core.Stages
			Expect(json.Unmarshal(b, &stages)).To(Succeed())
			Expect(stages).ToNot(BeEmpty())

			for _, s := range stages {
				Expect(s.Enabled).To(BeFalse())
				Expect(s.Name).ToNot(BeEmpty())
			}
		})

		It("includes expected stage names", func() {
			Expect(res.StatusCode).To(Equal(http.StatusOK))

			b, readErr := io.ReadAll(res.Body)
			Expect(readErr).To(BeNil())

			var stages core.Stages
			Expect(json.Unmarshal(b, &stages)).To(Succeed())

			names := []string{}
			for _, s := range stages {
				names = append(names, s.Name)
			}

			Expect(names).To(ContainElement("deployManifest"))
			Expect(names).To(ContainElement("deleteManifest"))
			Expect(names).To(ContainElement("patchManifest"))
			Expect(names).To(ContainElement("runJob"))
		})
	})
})
