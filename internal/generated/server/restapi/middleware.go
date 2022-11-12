package restapi

import (
	"net/http"
	"runtime/debug"

	loggerpkg "github/user-manager/tools/logger"
)

func loggerMiddleware(handler http.Handler, logger loggerpkg.ILogger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := loggerpkg.EnrichContext(r.Context(), logger)
		r = r.WithContext(ctx)

		handler.ServeHTTP(w, r)
	})
}

func panicRecoveryMiddleware(handler http.Handler, logger loggerpkg.ILogger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if p := recover(); p != nil {
				logger.WithPanic(string(debug.Stack())).Error("panic recovered")
			}
		}()

		handler.ServeHTTP(w, r)
	})
}
