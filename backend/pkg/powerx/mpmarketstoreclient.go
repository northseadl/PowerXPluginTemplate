package powerx

import (
	"PluginTemplate/pkg/powerx/types"

	"net/http"
)

type MpMarketStore struct {
	*PowerX
}

func (c *MpMarketStore) ListStoresPage(req *types.ListStoresPageRequest) (*types.ListStoresPageReply, error) {
	res := &types.ListStoresPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/mp/market/stores/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
