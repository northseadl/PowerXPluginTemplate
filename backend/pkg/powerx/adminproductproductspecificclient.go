package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"fmt"
	"net/http"
)

type AdminProductProductspecific struct {
	*PowerX
}

func (c *AdminProductProductspecific) ListProductSpecificPage(req *powerxtypes.ListProductSpecificPageRequest) (*powerxtypes.ListProductSpecificPageReply, error) {
	res := &powerxtypes.ListProductSpecificPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/product/product-specifics/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductProductspecific) GetProductSpecific(req *powerxtypes.GetProductSpecificRequest) (*powerxtypes.GetProductSpecificReply, error) {
	res := &powerxtypes.GetProductSpecificReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/product/product-specifics/%v", req.ProductSpecificId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductProductspecific) CreateProductSpecific(req *powerxtypes.CreateProductSpecificRequest) (*powerxtypes.CreateProductSpecificReply, error) {
	res := &powerxtypes.CreateProductSpecificReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/product/product-specifics").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductProductspecific) ConfigProductSpecific(req *powerxtypes.ConfigProductSpecificRequest) (*powerxtypes.ConfigProductSpecificReply, error) {
	res := &powerxtypes.ConfigProductSpecificReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/product/product-specifics/config").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductProductspecific) PutProductSpecific(req *powerxtypes.PutProductSpecificRequest) (*powerxtypes.PutProductSpecificReply, error) {
	res := &powerxtypes.PutProductSpecificReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/product/product-specifics/%v", req.ProductSpecificId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductProductspecific) PatchProductSpecific(req *powerxtypes.PatchProductSpecificRequest) (*powerxtypes.PatchProductSpecificReply, error) {
	res := &powerxtypes.PatchProductSpecificReply{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri(fmt.Sprintf("/api/v1/admin/product/product-specifics/%v", req.ProductSpecificId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductProductspecific) DeleteProductSpecific(req *powerxtypes.DeleteProductSpecificRequest) (*powerxtypes.DeleteProductSpecificReply, error) {
	res := &powerxtypes.DeleteProductSpecificReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/product/product-specifics/%v", req.ProductSpecificId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
