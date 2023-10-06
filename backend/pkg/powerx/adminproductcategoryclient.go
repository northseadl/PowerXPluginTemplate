package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"fmt"
	"net/http"
)

type AdminProductCategory struct {
	*PowerX
}

func (c *AdminProductCategory) ListProductCategoryTree(req *powerxtypes.ListProductCategoryTreeRequest) (*powerxtypes.ListProductCategoryTreeReply, error) {
	res := &powerxtypes.ListProductCategoryTreeReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/product/product-category-tree").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductCategory) GetProductCategory(req *powerxtypes.GetProductCategoryRequest) (*powerxtypes.GetProductCategoryReply, error) {
	res := &powerxtypes.GetProductCategoryReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/product/product-categories/%v", req.ProductCategoryId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductCategory) CreateProductCategory(req *powerxtypes.CreateProductCategoryRequest) (*powerxtypes.CreateProductCategoryReply, error) {
	res := &powerxtypes.CreateProductCategoryReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/product/product-categories").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductCategory) UpdateProductCategory(req *powerxtypes.UpdateProductCategoryRequest) (*powerxtypes.UpdateProductCategoryReply, error) {
	res := &powerxtypes.UpdateProductCategoryReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/product/product-categories/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductCategory) PatchProductCategory(req *powerxtypes.PatchProductCategoryRequest) (*powerxtypes.PatchProductCategoryReply, error) {
	res := &powerxtypes.PatchProductCategoryReply{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri(fmt.Sprintf("/api/v1/admin/product/product-categories/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductCategory) DeleteProductCategory(req *powerxtypes.DeleteProductCategoryRequest) (*powerxtypes.DeleteProductCategoryReply, error) {
	res := &powerxtypes.DeleteProductCategoryReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/product/product-categories/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
