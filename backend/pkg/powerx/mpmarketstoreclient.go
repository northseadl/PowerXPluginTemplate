package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"

	"net/http"
)

type MpMarketStore struct {
	*PowerX
}

func (c *MpMarketStore) ListStoresPage(ctx context.Context, req *powerxtypes.ListStoresPageRequest) (*powerxtypes.ListStoresPageReply, error) {
	res := &powerxtypes.ListStoresPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/mp/market/stores/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
