package clouddriver_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	clouddriver "github.com/homedepot/go-clouddriver/pkg"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Task", func() {
	Describe("#NewDefaultTask", func() {
		It("returns a task with correct default status fields", func() {
			task := clouddriver.NewDefaultTask("test-id")

			Expect(task.ID).To(Equal("test-id"))
			Expect(task.ResultObjects).To(BeEmpty())
			Expect(task.Status.Complete).To(BeTrue())
			Expect(task.Status.Completed).To(BeTrue())
			Expect(task.Status.Failed).To(BeFalse())
			Expect(task.Status.Phase).To(Equal("ORCHESTRATION"))
			Expect(task.Status.Retryable).To(BeFalse())
			Expect(task.Status.Status).To(Equal("Orchestration completed."))
		})
	})

	Describe("#TaskIDFromContext", func() {
		It("retrieves the task ID from a gin context", func() {
			gin.SetMode(gin.ReleaseMode)
			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			r, err := http.NewRequest(http.MethodGet, "", nil)
			Expect(err).To(BeNil())
			c.Request = r
			c.Set(clouddriver.TaskIDKey, "my-task-id")

			taskID := clouddriver.TaskIDFromContext(c)
			Expect(taskID).To(Equal("my-task-id"))
		})
	})
})
