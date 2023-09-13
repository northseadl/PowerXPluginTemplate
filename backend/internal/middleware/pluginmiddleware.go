package middleware

import (
	"net/http"
)

type PluginMiddleware struct {
}

func NewPluginMiddleware() *PluginMiddleware {
	return &PluginMiddleware{}
}

func (m *PluginMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := server.GetServerHeaderAuthorization(r)
		r = r.WithContext(server.WithCtxAuthorization(r.Context(), token))
		next(w, r)
	}
}
