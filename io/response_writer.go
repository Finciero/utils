package io

// JSONResponseWriter ...
import (
	"encoding/json"
	"net/http"

	"github.com/Finciero/errors"
	"github.com/Finciero/log"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// JSONResponseWriter ...
type JSONResponseWriter struct {
	http.ResponseWriter
	Logger log.Logger
	Span   opentracing.Span
}

// NewJSONResponseWriter ...
func NewJSONResponseWriter(w http.ResponseWriter, l log.Logger, s opentracing.Span) *JSONResponseWriter {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	return &JSONResponseWriter{
		ResponseWriter: w,
		Logger:         l,
		Span:           s,
	}
}

// WriteError ...
func (w *JSONResponseWriter) WriteError(err error) {
	v := errors.BuildError(err)
	if v.StatusCode == errors.StatusInternalServerError {
		w.Logger.Error(v)
	}

	ext.Error.Set(w.Span, true)

	w.Write(v.Code(), v)
}

// Write ...
func (w *JSONResponseWriter) Write(code int, v interface{}) {
	w.WriteHeader(code)
	ext.HTTPStatusCode.Set(w.Span, uint16(code))

	err := json.NewEncoder(w.ResponseWriter).Encode(v)

	if err != nil {
		w.Logger.Error(err,
			"json_response_writer", "failed to econde message",
			"value", v,
		)
	}
}
