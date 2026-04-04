package req

import (
	"net/http"

	"github.com/mukhinfa/golang-advanced/4-order-api/pkg/res"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	var body T
	body, err := Decode[T](r.Body)
	if err != nil {
		res.JSON(*w, http.StatusBadRequest, err.Error())
		return nil, err
	}
	if err := isValid(body); err != nil {
		res.JSON(*w, http.StatusBadRequest, err.Error())
		return nil, err
	}
	return &body, nil
}
