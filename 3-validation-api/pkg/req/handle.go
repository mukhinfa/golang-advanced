package req

import (
	"net/http"

	"github.com/mukhinfa/golang-advanced/3-validation-api/pkg/res"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	var body T
	if body, err := decode[T](r.Body); err != nil {
		res.JSONResponse(*w, http.StatusBadRequest, err)
		return &body, err
	}
	if err := isValid(body); err != nil {
		res.JSONResponse(*w, http.StatusBadRequest, err)
		return &body, err
	}
	return &body, nil
}
