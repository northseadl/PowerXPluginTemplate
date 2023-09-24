package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type MpProductArtisan struct {
	*PowerX
}

func (c *MpProductArtisan) ListArtisansPage(req *types.ListArtisansPageRequest) (*types.ListArtisansPageReply, error) {
	res := &types.ListArtisansPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/mp/product/artisans/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpProductArtisan) GetArtisan(req *types.GetArtisanRequest) (*types.GetArtisanReply, error) {
	res := &types.GetArtisanReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/mp/product/artisans/%v", req.ArtisanId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
