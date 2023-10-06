package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"net/http"
)

type MpMarketStore struct {
	*PowerX
}

func (c *MpMarketStore) ListStoresPage(req *powerxtypes.ListStoresPageRequest) (*powerxtypes.ListStoresPageReply, error) {
	res := &powerxtypes.ListStoresPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/mp/market/stores/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
