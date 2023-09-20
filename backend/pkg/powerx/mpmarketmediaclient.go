package powerx

import (
	"PluginTemplate/pkg/powerx/types"

	"net/http"
)

type MpMarketMedia struct {
	*PowerX
}

func (c *MpMarketMedia) ListMediasPage(req *types.ListMediasPageRequest) (*types.ListMediasPageReply, error) {
	res := &types.ListMediasPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/mp/market/medias/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}