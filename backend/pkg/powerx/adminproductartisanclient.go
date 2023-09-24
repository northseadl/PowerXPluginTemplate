package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type AdminProductArtisan struct {
	*PowerX
}

func (c *AdminProductArtisan) ListArtisansPage(req *types.ListArtisansPageRequest) (*types.ListArtisansPageReply, error) {
	res := &types.ListArtisansPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/product/artisans/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductArtisan) GetArtisan(req *types.GetArtisanRequest) (*types.GetArtisanReply, error) {
	res := &types.GetArtisanReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/product/artisans/%v", req.ArtisanId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductArtisan) CreateArtisan(req *types.CreateArtisanRequest) (*types.CreateArtisanReply, error) {
	res := &types.CreateArtisanReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/product/artisans").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductArtisan) PutArtisan(req *types.PutArtisanRequest) (*types.PutArtisanReply, error) {
	res := &types.PutArtisanReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/product/artisans/%v", req.ArtisanId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductArtisan) DeleteArtisan(req *types.DeleteArtisanRequest) (*types.DeleteArtisanReply, error) {
	res := &types.DeleteArtisanReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/product/artisans/%v", req.ArtisanId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
