package rest

import (
	"errors"
	"net/http"
	"strconv"

	apiErrors "pokeapi/api/internal/errors"
)

func Error(w http.ResponseWriter, err error) {
	var e apiErrors.ApiError
	if ok := errors.As(err, &e); !ok {
		e = apiErrors.ErrUnhandled
	}
	http.Error(w, e.Error(), e.Status())
}

func IDFromParams(r *http.Request) (int, error) {
	params := r.URL.Query()["id"]
	if len(params) != 1 {
		return 0, apiErrors.ErrBadRequest
	}
	idStr := params[0]
	if idStr == "" {
		return 0, apiErrors.ErrBadRequest
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, apiErrors.ErrBadRequest
	}
	return int(id), nil
}
