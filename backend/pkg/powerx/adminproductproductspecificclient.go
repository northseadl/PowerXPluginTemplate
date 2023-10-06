package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"
	"fmt"
	"net/http"
)

type AdminProductProductspecific struct {
	*PowerX
}

func (c *AdminProductProductspecific) ListProductSpecificPage(ctx context.Context, req *powerxtypes.ListProductSpecificPageRequest) (*powerxtypes.ListProductSpecificPageReply, error) {
	res := &powerxtypes.ListProductSpecificPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/admin/product/product-specifics/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductProductspecific) GetProductSpecific(ctx context.Context, req *powerxtypes.GetProductSpecificRequest) (*powerxtypes.GetProductSpecificReply, error) {
	res := &powerxtypes.GetProductSpecificReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/product/product-specifics/%v", req.ProductSpecificId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductProductspecific) CreateProductSpecific(ctx context.Context, req *powerxtypes.CreateProductSpecificRequest) (*powerxtypes.CreateProductSpecificReply, error) {
	res := &powerxtypes.CreateProductSpecificReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/admin/product/product-specifics").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductProductspecific) ConfigProductSpecific(ctx context.Context, req *powerxtypes.ConfigProductSpecificRequest) (*powerxtypes.ConfigProductSpecificReply, error) {
	res := &powerxtypes.ConfigProductSpecificReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/admin/product/product-specifics/config").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductProductspecific) PutProductSpecific(ctx context.Context, req *powerxtypes.PutProductSpecificRequest) (*powerxtypes.PutProductSpecificReply, error) {
	res := &powerxtypes.PutProductSpecificReply{}
	err := c.H.Df().Method(http.MethodPut).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/product/product-specifics/%v", req.ProductSpecificId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductProductspecific) PatchProductSpecific(ctx context.Context, req *powerxtypes.PatchProductSpecificRequest) (*powerxtypes.PatchProductSpecificReply, error) {
	res := &powerxtypes.PatchProductSpecificReply{}
	err := c.H.Df().Method(http.MethodPatch).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/product/product-specifics/%v", req.ProductSpecificId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductProductspecific) DeleteProductSpecific(ctx context.Context, req *powerxtypes.DeleteProductSpecificRequest) (*powerxtypes.DeleteProductSpecificReply, error) {
	res := &powerxtypes.DeleteProductSpecificReply{}
	err := c.H.Df().Method(http.MethodDelete).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/product/product-specifics/%v", req.ProductSpecificId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
