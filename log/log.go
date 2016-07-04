package log

import (
	"net/http"

	"github.com/Finciero/log"
	"github.com/gorilla/context"
)

// LoggerFromRequest checks if a logger exists
// in the gorilla context, if does, it cast it and then
// returns it, if don't, returns a NOOPLogger
func LoggerFromRequest(r *http.Request, keyvals ...interface{}) log.Logger {
	if logger, ok := context.GetOk(r, "logger"); ok {
		return logger.(log.Logger).With(keyvals...)
	}
	return log.NewNOOPContext()
}
