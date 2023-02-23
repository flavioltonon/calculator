package middleware

import (
	"net/http"

	"calculator/application/http/router"
	"calculator/domain/valueobject"
	"calculator/shared/logging"

	"github.com/urfave/negroni"
)

func Log(logger logging.Logger) router.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := valueobject.NewRequestID()

			logger.Info("Got request",
				logging.String("http_method", r.Method),
				logging.String("request_uri", r.RequestURI),
				logging.Stringer("request_id", requestID),
			)

			rw := negroni.NewResponseWriter(w)

			next.ServeHTTP(rw, r)

			logger.Info("Returning response",
				logging.Stringer("request_id", requestID),
				logging.Int("http_status_code", rw.Status()),
			)
		})
	}
}
