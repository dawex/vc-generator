package integration_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Healthcheck IT", func() {
	Context("GET /healthcheck", func() {
		It("should get a healtcheck ok response", func() {
			// Prepare the HTTP GET request
			req := httptest.NewRequest(http.MethodGet, "/v1/healthcheck", nil)
			rec := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(rec, req)

			// Verify the HTTP response code
			Expect(rec.Code).To(Equal(http.StatusOK))
		})
	})
})
