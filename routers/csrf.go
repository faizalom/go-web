package routers

import (
	"net/http"

	"github.com/faizalom/go-web/config"

	"github.com/gorilla/csrf"
)

func CSRFRoute(route http.Handler) http.Handler {
	CSRF := csrf.Protect([]byte(config.Cipher), csrf.ErrorHandler(http.HandlerFunc(pageExpiredHandler)))
	SKIP := NewSkipper()
	return SKIP(CSRF(route))
}

type SkipCSRF struct {
	h http.Handler
}

func NewSkipper() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		sk := &SkipCSRF{
			h: h,
		}
		return sk
	}
}

func (sr *SkipCSRF) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/skip/me" {
		r = csrf.UnsafeSkipCheck(r)
	}
	sr.h.ServeHTTP(w, r)
}
