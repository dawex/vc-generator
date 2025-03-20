package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/pkg/errors"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

const fieldErrMsg = "Field validation failed on the '%s' tag"

type ErrorResponse struct {
	ErrorMessage   string         `json:"error"`
	Details        []ErrorDetails `json:"details,omitempty"`
	Detail         string         `json:"detail,omitempty"`
	HTTPStatusCode int            `json:"status"`
}

func (a *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type ErrorDetails struct {
	Field string `json:"field"`
	Msg   string `json:"msg"`
}

func RespondError(w http.ResponseWriter, r *http.Request, err error, code int) {
	log.Error().Stack().Err(err).Msg("Error during process")

	if code == http.StatusBadRequest {

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorDetails, len(ve))
			for i, fe := range ve {
				out[i] = ErrorDetails{fe.Field(), fmt.Sprintf(fieldErrMsg, fe.Tag())}
			}

			RespondWithBody(w, r, &ErrorResponse{
				HTTPStatusCode: code,
				ErrorMessage:   http.StatusText(code),
				Details:        out,
			}, code)
		}
	} else {
		RespondWithBody(w, r, &ErrorResponse{
			HTTPStatusCode: code,
			ErrorMessage:   http.StatusText(code),
		}, code)
	}
}

func RespondWithBody(w http.ResponseWriter, r *http.Request, respData render.Renderer, code int) {
	render.Status(r, code)
	if err := render.Render(w, r, respData); err != nil {
		RespondError(w, r, err, http.StatusInternalServerError)
	}
}
