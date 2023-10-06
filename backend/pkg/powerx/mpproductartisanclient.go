package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"
	"fmt"
	"net/http"
)

type MpProductArtisan struct {
	*PowerX
}

func (c *MpProductArtisan) ListArtisansPage(ctx context.Context, req *powerxtypes.ListArtisansPageRequest) (*powerxtypes.ListArtisansPageReply, error) {
	res := &powerxtypes.ListArtisansPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/mp/product/artisans/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpProductArtisan) GetArtisan(ctx context.Context, req *powerxtypes.GetArtisanRequest) (*powerxtypes.GetArtisanReply, error) {
	res := &powerxtypes.GetArtisanReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/mp/product/artisans/%v", req.ArtisanId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
