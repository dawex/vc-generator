package server

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

func BindAndValidate(r *http.Request, binder render.Binder) (int, error) {
	if err := render.Bind(r, binder); err != nil {
		return http.StatusInternalServerError, err
	}

	return Validate(binder)
}

func Validate(s interface{}) (int, error) {
	validate := validator.New()
	if err := validate.Struct(s); err != nil {
		return http.StatusBadRequest, err
	}

	return 0, nil
}
