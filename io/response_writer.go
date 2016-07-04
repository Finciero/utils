package io

// JSONResponseWriter ...
import (
	"encoding/json"
	"net/http"

	"github.com/Finciero/errors"
	"github.com/Finciero/log"
)

type JSONResponseWriter struct {
	http.ResponseWriter
	log.Logger
}

// NewJSONResponseWriter ...
func NewJSONResponseWriter(w http.ResponseWriter, l log.Logger) *JSONResponseWriter {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	return &JSONResponseWriter{
		ResponseWriter: w,
		Logger:         l,
	}
}

// WriteError ...
func (w *JSONResponseWriter) WriteError(err error) {
	v := errors.BuildError(err)
	w.Write(v.Code(), v)
}

// Write ...
func (w *JSONResponseWriter) Write(code int, v interface{}) {
	w.WriteHeader(code)

	err := json.NewEncoder(w.ResponseWriter).Encode(v)

	if err != nil {
		w.Logger.Error(err,
			"json_response_writer", "failed to econde message",
			"value", v,
		)
	}
}
