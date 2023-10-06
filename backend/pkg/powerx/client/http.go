package client

import (
	"github.com/artisancloud/httphelper"
	"github.com/artisancloud/httphelper/client"
	"github.com/artisancloud/httphelper/dataflow"
	"net/http"
	"time"
)

type PClient struct {
	endpoint string
	H        httphelper.Helper
}

func NewPClient(endpoint string, debug bool) *PClient {
	conf := &httphelper.Config{
		Config: &client.Config{
			Timeout: 30 * time.Second,
		},
		BaseUrl: endpoint,
	}

	helper, err := httphelper.NewRequestHelper(conf)
	if err != nil {
		panic(err)
	}
	helper.WithMiddleware(httphelper.HttpDebugMiddleware(debug))
	helper.WithMiddleware(func(handle dataflow.RequestHandle) dataflow.RequestHandle {
		return func(request *http.Request, response *http.Response) error {
			if token, ok := FromCtxAuthorization(request.Context()); ok {
				request.Header.Set("Authorization", token)
			}
			return handle(request, response)
		}
	})
	return &PClient{
		endpoint: endpoint,
		H:        helper,
	}
}
