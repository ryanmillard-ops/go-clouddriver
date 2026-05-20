package middleware_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/homedepot/go-clouddriver/internal/middleware"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Log", func() {
	Describe("#LogRequest", func() {
		var (
			c *gin.Context
		)

		BeforeEach(func() {
			gin.SetMode(gin.ReleaseMode)
			c, _ = gin.CreateTestContext(httptest.NewRecorder())
		})

		It("reads the request body and still allows downstream handlers to read it", func() {
			requestBody := `{"key":"value"}`
			r, err := http.NewRequest(http.MethodPost, "http://localhost/test", strings.NewReader(requestBody))
			Expect(err).To(BeNil())
			c.Request = r

			mw := middleware.LogRequest()
			mw(c)

			// The body should still be readable after the middleware runs.
			body, readErr := io.ReadAll(c.Request.Body)
			Expect(readErr).To(BeNil())
			Expect(string(body)).To(Equal(requestBody))
		})

		It("calls c.Next to pass through to the next handler", func() {
			nextCalled := false

			engine := gin.New()
			engine.Use(middleware.LogRequest())
			engine.Use(func(ctx *gin.Context) {
				nextCalled = true
			})
			engine.GET("/test", func(ctx *gin.Context) {})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/test", strings.NewReader(""))
			engine.ServeHTTP(w, req)

			Expect(nextCalled).To(BeTrue())
		})
	})
})
