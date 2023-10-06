package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"fmt"
	"net/http"
)

type AdminProductSku struct {
	*PowerX
}

func (c *AdminProductSku) ListSKUPage(req *powerxtypes.ListSKUPageRequest) (*powerxtypes.ListSKUPageReply, error) {
	res := &powerxtypes.ListSKUPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/product/skus/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductSku) GetSKU(req *powerxtypes.GetSKURequest) (*powerxtypes.GetSKUReply, error) {
	res := &powerxtypes.GetSKUReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/product/skus/%v", req.SKUId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductSku) CreateSKU(req *powerxtypes.CreateSKURequest) (*powerxtypes.CreateSKUReply, error) {
	res := &powerxtypes.CreateSKUReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/product/skus").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductSku) ConfigSKU(req *powerxtypes.ConfigSKURequest) (*powerxtypes.ConfigSKUReply, error) {
	res := &powerxtypes.ConfigSKUReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/product/skus/config").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductSku) PutSKU(req *powerxtypes.PutSKURequest) (*powerxtypes.PutSKUReply, error) {
	res := &powerxtypes.PutSKUReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/product/skus/%v", req.SKUId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductSku) PatchSKU(req *powerxtypes.PatchSKURequest) (*powerxtypes.PatchSKUReply, error) {
	res := &powerxtypes.PatchSKUReply{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri(fmt.Sprintf("/api/v1/admin/product/skus/%v", req.SKUId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductSku) DeleteSKU(req *powerxtypes.DeleteSKURequest) (*powerxtypes.DeleteSKUReply, error) {
	res := &powerxtypes.DeleteSKUReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/product/skus/%v", req.SKUId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
