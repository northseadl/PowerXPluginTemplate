package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"
	"fmt"
	"net/http"
)

type AdminProductCategory struct {
	*PowerX
}

func (c *AdminProductCategory) ListProductCategoryTree(ctx context.Context, req *powerxtypes.ListProductCategoryTreeRequest) (*powerxtypes.ListProductCategoryTreeReply, error) {
	res := &powerxtypes.ListProductCategoryTreeReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/admin/product/product-category-tree").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductCategory) GetProductCategory(ctx context.Context, req *powerxtypes.GetProductCategoryRequest) (*powerxtypes.GetProductCategoryReply, error) {
	res := &powerxtypes.GetProductCategoryReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/product/product-categories/%v", req.ProductCategoryId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductCategory) CreateProductCategory(ctx context.Context, req *powerxtypes.CreateProductCategoryRequest) (*powerxtypes.CreateProductCategoryReply, error) {
	res := &powerxtypes.CreateProductCategoryReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/admin/product/product-categories").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductCategory) UpdateProductCategory(ctx context.Context, req *powerxtypes.UpdateProductCategoryRequest) (*powerxtypes.UpdateProductCategoryReply, error) {
	res := &powerxtypes.UpdateProductCategoryReply{}
	err := c.H.Df().Method(http.MethodPut).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/product/product-categories/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductCategory) PatchProductCategory(ctx context.Context, req *powerxtypes.PatchProductCategoryRequest) (*powerxtypes.PatchProductCategoryReply, error) {
	res := &powerxtypes.PatchProductCategoryReply{}
	err := c.H.Df().Method(http.MethodPatch).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/product/product-categories/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductCategory) DeleteProductCategory(ctx context.Context, req *powerxtypes.DeleteProductCategoryRequest) (*powerxtypes.DeleteProductCategoryReply, error) {
	res := &powerxtypes.DeleteProductCategoryReply{}
	err := c.H.Df().Method(http.MethodDelete).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/product/product-categories/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
