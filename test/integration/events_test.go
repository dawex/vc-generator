package integration_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/dawex/vc-generator/internal/common/db/models"
	event_ports "github.com/dawex/vc-generator/internal/services/event/ports"
	vc_ports "github.com/dawex/vc-generator/internal/services/verifiable-credential/ports"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Events API IT", func() {
	var (
		event1CreatedID uuid.UUID
		event2CreatedID uuid.UUID
	)

	Context("Insert Events and Fetch Events List", func() {
		It("OK - should fetch an empty events list. No records yet.", func() {
			// Prepare the HTTP GET request
			req := httptest.NewRequest(http.MethodGet, "/v1/events?contractId=gfjfgd45555sd5s5s5s5s&executionId=uktdqs4661qsdqs664141", nil)
			rec := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(rec, req)

			// Verify the HTTP response code
			Expect(rec.Code).To(Equal(http.StatusOK))

			var body []event_ports.Event
			err := json.NewDecoder(rec.Body).Decode(&body)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(body)).To(Equal(0))
		})

		It("OK - should insert a event start item into the database", func() {
			// Prepare the HTTP POST request
			reqBody := `{
				"contractId": "gfjfgd45555sd5s5s5s5s",
				"executionId": "uktdqs4661qsdqs664141",
				"monitoringEvent": {
					"source": "transfer-plugin",
					"timestamp": "2025-02-26T13:28:09Z",
					"metric": "action-start",
					"value": "calculation of statistics",
					"log": "Root WebApplicationContext: initialization completed in 2758 ms"
				}
			}`
			req := httptest.NewRequest(http.MethodPost, "/v1/events", strings.NewReader(reqBody))
			rec := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(rec, req)

			// Verify the HTTP response code
			Expect(rec.Code).To(Equal(http.StatusOK))

			var body event_ports.Event
			err := json.NewDecoder(rec.Body).Decode(&body)
			Expect(err).ToNot(HaveOccurred())
			Expect(body.Id).ShouldNot(BeNil())
			Expect(body.CreatedAt).ShouldNot(BeNil())
			Expect(body.ContractId).To(Equal("gfjfgd45555sd5s5s5s5s"))
			Expect(body.ExecutionId).To(Equal("uktdqs4661qsdqs664141"))
			Expect(body.MonitoringEvent.Source).To(Equal("transfer-plugin"))
			Expect(body.MonitoringEvent.Metric).To(Equal(event_ports.ActionStart))

			// Verify the item was inserted into the database
			event1CreatedID = *body.Id
			model := &models.Event{}
			err = db.First(model, event1CreatedID).Error
			Expect(err).ToNot(HaveOccurred())
			Expect(model.ContractID).To(Equal("gfjfgd45555sd5s5s5s5s"))
			Expect(model.ExecutionID).To(Equal("uktdqs4661qsdqs664141"))
			Expect(model.Source).To(Equal("transfer-plugin"))
			Expect(model.Metric).To(Equal("action-start"))
			Expect(model.Log).To(Equal("Root WebApplicationContext: initialization completed in 2758 ms"))
		})

		It("OK - should insert a event stop item into the database", func() {
			// Prepare the HTTP POST request
			reqBody := `{
				"contractId": "gfjfgd45555sd5s5s5s5s",
				"executionId": "uktdqs4661qsdqs664141",
				"monitoringEvent": {
					"source": "transfer-plugin",
					"timestamp": "2025-02-26T14:28:09Z",
					"metric": "action-stop",
					"value": "calculation of statistics",
					"log": "Process Complete"
				}
			}`
			req := httptest.NewRequest(http.MethodPost, "/v1/events", strings.NewReader(reqBody))
			rec := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(rec, req)

			// Verify the HTTP response code
			Expect(rec.Code).To(Equal(http.StatusOK))

			var body event_ports.Event
			err := json.NewDecoder(rec.Body).Decode(&body)
			Expect(err).ToNot(HaveOccurred())
			Expect(body.Id).ShouldNot(BeNil())
			Expect(body.CreatedAt).ShouldNot(BeNil())

			// Verify the item was inserted into the database
			event2CreatedID = *body.Id
			model := &models.Event{}
			err = db.First(model, event2CreatedID).Error
			Expect(err).ToNot(HaveOccurred())
			Expect(model.ContractID).To(Equal("gfjfgd45555sd5s5s5s5s"))
			Expect(model.ExecutionID).To(Equal("uktdqs4661qsdqs664141"))
		})

		It("OK - should fetch 2 events in events list", func() {
			// Prepare the HTTP GET request
			req := httptest.NewRequest(http.MethodGet, "/v1/events?contractId=gfjfgd45555sd5s5s5s5s&executionId=uktdqs4661qsdqs664141", nil)
			rec := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(rec, req)

			// Verify the HTTP response code
			Expect(rec.Code).To(Equal(http.StatusOK))

			var body []event_ports.Event
			err := json.NewDecoder(rec.Body).Decode(&body)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(body)).To(Equal(2))
		})
	})

	Context("Fetch Public Key", func() {
		It("OK - should fetch public key", func() {
			// Prepare the HTTP GET request
			req := httptest.NewRequest(http.MethodGet, "/v1/verifiable-credential/publicKey", nil)
			rec := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(rec, req)

			// Verify the HTTP response code
			Expect(rec.Code).To(Equal(http.StatusOK))

			var body vc_ports.PublicKey
			err := json.NewDecoder(rec.Body).Decode(&body)
			Expect(err).ToNot(HaveOccurred())
			Expect(body.Key).ShouldNot(BeNil())
			Expect(body.Type).To(Equal(vc_ports.Ed25519))
		})
	})

	Context("Sign VC and Fetch VCs List", func() {
		It("OK - should fetch 0 VCs", func() {
			// Prepare the HTTP GET request
			req := httptest.NewRequest(http.MethodGet, "/v1/verifiable-credential", nil)
			rec := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(rec, req)

			// Verify the HTTP response code
			Expect(rec.Code).To(Equal(http.StatusOK))

			var body []vc_ports.VcSigned
			err := json.NewDecoder(rec.Body).Decode(&body)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(body)).To(Equal(0))
		})

		It("OK - should create and sign VC", func() {
			// Prepare the HTTP GET request
			req := httptest.NewRequest(http.MethodPost, "/v1/verifiable-credential/_sign?contractId=gfjfgd45555sd5s5s5s5s&executionId=uktdqs4661qsdqs664141", nil)
			rec := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(rec, req)

			// Verify the HTTP response code
			Expect(rec.Code).To(Equal(http.StatusOK))

			var body vc_ports.VcSigned
			err := json.NewDecoder(rec.Body).Decode(&body)
			Expect(err).ToNot(HaveOccurred())
			Expect(body.CredentialSubject.Id).To(Equal("gfjfgd45555sd5s5s5s5s"))
			Expect(body.CredentialSubject.ExecutionId).To(Equal("uktdqs4661qsdqs664141"))
			Expect(len(body.CredentialSubject.MonitoringEvents)).To(Equal(2))
		})

		It("OK - should fetch 1 signed VC", func() {
			// Prepare the HTTP GET request
			req := httptest.NewRequest(http.MethodGet, "/v1/verifiable-credential", nil)
			rec := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(rec, req)

			// Verify the HTTP response code
			Expect(rec.Code).To(Equal(http.StatusOK))

			var body []vc_ports.VcSigned
			err := json.NewDecoder(rec.Body).Decode(&body)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(body)).To(Equal(1))
		})
	})
})
