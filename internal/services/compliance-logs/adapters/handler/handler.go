package handler

import (
	"net/http"

	"github.com/dawex/vc-generator/internal/common/server"
	"github.com/dawex/vc-generator/internal/services/compliance-logs/ports"
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

// List Compliance Logs
// (GET /compliance-logs)
func (h *Handler) ListComplianceLogs(w http.ResponseWriter, r *http.Request, params ports.ListComplianceLogsParams) {
	// Validate query params
	validate := validator.New()
	if err := validate.Struct(params); err != nil {
		server.RespondError(w, r, err, http.StatusBadRequest)
		return
	}

	// Call service
	modelsFound, err := h.service.ListComplianceLogs(r.Context(), params.ContractId, params.ExecutionId)
	if err != nil {
		server.RespondError(w, r, err, http.StatusInternalServerError)
		return
	}

	// Map to response body
	var elems []*ports.ComplianceLog
	for _, model := range modelsFound {
		elem, err := modelToEntity(&model)
		if err != nil {
			server.RespondError(w, r, err, http.StatusInternalServerError)
			return
		}
		elems = append(elems, elem)
	}

	render.Status(r, http.StatusOK)
	render.RenderList(w, r, ports.ConvertComplianceLogListRenders(elems))
}

// Save Compliance Log
// (POST /compliance-logs)
func (h *Handler) SaveComplianceLog(w http.ResponseWriter, r *http.Request) {
	// Decode & Validate request body
	var body ports.ComplianceLogIn
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
	modelCreated, err := h.service.SaveComplianceLog(r.Context(), model)
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
