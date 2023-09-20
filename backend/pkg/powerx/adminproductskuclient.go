package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type AdminProductSku struct {
	*PowerX
}

func (c *AdminProductSku) ListSKUPage(req *types.ListSKUPageRequest) (*types.ListSKUPageReply, error) {
	res := &types.ListSKUPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/product/skus/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductSku) GetSKU(req *types.GetSKURequest) (*types.GetSKUReply, error) {
	res := &types.GetSKUReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/product/skus/%s", req.SKUId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductSku) CreateSKU(req *types.CreateSKURequest) (*types.CreateSKUReply, error) {
	res := &types.CreateSKUReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/product/skus").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductSku) ConfigSKU(req *types.ConfigSKURequest) (*types.ConfigSKUReply, error) {
	res := &types.ConfigSKUReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/product/skus/config").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductSku) PutSKU(req *types.PutSKURequest) (*types.PutSKUReply, error) {
	res := &types.PutSKUReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/product/skus/%s", req.SKUId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductSku) PatchSKU(req *types.PatchSKURequest) (*types.PatchSKUReply, error) {
	res := &types.PatchSKUReply{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri(fmt.Sprintf("/api/v1/admin/product/skus/%s", req.SKUId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductSku) DeleteSKU(req *types.DeleteSKURequest) (*types.DeleteSKUReply, error) {
	res := &types.DeleteSKUReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/product/skus/%s", req.SKUId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
