package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type AdminProductCategory struct {
	*PowerX
}

func (c *AdminProductCategory) ListProductCategoryTree(req *types.ListProductCategoryTreeRequest) (*types.ListProductCategoryTreeReply, error) {
	res := &types.ListProductCategoryTreeReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/product/product-category-tree").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductCategory) GetProductCategory(req *types.GetProductCategoryRequest) (*types.GetProductCategoryReply, error) {
	res := &types.GetProductCategoryReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/product/product-categories/%s", req.ProductCategoryId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductCategory) CreateProductCategory(req *types.CreateProductCategoryRequest) (*types.CreateProductCategoryReply, error) {
	res := &types.CreateProductCategoryReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/product/product-categories").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductCategory) UpdateProductCategory(req *types.UpdateProductCategoryRequest) (*types.UpdateProductCategoryReply, error) {
	res := &types.UpdateProductCategoryReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/product/product-categories/%s", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductCategory) PatchProductCategory(req *types.PatchProductCategoryRequest) (*types.PatchProductCategoryReply, error) {
	res := &types.PatchProductCategoryReply{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri(fmt.Sprintf("/api/v1/admin/product/product-categories/%s", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductCategory) DeleteProductCategory(req *types.DeleteProductCategoryRequest) (*types.DeleteProductCategoryReply, error) {
	res := &types.DeleteProductCategoryReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/product/product-categories/%s", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
