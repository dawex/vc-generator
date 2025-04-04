package handler

import (
	"net/http"

	"github.com/dawex/vc-generator/internal/common/server"
	"github.com/dawex/vc-generator/internal/services/negotiation-contracts/ports"
	"github.com/go-chi/render"
)

type Handler struct {
	service ports.Service
}

func NewHandler(service ports.Service) *Handler {
	return &Handler{
		service: service,
	}
}

// List Negotiation Contracts
// (GET /negotiation-contracts)
func (h *Handler) ListNegotiationContracts(w http.ResponseWriter, r *http.Request) {
	// Call service
	modelsFound, err := h.service.ListNegotiationContracts(r.Context())
	if err != nil {
		server.RespondError(w, r, err, http.StatusInternalServerError)
		return
	}

	// Map to response body
	var elems []*ports.NegotiationContract
	for _, model := range modelsFound {
		elem, err := modelToEntity(&model)
		if err != nil {
			server.RespondError(w, r, err, http.StatusInternalServerError)
			return
		}
		elems = append(elems, elem)
	}

	render.Status(r, http.StatusOK)
	render.RenderList(w, r, ports.ConvertNegotiationContractListRenders(elems))
}

// Save Negotiation Contract
// (POST /negotiation-contracts)
func (h *Handler) SaveNegotiationContract(w http.ResponseWriter, r *http.Request) {
	// Decode & Validate request body
	var body ports.NegotiationContract
	if code, err := server.BindAndValidate(r, &body); err != nil {
		server.RespondError(w, r, err, code)
		return
	}

	// Map request body to model
	model, err := entityToModel(&body)
	if err != nil {
		server.RespondError(w, r, err, http.StatusBadRequest)
		return
	}

	// Call service
	modelCreated, err := h.service.SaveNegotiationContract(r.Context(), model)
	if err != nil {
		server.RespondError(w, r, err, http.StatusInternalServerError)
		return
	}

	// Map model to response body
	respBody, err := modelToEntity(modelCreated)
	if err != nil {
		server.RespondError(w, r, err, http.StatusInternalServerError)
		return
	}

	server.RespondWithBody(w, r, respBody, http.StatusOK)
}
