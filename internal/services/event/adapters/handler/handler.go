package handler

import (
	"net/http"

	"github.com/dawex/vc-generator/internal/common/server"
	"github.com/dawex/vc-generator/internal/services/event/ports"
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

// List Dataset Execution Events
// (GET /events)
func (h *Handler) GetEvents(w http.ResponseWriter, r *http.Request, params ports.GetEventsParams) {
	// Validate query params
	validate := validator.New()
	if err := validate.Struct(params); err != nil {
		server.RespondError(w, r, err, http.StatusBadRequest)
		return
	}

	// Call service
	modelsFound, err := h.service.GetEvents(r.Context(), params.ContractId, params.ExecutionId)
	if err != nil {
		server.RespondError(w, r, err, http.StatusInternalServerError)
		return
	}

	// Map to response body
	var elems []*ports.Event
	for _, model := range modelsFound {
		elems = append(elems, modelToEntity(&model))
	}

	render.Status(r, http.StatusOK)
	render.RenderList(w, r, ports.ConvertEventListRenders(elems))
}

// Save Dataset Execution Event
// (POST /events)
func (h *Handler) SaveEvent(w http.ResponseWriter, r *http.Request) {
	// Decode & Validate request body
	var body ports.EventIn
	if code, err := server.BindAndValidate(r, &body); err != nil {
		server.RespondError(w, r, err, code)
		return
	}

	// Map request body to model
	model := entityToModel(&body)

	// Call service
	modelCreated, err := h.service.SaveEvent(r.Context(), model)
	if err != nil {
		server.RespondError(w, r, err, http.StatusInternalServerError)
		return
	}

	// Map model to response body
	respBody := modelToEntity(modelCreated)

	server.RespondWithBody(w, r, respBody, http.StatusOK)
}
