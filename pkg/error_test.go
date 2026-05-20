package clouddriver_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/gin-gonic/gin"
	clouddriver "github.com/homedepot/go-clouddriver/pkg"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Error", func() {
	Describe("#NewError", func() {
		It("returns an ErrorResponse with correct fields", func() {
			before := time.Now().UnixNano() / 1000000
			e := clouddriver.NewError("Not Found", "resource not found", http.StatusNotFound)
			after := time.Now().UnixNano() / 1000000

			Expect(e.Error).To(Equal("Not Found"))
			Expect(e.Message).To(Equal("resource not found"))
			Expect(e.Status).To(Equal(http.StatusNotFound))
			Expect(e.Timestamp).To(BeNumerically(">=", before))
			Expect(e.Timestamp).To(BeNumerically("<=", after))
		})
	})

	Describe("#Error", func() {
		var (
			c *gin.Context
			w *httptest.ResponseRecorder
		)

		BeforeEach(func() {
			gin.SetMode(gin.ReleaseMode)
			w = httptest.NewRecorder()
			c, _ = gin.CreateTestContext(w)
		})

		It("sets the gin context status and attaches a public error with ErrorMeta", func() {
			clouddriver.Error(c, http.StatusBadRequest, errors.New("bad request"))

			Expect(c.Writer.Status()).To(Equal(http.StatusBadRequest))
			Expect(c.Errors).To(HaveLen(1))

			ginErr := c.Errors[0]
			Expect(ginErr.Error()).To(Equal("bad request"))
			Expect(ginErr.Type).To(Equal(gin.ErrorTypePublic))

			meta := ginErr.Meta.(clouddriver.ErrorMeta)
			Expect(meta.FuncName).ToNot(BeEmpty())
			Expect(meta.FileName).ToNot(BeEmpty())
			Expect(meta.GUID).ToNot(BeEmpty())
			Expect(meta.LineNum).To(BeNumerically(">", 0))
		})
	})

	Describe("#Meta", func() {
		It("extracts ErrorMeta from a gin.Error", func() {
			expectedMeta := clouddriver.ErrorMeta{
				FuncName: "some.Func",
				FileName: "some_file.go",
				GUID:     "abc-123",
				LineNum:  42,
			}

			ginErr := &gin.Error{
				Err:  errors.New("test"),
				Type: gin.ErrorTypePublic,
				Meta: expectedMeta,
			}

			meta := clouddriver.Meta(ginErr)
			Expect(meta.FuncName).To(Equal("some.Func"))
			Expect(meta.FileName).To(Equal("some_file.go"))
			Expect(meta.GUID).To(Equal("abc-123"))
			Expect(meta.LineNum).To(Equal(42))
		})
	})
})
