package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"fmt"
	"net/http"
)

type AdminProductArtisan struct {
	*PowerX
}

func (c *AdminProductArtisan) ListArtisansPage(req *powerxtypes.ListArtisansPageRequest) (*powerxtypes.ListArtisansPageReply, error) {
	res := &powerxtypes.ListArtisansPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/product/artisans/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductArtisan) GetArtisan(req *powerxtypes.GetArtisanRequest) (*powerxtypes.GetArtisanReply, error) {
	res := &powerxtypes.GetArtisanReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/product/artisans/%v", req.ArtisanId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductArtisan) CreateArtisan(req *powerxtypes.CreateArtisanRequest) (*powerxtypes.CreateArtisanReply, error) {
	res := &powerxtypes.CreateArtisanReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/product/artisans").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductArtisan) PutArtisan(req *powerxtypes.PutArtisanRequest) (*powerxtypes.PutArtisanReply, error) {
	res := &powerxtypes.PutArtisanReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/product/artisans/%v", req.ArtisanId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductArtisan) DeleteArtisan(req *powerxtypes.DeleteArtisanRequest) (*powerxtypes.DeleteArtisanReply, error) {
	res := &powerxtypes.DeleteArtisanReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/product/artisans/%v", req.ArtisanId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
