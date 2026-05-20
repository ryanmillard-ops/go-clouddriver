package manifest_test

import (
	"github.com/homedepot/go-clouddriver/internal/kubernetes/manifest"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Status", func() {
	Describe("DefaultStatus", func() {
		It("has Available.State=true", func() {
			Expect(manifest.DefaultStatus.Available.State).To(BeTrue())
			Expect(manifest.DefaultStatus.Available.Message).To(BeEmpty())
		})

		It("has Failed.State=false", func() {
			Expect(manifest.DefaultStatus.Failed.State).To(BeFalse())
			Expect(manifest.DefaultStatus.Failed.Message).To(BeEmpty())
		})

		It("has Paused.State=false", func() {
			Expect(manifest.DefaultStatus.Paused.State).To(BeFalse())
			Expect(manifest.DefaultStatus.Paused.Message).To(BeEmpty())
		})

		It("has Stable.State=true", func() {
			Expect(manifest.DefaultStatus.Stable.State).To(BeTrue())
			Expect(manifest.DefaultStatus.Stable.Message).To(BeEmpty())
		})
	})

	Describe("NoneReported", func() {
		It("has Available.State=false with correct message", func() {
			Expect(manifest.NoneReported.Available.State).To(BeFalse())
			Expect(manifest.NoneReported.Available.Message).To(Equal("No availability reported"))
		})

		It("has Failed.State=false", func() {
			Expect(manifest.NoneReported.Failed.State).To(BeFalse())
			Expect(manifest.NoneReported.Failed.Message).To(BeEmpty())
		})

		It("has Paused.State=false", func() {
			Expect(manifest.NoneReported.Paused.State).To(BeFalse())
			Expect(manifest.NoneReported.Paused.Message).To(BeEmpty())
		})

		It("has Stable.State=false with correct message", func() {
			Expect(manifest.NoneReported.Stable.State).To(BeFalse())
			Expect(manifest.NoneReported.Stable.Message).To(Equal("No status reported yet"))
		})
	})
})
