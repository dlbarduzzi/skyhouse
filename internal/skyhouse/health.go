package skyhouse

import (
	"net/http"

	"github.com/dlbarduzzi/skyhouse/internal/jsontil"
)

type healthResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (s *Skyhouse) healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodHead {
		w.WriteHeader(http.StatusOK)
		return
	}

	res := healthResponse{
		Code:    http.StatusOK,
		Message: "API is healthy.",
	}

	if err := jsontil.Marshal(w, res, res.Code, nil); err != nil {
		s.serverError(w, r, err)
		return
	}
}
