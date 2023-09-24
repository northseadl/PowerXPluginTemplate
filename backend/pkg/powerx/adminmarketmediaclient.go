package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type AdminMarketMedia struct {
	*PowerX
}

func (c *AdminMarketMedia) ListMediasPage(req *types.ListMediasPageRequest) (*types.ListMediasPageReply, error) {
	res := &types.ListMediasPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/market/medias/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminMarketMedia) CreateMedia(req *types.CreateMediaRequest) (*types.CreateMediaReply, error) {
	res := &types.CreateMediaReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/market/medias").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminMarketMedia) UpdateMedia(req *types.UpdateMediaRequest) (*types.UpdateMediaReply, error) {
	res := &types.UpdateMediaReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/market/medias/%v", req.MediaId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminMarketMedia) GetMedia(req *types.GetMediaRequest) (*types.GetMediaReply, error) {
	res := &types.GetMediaReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/market/medias/%v", req.MediaId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminMarketMedia) DeleteMedia(req *types.DeleteMediaRequest) (*types.DeleteMediaReply, error) {
	res := &types.DeleteMediaReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/market/medias/%v", req.MediaId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
