package handler

import (
	"context"
	"github.com/google/uuid"
	"net/http"
)

func middlewareRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "uuid", uuid.New().String())

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
