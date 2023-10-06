package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"fmt"
	"net/http"
)

type MpProductArtisan struct {
	*PowerX
}

func (c *MpProductArtisan) ListArtisansPage(req *powerxtypes.ListArtisansPageRequest) (*powerxtypes.ListArtisansPageReply, error) {
	res := &powerxtypes.ListArtisansPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/mp/product/artisans/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpProductArtisan) GetArtisan(req *powerxtypes.GetArtisanRequest) (*powerxtypes.GetArtisanReply, error) {
	res := &powerxtypes.GetArtisanReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/mp/product/artisans/%v", req.ArtisanId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
