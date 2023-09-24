package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type AdminProductProductspecific struct {
	*PowerX
}

func (c *AdminProductProductspecific) ListProductSpecificPage(req *types.ListProductSpecificPageRequest) (*types.ListProductSpecificPageReply, error) {
	res := &types.ListProductSpecificPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/product/product-specifics/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductProductspecific) GetProductSpecific(req *types.GetProductSpecificRequest) (*types.GetProductSpecificReply, error) {
	res := &types.GetProductSpecificReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/product/product-specifics/%v", req.ProductSpecificId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductProductspecific) CreateProductSpecific(req *types.CreateProductSpecificRequest) (*types.CreateProductSpecificReply, error) {
	res := &types.CreateProductSpecificReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/product/product-specifics").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductProductspecific) ConfigProductSpecific(req *types.ConfigProductSpecificRequest) (*types.ConfigProductSpecificReply, error) {
	res := &types.ConfigProductSpecificReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/product/product-specifics/config").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductProductspecific) PutProductSpecific(req *types.PutProductSpecificRequest) (*types.PutProductSpecificReply, error) {
	res := &types.PutProductSpecificReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/product/product-specifics/%v", req.ProductSpecificId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductProductspecific) PatchProductSpecific(req *types.PatchProductSpecificRequest) (*types.PatchProductSpecificReply, error) {
	res := &types.PatchProductSpecificReply{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri(fmt.Sprintf("/api/v1/admin/product/product-specifics/%v", req.ProductSpecificId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductProductspecific) DeleteProductSpecific(req *types.DeleteProductSpecificRequest) (*types.DeleteProductSpecificReply, error) {
	res := &types.DeleteProductSpecificReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/product/product-specifics/%v", req.ProductSpecificId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
