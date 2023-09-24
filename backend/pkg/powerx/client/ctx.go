package client

import (
	"context"
	"net/http"
)

func WithCtxAuthorization(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, "Authorization", token)
}

func FromCtxAuthorization(ctx context.Context) (string, bool) {
	return ctx.Value("Authorization").(string), true
}

func GetServerHeaderAuthorization(request *http.Request) string {
	return request.Header.Get("Authorization")
}
