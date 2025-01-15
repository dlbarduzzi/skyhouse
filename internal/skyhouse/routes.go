package skyhouse

import "net/http"

func (s *Skyhouse) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/health", s.healthHandler)
	return mux
}
