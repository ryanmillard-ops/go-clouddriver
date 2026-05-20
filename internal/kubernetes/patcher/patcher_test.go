package patcher

import (
	"encoding/json"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/cli-runtime/pkg/resource"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Patcher", func() {
	Describe("#New", func() {
		It("returns a Patcher with correct default field values", func() {
			info := &resource.Info{
				Mapping: &meta.RESTMapping{
					Resource: schema.GroupVersionResource{
						Group:    "",
						Version:  "v1",
						Resource: "pods",
					},
					GroupVersionKind: schema.GroupVersionKind{
						Group:   "",
						Version: "v1",
						Kind:    "Pod",
					},
				},
			}
			helper := &resource.Helper{}

			p, err := New(info, helper)
			Expect(err).To(BeNil())
			Expect(p).ToNot(BeNil())
			Expect(p.Overwrite).To(BeTrue())
			Expect(p.Force).To(BeFalse())
			Expect(p.Cascade).To(BeTrue())
			Expect(p.Timeout).To(Equal(time.Duration(0)))
			Expect(p.GracePeriod).To(Equal(-1))
			Expect(p.Retries).To(Equal(0))
			Expect(p.ResourceVersion).To(BeNil())
			Expect(p.Mapping).To(Equal(info.Mapping))
			Expect(p.Helper).To(Equal(helper))
		})
	})

	Describe("#addResourceVersion", func() {
		It("injects a resourceVersion into JSON patch bytes", func() {
			original := map[string]interface{}{
				"apiVersion": "v1",
				"kind":       "ConfigMap",
				"metadata": map[string]interface{}{
					"name": "test",
				},
			}
			patchBytes, err := json.Marshal(original)
			Expect(err).To(BeNil())

			result, err := addResourceVersion(patchBytes, "12345")
			Expect(err).To(BeNil())

			var resultMap map[string]interface{}
			err = json.Unmarshal(result, &resultMap)
			Expect(err).To(BeNil())

			metadata, ok := resultMap["metadata"].(map[string]interface{})
			Expect(ok).To(BeTrue())
			Expect(metadata["resourceVersion"]).To(Equal("12345"))
			Expect(metadata["name"]).To(Equal("test"))
		})

		It("returns an error for invalid JSON", func() {
			_, err := addResourceVersion([]byte("not-json"), "12345")
			Expect(err).ToNot(BeNil())
		})
	})

	Describe("#asDeleteOptions", func() {
		When("cascade is true", func() {
			It("returns DeletePropagationForeground", func() {
				opts := asDeleteOptions(true, -1)
				Expect(*opts.PropagationPolicy).To(Equal(metav1.DeletePropagationForeground))
			})
		})

		When("cascade is false", func() {
			It("returns DeletePropagationOrphan", func() {
				opts := asDeleteOptions(false, -1)
				Expect(*opts.PropagationPolicy).To(Equal(metav1.DeletePropagationOrphan))
			})
		})

		When("gracePeriod >= 0", func() {
			It("sets GracePeriodSeconds on the delete options", func() {
				opts := asDeleteOptions(true, 30)
				Expect(*opts.GracePeriodSeconds).To(Equal(int64(30)))
				Expect(*opts.PropagationPolicy).To(Equal(metav1.DeletePropagationForeground))
			})
		})

		When("gracePeriod is 0", func() {
			It("sets GracePeriodSeconds to 0", func() {
				opts := asDeleteOptions(true, 0)
				Expect(*opts.GracePeriodSeconds).To(Equal(int64(0)))
			})
		})

		When("gracePeriod < 0", func() {
			It("does not set GracePeriodSeconds", func() {
				opts := asDeleteOptions(true, -1)
				Expect(opts.GracePeriodSeconds).To(BeNil())
			})
		})

		When("cascade is false and gracePeriod >= 0", func() {
			It("returns DeletePropagationOrphan with GracePeriodSeconds", func() {
				opts := asDeleteOptions(false, 10)
				Expect(*opts.PropagationPolicy).To(Equal(metav1.DeletePropagationOrphan))
				Expect(*opts.GracePeriodSeconds).To(Equal(int64(10)))
			})
		})
	})
})
