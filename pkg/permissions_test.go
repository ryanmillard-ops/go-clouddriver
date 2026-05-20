package clouddriver_test

import (
	clouddriver "github.com/homedepot/go-clouddriver/pkg"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Permissions", func() {
	Describe("ReadPermission", func() {
		It("returns the correct table name", func() {
			rp := clouddriver.ReadPermission{}
			Expect(rp.TableName()).To(Equal("provider_read_permissions"))
		})
	})

	Describe("WritePermission", func() {
		It("returns the correct table name", func() {
			wp := clouddriver.WritePermission{}
			Expect(wp.TableName()).To(Equal("provider_write_permissions"))
		})
	})
})
