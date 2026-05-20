package middleware_test

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/homedepot/go-clouddriver/internal/middleware"
	clouddriver "github.com/homedepot/go-clouddriver/pkg"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Error", func() {
	Describe("#HandleError", func() {
		var (
			w *httptest.ResponseRecorder
			c *gin.Context
		)

		BeforeEach(func() {
			gin.SetMode(gin.ReleaseMode)
			w = httptest.NewRecorder()
			c, _ = gin.CreateTestContext(w)
			r, err := http.NewRequest(http.MethodGet, "", nil)
			Expect(err).To(BeNil())
			c.Request = r
		})

		When("no error occurs", func() {
			It("does not write a JSON body", func() {
				mw := middleware.HandleError()
				mw(c)

				Expect(w.Body.Len()).To(BeZero())
			})
		})

		When("a public error is attached with status 400", func() {
			BeforeEach(func() {
				clouddriver.Error(c, http.StatusBadRequest, errors.New("invalid input"))
			})

			It("returns a JSON error response with the correct structure", func() {
				mw := middleware.HandleError()
				mw(c)

				Expect(w.Code).To(Equal(http.StatusBadRequest))

				body, readErr := io.ReadAll(w.Body)
				Expect(readErr).To(BeNil())

				var errResp clouddriver.ErrorResponse
				Expect(json.Unmarshal(body, &errResp)).To(Succeed())
				Expect(errResp.Error).To(Equal("Bad Request"))
				Expect(errResp.Message).To(Equal("invalid input"))
				Expect(errResp.Status).To(Equal(http.StatusBadRequest))
				Expect(errResp.Timestamp).To(BeNumerically(">", 0))
			})
		})

		When("a public error is attached with status 500", func() {
			BeforeEach(func() {
				clouddriver.Error(c, http.StatusInternalServerError, errors.New("server failure"))
			})

			It("includes the GUID in the response text", func() {
				mw := middleware.HandleError()
				mw(c)

				Expect(w.Code).To(Equal(http.StatusInternalServerError))

				body, readErr := io.ReadAll(w.Body)
				Expect(readErr).To(BeNil())

				var errResp clouddriver.ErrorResponse
				Expect(json.Unmarshal(body, &errResp)).To(Succeed())
				Expect(errResp.Error).To(ContainSubstring("Internal Server Error"))
				Expect(errResp.Error).To(ContainSubstring("error ID:"))
				Expect(errResp.Message).To(Equal("server failure"))
				Expect(errResp.Status).To(Equal(http.StatusInternalServerError))
			})
		})
	})
})
