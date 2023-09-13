package client

import (
	"context"
	"net/http"
)

func withCtxAuthorization(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, "Authorization", token)
}

func fromCtxAuthorization(ctx context.Context) (string, bool) {
	return ctx.Value("Authorization").(string), true
}

func getServerHeaderAuthorization(request *http.Request) string {
	return request.Header.Get("Authorization")
}
