package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"
	"fmt"
	"net/http"
)

type AdminMarketMedia struct {
	*PowerX
}

func (c *AdminMarketMedia) ListMediasPage(ctx context.Context, req *powerxtypes.ListMediasPageRequest) (*powerxtypes.ListMediasPageReply, error) {
	res := &powerxtypes.ListMediasPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/admin/market/medias/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminMarketMedia) CreateMedia(ctx context.Context, req *powerxtypes.CreateMediaRequest) (*powerxtypes.CreateMediaReply, error) {
	res := &powerxtypes.CreateMediaReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/admin/market/medias").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminMarketMedia) UpdateMedia(ctx context.Context, req *powerxtypes.UpdateMediaRequest) (*powerxtypes.UpdateMediaReply, error) {
	res := &powerxtypes.UpdateMediaReply{}
	err := c.H.Df().Method(http.MethodPut).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/market/medias/%v", req.MediaId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminMarketMedia) GetMedia(ctx context.Context, req *powerxtypes.GetMediaRequest) (*powerxtypes.GetMediaReply, error) {
	res := &powerxtypes.GetMediaReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/market/medias/%v", req.MediaId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminMarketMedia) DeleteMedia(ctx context.Context, req *powerxtypes.DeleteMediaRequest) (*powerxtypes.DeleteMediaReply, error) {
	res := &powerxtypes.DeleteMediaReply{}
	err := c.H.Df().Method(http.MethodDelete).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/market/medias/%v", req.MediaId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
