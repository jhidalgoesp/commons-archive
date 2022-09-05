package middleware

import (
	"context"
	"github.com/jhidalgoesp/commons/auth"
	"github.com/jhidalgoesp/commons/validate"
	"github.com/jhidalgoesp/commons/web"
	"net/http"
)

// Authenticate validate a user is set inside the context.
func Authenticate() Middleware {
	// This is the actual middleware function to be executed.
	m := func(handler web.Handler) web.Handler {
		// Create the handler that will be attached in the middleware chain.
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			_, err := auth.GetClaims(ctx)
			if err != nil {
				return validate.NewRequestError(auth.ErrForbidden, http.StatusForbidden)
			}

			// Call the next handler.
			return handler(ctx, w, r)
		}
		return h
	}
	return m
}
