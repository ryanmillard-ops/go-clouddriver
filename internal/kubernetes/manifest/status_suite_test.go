package manifest_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestManifestStatus(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Manifest Status Suite")
}
