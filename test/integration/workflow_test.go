package integration_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/dawex/vc-generator/internal/common/db/models"
	compliancelogs_ports "github.com/dawex/vc-generator/internal/services/compliance-logs/ports"
	negotiationcontracts_ports "github.com/dawex/vc-generator/internal/services/negotiation-contracts/ports"
	vc_ports "github.com/dawex/vc-generator/internal/services/verifiable-credential/ports"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Workflow IT", func() {
	var (
		compliancelogs1CreatedID uuid.UUID
		compliancelogs2CreatedID uuid.UUID
	)

	Context("Insert Negotiation Contract and Fetch Negotiation Contracts List", func() {
		It("OK - should fetch an empty negotiationcontracts list. No records yet.", func() {
			// Prepare the HTTP GET request
			req := httptest.NewRequest(http.MethodGet, "/v1/negotiation-contracts", nil)
			rec := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(rec, req)

			// Verify the HTTP response code
			Expect(rec.Code).To(Equal(http.StatusOK))

			var body []negotiationcontracts_ports.NegotiationContract
			err := json.NewDecoder(rec.Body).Decode(&body)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(body)).To(Equal(0))
		})

		It("OK - should insert a negotiationcontract item into the database", func() {
			// Prepare the HTTP POST request
			reqBody := `{
				"type": "Contract Type",
				"consumer_id": "Consumer ID",
				"producer_id": "Producer ID",
				"data_processing_workflow_object": "Object:6158f68sdf165ds615fs",
				"natural_language_document": "EN",
				"resource_description_object": {
					"test": "test"
				},
				"odrl_policy": {
					"@type": "odrl:Offer",
					"odrl:permission": [
						{
							"odrl:action": "odrl:derive",
							"odrl:constraint": [
								{
									"odrl:rightOperand": "xls",
									"odrl:leftOperand": "odrl:fileFormat",
									"odrl:operator": "odrl:eq"
								}
							]
						}
					]
				},
				"id": "contractId123gt",
				"title": "Contract Tiltle",
				"negotiation_id": "Negotiation ID",
				"created_at": "2025-03-30T20:29:45Z",
				"updated_at": "2025-04-03T22:29:45Z"
			}`
			req := httptest.NewRequest(http.MethodPost, "/v1/negotiation-contracts", strings.NewReader(reqBody))
			rec := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(rec, req)

			// Verify the HTTP response code
			Expect(rec.Code).To(Equal(http.StatusOK))

			var body negotiationcontracts_ports.NegotiationContract
			err := json.NewDecoder(rec.Body).Decode(&body)
			Expect(err).ToNot(HaveOccurred())
			Expect(body.Id).ShouldNot(BeNil())
			Expect(body.CreatedAt).ShouldNot(BeNil())
			Expect(body.UpdatedAt).ShouldNot(BeNil())
			Expect(body.ResourceDescriptionObject).ShouldNot(BeNil())
			Expect(body.OdrlPolicy).ShouldNot(BeNil())
			Expect(*body.NegotiationId).To(Equal("Negotiation ID"))
			Expect(*body.Title).To(Equal("Contract Tiltle"))
			Expect(body.ConsumerId).To(Equal("Consumer ID"))
			Expect(body.ProducerId).To(Equal("Producer ID"))
			Expect(body.DataProcessingWorkflowObject).To(Equal("Object:6158f68sdf165ds615fs"))
			Expect(body.NaturalLanguageDocument).To(Equal("EN"))

			// Verify the item was inserted into the database
			model := &models.NegotiationContract{}
			model.ID = body.Id
			err = db.First(model).Error
			Expect(err).ToNot(HaveOccurred())
			Expect(model.CreatedAt).ShouldNot(BeNil())
			Expect(model.UpdatedAt).ShouldNot(BeNil())
			Expect(model.ResourceDescriptionObject).ShouldNot(BeNil())
			Expect(model.OdrlPolicy).ShouldNot(BeNil())
			Expect(*model.NegotiationID).To(Equal("Negotiation ID"))
			Expect(*model.Title).To(Equal("Contract Tiltle"))
			Expect(model.ConsumerID).To(Equal("Consumer ID"))
			Expect(model.ProducerID).To(Equal("Producer ID"))
			Expect(model.DataProcessingWorkflowObject).To(Equal("Object:6158f68sdf165ds615fs"))
			Expect(model.NaturalLanguageDocument).To(Equal("EN"))
		})

		It("OK - should fetch 1 negotiationcontract in negotiationcontracts list", func() {
			// Prepare the HTTP GET request
			req := httptest.NewRequest(http.MethodGet, "/v1/negotiation-contracts", nil)
			rec := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(rec, req)

			// Verify the HTTP response code
			Expect(rec.Code).To(Equal(http.StatusOK))

			var body []negotiationcontracts_ports.NegotiationContract
			err := json.NewDecoder(rec.Body).Decode(&body)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(body)).To(Equal(1))
		})
	})

	Context("Insert Compliance Logs and Fetch Compliance Logs List", func() {
		It("OK - should fetch an empty compliancelogs list. No records yet.", func() {
			// Prepare the HTTP GET request
			req := httptest.NewRequest(http.MethodGet, "/v1/compliance-logs?contractId=contractId123gt&executionId=executionId123gt", nil)
			rec := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(rec, req)

			// Verify the HTTP response code
			Expect(rec.Code).To(Equal(http.StatusOK))

			var body []compliancelogs_ports.ComplianceLog
			err := json.NewDecoder(rec.Body).Decode(&body)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(body)).To(Equal(0))
		})

		It("OK - should insert a compliancelogs start item into the database", func() {
			// Prepare the HTTP POST request
			reqBody := `{
				"contractId": "contractId123gt",
				"executionId": "executionId123gt",
				"monitoringEvent": {
					"log": "LOG: start of reader:read-data",
					"metric": "action-start",
					"params": [
						"datafile"
					],
					"result": null,
					"source": "reader",
					"value": "read-data",
					"groups": "adm",
					"timestamp": "2025-03-30T20:29:45Z"
				},
				"complianceLogs": [
					{
						"log_lvl": "ALERT",
						"log_msg": "rule violation",
						"rule_num": 1,
						"rule_context": "*",
						"rule_expr": "$name in {'a', 'b'}"
					},
					{
						"log_lvl": "INFO",
						"log_msg": "rule match",
						"rule_num": 5,
						"rule_context": "*",
						"rule_expr": "'adm' in $groups"
					}
				]
			}`
			req := httptest.NewRequest(http.MethodPost, "/v1/compliance-logs", strings.NewReader(reqBody))
			rec := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(rec, req)

			// Verify the HTTP response code
			Expect(rec.Code).To(Equal(http.StatusOK))

			var body compliancelogs_ports.ComplianceLog
			err := json.NewDecoder(rec.Body).Decode(&body)
			Expect(err).ToNot(HaveOccurred())
			Expect(body.Id).ShouldNot(BeNil())
			Expect(body.CreatedAt).ShouldNot(BeNil())
			Expect(body.ComplianceLogs).ShouldNot(BeNil())
			Expect(body.ContractId).To(Equal("contractId123gt"))
			Expect(body.ExecutionId).To(Equal("executionId123gt"))
			Expect(body.MonitoringEvent.Source).To(Equal("reader"))
			Expect(body.MonitoringEvent.Metric).To(Equal("action-start"))
			Expect(body.MonitoringEvent.Log).To(Equal("LOG: start of reader:read-data"))
			Expect(body.MonitoringEvent.Value).To(Equal("read-data"))
			Expect(body.MonitoringEvent.Groups).To(Equal("adm"))

			// Verify the item was inserted into the database
			compliancelogs1CreatedID = *body.Id
			model := &models.ComplianceLog{}
			err = db.First(model, compliancelogs1CreatedID).Error
			Expect(err).ToNot(HaveOccurred())
			Expect(model.ContractID).To(Equal("contractId123gt"))
			Expect(model.ExecutionID).To(Equal("executionId123gt"))
			Expect(model.Source).To(Equal("reader"))
			Expect(model.Metric).To(Equal("action-start"))
			Expect(model.Log).To(Equal("LOG: start of reader:read-data"))
		})

		It("OK - should insert a compliancelog stop item into the database", func() {
			// Prepare the HTTP POST request
			reqBody := `{
				"contractId": "contractId123gt",
				"executionId": "executionId123gt",
				"monitoringEvent": {
					"log": "LOG: end of integrator:integrate",
					"metric": "action-stop",
					"params": [],
					"result": null,
					"source": "integrator",
					"value": "integrate",
					"groups": "adm",
					"timestamp": "2025-03-30T20:29:54Z"
				},
				"complianceLogs": [
					{
						"log_lvl": "ALERT",
						"log_msg": "rule violation",
						"rule_num": 7,
						"rule_context": "integrate",
						"rule_expr": "$timestamp < '2025-3-19T12:47:54Z'"
					}
				]
			}`
			req := httptest.NewRequest(http.MethodPost, "/v1/compliance-logs", strings.NewReader(reqBody))
			rec := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(rec, req)

			// Verify the HTTP response code
			Expect(rec.Code).To(Equal(http.StatusOK))

			var body compliancelogs_ports.ComplianceLog
			err := json.NewDecoder(rec.Body).Decode(&body)
			Expect(err).ToNot(HaveOccurred())
			Expect(body.Id).ShouldNot(BeNil())
			Expect(body.CreatedAt).ShouldNot(BeNil())

			// Verify the item was inserted into the database
			compliancelogs2CreatedID = *body.Id
			model := &models.ComplianceLog{}
			err = db.First(model, compliancelogs2CreatedID).Error
			Expect(err).ToNot(HaveOccurred())
			Expect(model.ContractID).To(Equal("contractId123gt"))
			Expect(model.ExecutionID).To(Equal("executionId123gt"))
		})

		It("OK - should fetch 2 compliancelogs in compliancelogs list", func() {
			// Prepare the HTTP GET request
			req := httptest.NewRequest(http.MethodGet, "/v1/compliance-logs?contractId=contractId123gt&executionId=executionId123gt", nil)
			rec := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(rec, req)

			// Verify the HTTP response code
			Expect(rec.Code).To(Equal(http.StatusOK))

			var body []compliancelogs_ports.ComplianceLog
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
			req := httptest.NewRequest(http.MethodPost, "/v1/verifiable-credential/_sign?contractId=contractId123gt&executionId=executionId123gt", nil)
			rec := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(rec, req)

			// Verify the HTTP response code
			Expect(rec.Code).To(Equal(http.StatusOK))

			var body vc_ports.VcSigned
			err := json.NewDecoder(rec.Body).Decode(&body)
			Expect(err).ToNot(HaveOccurred())
			Expect(body.CredentialSubject.Id).To(Equal("contractId123gt"))
			Expect(body.CredentialSubject.ExecutionId).To(Equal("executionId123gt"))
			Expect(len(body.CredentialSubject.ComplianceAudit)).To(Equal(2))
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
