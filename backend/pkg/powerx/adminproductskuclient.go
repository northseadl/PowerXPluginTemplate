package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"
	"fmt"
	"net/http"
)

type AdminProductSku struct {
	*PowerX
}

func (c *AdminProductSku) ListSKUPage(ctx context.Context, req *powerxtypes.ListSKUPageRequest) (*powerxtypes.ListSKUPageReply, error) {
	res := &powerxtypes.ListSKUPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/admin/product/skus/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductSku) GetSKU(ctx context.Context, req *powerxtypes.GetSKURequest) (*powerxtypes.GetSKUReply, error) {
	res := &powerxtypes.GetSKUReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/product/skus/%v", req.SKUId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductSku) CreateSKU(ctx context.Context, req *powerxtypes.CreateSKURequest) (*powerxtypes.CreateSKUReply, error) {
	res := &powerxtypes.CreateSKUReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/admin/product/skus").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductSku) ConfigSKU(ctx context.Context, req *powerxtypes.ConfigSKURequest) (*powerxtypes.ConfigSKUReply, error) {
	res := &powerxtypes.ConfigSKUReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/admin/product/skus/config").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductSku) PutSKU(ctx context.Context, req *powerxtypes.PutSKURequest) (*powerxtypes.PutSKUReply, error) {
	res := &powerxtypes.PutSKUReply{}
	err := c.H.Df().Method(http.MethodPut).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/product/skus/%v", req.SKUId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductSku) PatchSKU(ctx context.Context, req *powerxtypes.PatchSKURequest) (*powerxtypes.PatchSKUReply, error) {
	res := &powerxtypes.PatchSKUReply{}
	err := c.H.Df().Method(http.MethodPatch).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/product/skus/%v", req.SKUId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductSku) DeleteSKU(ctx context.Context, req *powerxtypes.DeleteSKURequest) (*powerxtypes.DeleteSKUReply, error) {
	res := &powerxtypes.DeleteSKUReply{}
	err := c.H.Df().Method(http.MethodDelete).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/product/skus/%v", req.SKUId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
