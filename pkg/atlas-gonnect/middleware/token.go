package middleware

import (
	"net/http"

	"github.com/go-enjin/third_party/pkg/atlas-gonnect"
)

type TokenMiddleware struct {
	h     http.Handler
	addon *gonnect.Addon
}

func (h TokenMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: Add isTokenVerificationDisabled
	authHandler := NewAuthenticationMiddleware(h.addon, true)
	authHandler(h.h).ServeHTTP(w, r)
}

func NewTokenMiddleware(addon *gonnect.Addon) func(h http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return TokenMiddleware{handler, addon}
	}
}
