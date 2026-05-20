package clouddriver_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestClouddriver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Clouddriver Suite")
}
