package middleware

import (
	"PluginTemplate/pkg/powerx/client"
	"net/http"
)

type PluginMiddleware struct {
}

func NewPluginMiddleware() *PluginMiddleware {
	return &PluginMiddleware{}
}

func (m *PluginMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := client.GetServerHeaderAuthorization(r)
		r = r.WithContext(client.WithCtxAuthorization(r.Context(), token))
		next(w, r)
	}
}
