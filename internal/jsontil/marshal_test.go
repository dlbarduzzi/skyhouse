package jsontil

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMarshal(t *testing.T) {
	t.Parallel()

	res := map[string]interface{}{"foo": "bar"}

	headerKey := "Foo"
	headerValue := "Bar"

	headers := http.Header{}
	headers.Add(headerKey, headerValue)

	w := httptest.NewRecorder()

	err := Marshal(w, res, http.StatusOK, headers)
	if err != nil {
		t.Errorf("expected error to be nil; got %v", err)
	}

	if w.Code != http.StatusOK {
		t.Errorf("expected status code to be %d; got %d", http.StatusOK, w.Code)
	}

	wantHeaderValue := w.Header().Get(headerKey)
	if wantHeaderValue != headerValue {
		t.Errorf("expected header %s to be %s; got %s", headerKey, wantHeaderValue, headerValue)
	}
}

func TestMarshalError(t *testing.T) {
	t.Parallel()

	data := map[string]interface{}{"foo": "bar"}
	data["new-foo"] = data

	w := httptest.NewRecorder()
	err := Marshal(w, data, 0, nil)
	want := "json: unsupported value: encountered a cycle via map[string]interface {}"

	if err == nil || err.Error() != want {
		t.Errorf("expected error to be %v; got %v", want, err)
	}
}
