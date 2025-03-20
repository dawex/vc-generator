package handler

import (
	"encoding/hex"
	"net/http"

	"github.com/dawex/vc-generator/internal/common/server"
	"github.com/dawex/vc-generator/internal/services/verifiable-credential/ports"
	"github.com/go-chi/render"
	"github.com/go-playground/validator"
)

type Handler struct {
	service ports.Service
}

func NewHandler(service ports.Service) *Handler {
	return &Handler{
		service: service,
	}
}

// List Signed Verifiable Credentials
// (GET /verifiable-credential)
func (h *Handler) ListVerifiableCredentials(w http.ResponseWriter, r *http.Request) {
	// Call service
	modelsFound, err := h.service.ListVerifiableCredentials(r.Context())
	if err != nil {
		server.RespondError(w, r, err, http.StatusInternalServerError)
		return
	}

	// Map to response body
	var elems []*ports.VcSigned
	for _, model := range modelsFound {
		elem, err := modelToEntity(&model)
		if err != nil {
			server.RespondError(w, r, err, http.StatusInternalServerError)
			return
		}
		elems = append(elems, elem)
	}

	render.Status(r, http.StatusOK)
	render.RenderList(w, r, ports.ConvertVcSignedListRenders(elems))
}

// Sign Verifiable Credential for Dataset Execution
// (POST /verifiable-credential/_sign)
func (h *Handler) SignVerifiableCredential(w http.ResponseWriter, r *http.Request, params ports.SignVerifiableCredentialParams) {
	// Validate query params
	validate := validator.New()
	if err := validate.Struct(params); err != nil {
		server.RespondError(w, r, err, http.StatusBadRequest)
		return
	}

	// Call service
	vcSigned, err := h.service.SignVerifiableCredential(r.Context(), params.ContractId, params.ExecutionId)
	if err != nil {
		server.RespondError(w, r, err, http.StatusInternalServerError)
		return
	}

	// Map model to response body
	respBody, err := modelToEntity(vcSigned)
	if err != nil {
		server.RespondError(w, r, err, http.StatusInternalServerError)
		return
	}

	server.RespondWithBody(w, r, respBody, http.StatusOK)
}

// Get Public Key
// (GET /verifiable-credential/publicKey)
func (h *Handler) GetPublicKey(w http.ResponseWriter, r *http.Request) {
	// Call service
	publicKey, err := h.service.GetPublicKey(r.Context())
	if err != nil {
		server.RespondError(w, r, err, http.StatusInternalServerError)
		return
	}

	// Map model to response
	respData := ports.PublicKey{
		Type: ports.Ed25519,
		Key:  hex.EncodeToString(*publicKey),
	}

	server.RespondWithBody(w, r, respData, http.StatusOK)
}
