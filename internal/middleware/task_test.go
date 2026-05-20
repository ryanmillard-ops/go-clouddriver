package middleware_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/homedepot/go-clouddriver/internal/middleware"
	clouddriver "github.com/homedepot/go-clouddriver/pkg"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Task", func() {
	Describe("#TaskID", func() {
		var (
			c *gin.Context
		)

		BeforeEach(func() {
			gin.SetMode(gin.ReleaseMode)
			c, _ = gin.CreateTestContext(httptest.NewRecorder())
			r, err := http.NewRequest(http.MethodGet, "", nil)
			Expect(err).To(BeNil())
			c.Request = r
		})

		It("sets a UUID string on the gin context under TaskIDKey", func() {
			mw := middleware.TaskID()
			mw(c)

			val, exists := c.Get(clouddriver.TaskIDKey)
			Expect(exists).To(BeTrue())

			taskID, ok := val.(string)
			Expect(ok).To(BeTrue())
			Expect(taskID).ToNot(BeEmpty())
			Expect(taskID).To(HaveLen(36))
		})
	})
})
