package middleware

import (
	"context"
	"github.com/jhidalgoesp/commons/auth"
	"github.com/jhidalgoesp/commons/validate"
	"github.com/jhidalgoesp/commons/web"
	"net/http"
)

// CurrentUser validates a JWT from the `jwt` cookie.
func CurrentUser() web.Middleware {
	// This is the actual middleware function to be executed.
	m := func(handler web.Handler) web.Handler {

		// Create the handler that will be attached in the middleware chain.
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			cookie, err := r.Cookie("auth")
			if err != nil {
				return validate.NewRequestError(auth.ErrForbidden, http.StatusUnauthorized)
			}

			// Validate the token is signed by us.
			claims, err := auth.Verify(cookie)
			if err != nil {
				return validate.NewRequestError(auth.ErrForbidden, http.StatusUnauthorized)
			}

			// Add claims to the context, so they can be retrieved later.
			ctx = auth.SetClaims(ctx, claims)

			// Call the next handler.
			return handler(ctx, w, r)
		}

		return h
	}

	return m
}
