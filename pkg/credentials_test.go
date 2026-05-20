package clouddriver_test

import (
	clouddriver "github.com/homedepot/go-clouddriver/pkg"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Credentials", func() {
	Describe("Credentials struct", func() {
		It("stores and returns all fields correctly", func() {
			creds := clouddriver.Credentials{
				AccountType:                 "kubernetes",
				CacheThreads:                4,
				ChallengeDestructiveActions: true,
				CloudProvider:               "kubernetes",
				Enabled:                     true,
				Environment:                 "dev",
				Name:                        "test-account",
				Namespaces:                  []string{"default", "kube-system"},
				PrimaryAccount:              false,
				ProviderVersion:             "V2",
				Skin:                        "v2",
				Type:                        "kubernetes",
			}

			Expect(creds.AccountType).To(Equal("kubernetes"))
			Expect(creds.CacheThreads).To(Equal(4))
			Expect(creds.ChallengeDestructiveActions).To(BeTrue())
			Expect(creds.CloudProvider).To(Equal("kubernetes"))
			Expect(creds.Enabled).To(BeTrue())
			Expect(creds.Environment).To(Equal("dev"))
			Expect(creds.Name).To(Equal("test-account"))
			Expect(creds.Namespaces).To(ConsistOf("default", "kube-system"))
			Expect(creds.PrimaryAccount).To(BeFalse())
			Expect(creds.ProviderVersion).To(Equal("V2"))
			Expect(creds.Skin).To(Equal("v2"))
			Expect(creds.Type).To(Equal("kubernetes"))
		})

		It("has zero-value defaults", func() {
			creds := clouddriver.Credentials{}

			Expect(creds.Name).To(BeEmpty())
			Expect(creds.CacheThreads).To(BeZero())
			Expect(creds.Enabled).To(BeFalse())
			Expect(creds.Namespaces).To(BeNil())
			Expect(creds.DockerRegistries).To(BeNil())
			Expect(creds.RequiredGroupMembership).To(BeNil())
			Expect(creds.SpinnakerKindMap).To(BeNil())
		})
	})
})
