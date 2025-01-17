package skyhouse

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newTestSkyhouse(t *testing.T) *Skyhouse {
	t.Helper()
	return &Skyhouse{
		logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}
}

type testServer struct {
	http.HandlerFunc
}

func newTestServer(t *testing.T, h http.HandlerFunc) *testServer {
	t.Helper()
	return &testServer{h}
}

func (s *testServer) get(t *testing.T, url string) (int, string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	s.ServeHTTP(w, req)

	code := w.Result().StatusCode
	body := string(bytes.TrimSpace(w.Body.Bytes()))

	return code, body
}
