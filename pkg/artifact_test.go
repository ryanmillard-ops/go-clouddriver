package clouddriver_test

import (
	"github.com/homedepot/go-clouddriver/internal/artifact"
	clouddriver "github.com/homedepot/go-clouddriver/pkg"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Artifact", func() {
	Describe("Artifact struct", func() {
		It("stores and returns all fields correctly", func() {
			a := clouddriver.Artifact{
				CustomKind: true,
				Location:   "us-east-1",
				Metadata: clouddriver.ArtifactMetadata{
					Account: "my-account",
				},
				Name:      "my-artifact",
				Reference: "gs://bucket/path",
				Type:      artifact.TypeGCSObject,
				Version:   "v1",
			}

			Expect(a.CustomKind).To(BeTrue())
			Expect(a.Location).To(Equal("us-east-1"))
			Expect(a.Metadata.Account).To(Equal("my-account"))
			Expect(a.Name).To(Equal("my-artifact"))
			Expect(a.Reference).To(Equal("gs://bucket/path"))
			Expect(a.Type).To(Equal(artifact.TypeGCSObject))
			Expect(a.Version).To(Equal("v1"))
		})

		It("omits empty location and version in JSON", func() {
			a := clouddriver.Artifact{
				Name:      "test",
				Reference: "ref",
			}

			Expect(a.Location).To(BeEmpty())
			Expect(a.Version).To(BeEmpty())
		})
	})
})
