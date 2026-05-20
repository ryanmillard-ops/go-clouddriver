package clouddriver_test

import (
	"bytes"
	"errors"
	"log"

	clouddriver "github.com/homedepot/go-clouddriver/pkg"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Logger", func() {
	Describe("#Log", func() {
		var buf bytes.Buffer

		BeforeEach(func() {
			buf.Reset()
			log.SetOutput(&buf)
		})

		It("logs an error with auto-generated meta", func() {
			clouddriver.Log(errors.New("something went wrong"))

			output := buf.String()
			Expect(output).To(ContainSubstring("[CLOUDDRIVER]"))
			Expect(output).To(ContainSubstring("something went wrong"))
			Expect(output).To(ContainSubstring("LOG"))
		})

		It("logs an error with provided meta", func() {
			meta := clouddriver.ErrorMeta{
				FuncName: "test.Function",
				FileName: "test_file.go",
				GUID:     "guid-123",
				LineNum:  99,
			}

			clouddriver.Log(errors.New("custom error"), meta)

			output := buf.String()
			Expect(output).To(ContainSubstring("[CLOUDDRIVER]"))
			Expect(output).To(ContainSubstring("custom error"))
			Expect(output).To(ContainSubstring("test.Function"))
			Expect(output).To(ContainSubstring("test_file.go"))
			Expect(output).To(ContainSubstring("guid-123"))
			Expect(output).To(ContainSubstring("99"))
		})
	})
})
