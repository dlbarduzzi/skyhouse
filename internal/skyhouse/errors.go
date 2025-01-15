package skyhouse

import (
	"log/slog"
	"net/http"

	"github.com/dlbarduzzi/skyhouse/internal/jsontil"
)

type serverErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (s *Skyhouse) serverError(w http.ResponseWriter, r *http.Request, e error) {
	s.logger.Error(
		e.Error(),
		slog.String("path", r.URL.Path),
		slog.String("method", r.Method),
	)

	res := serverErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error.",
	}

	if err := jsontil.Marshal(w, res, res.Code, nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
