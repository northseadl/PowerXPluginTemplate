package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type AdminProduct struct {
	*PowerX
}

func (c *AdminProduct) ListProductsPage(req *types.ListProductsPageRequest) (*types.ListProductsPageReply, error) {
	res := &types.ListProductsPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/product/products/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProduct) GetProduct(req *types.GetProductRequest) (*types.GetProductReply, error) {
	res := &types.GetProductReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/product/products/%s", req.ProductId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProduct) CreateProduct(req *types.CreateProductRequest) (*types.CreateProductReply, error) {
	res := &types.CreateProductReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/product/products").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProduct) PutProduct(req *types.PutProductRequest) (*types.PutProductReply, error) {
	res := &types.PutProductReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/product/products/%s", req.ProductId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProduct) PatchProduct(req *types.PatchProductRequest) (*types.PatchProductReply, error) {
	res := &types.PatchProductReply{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri(fmt.Sprintf("/api/v1/admin/product/products/%s", req.ProductId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProduct) DeleteProduct(req *types.DeleteProductRequest) (*types.DeleteProductReply, error) {
	res := &types.DeleteProductReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/product/products/%s", req.ProductId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProduct) AssignProductToProductCategory(req *types.AssignProductToProductCategoryRequest) (*types.AssignProductToProductCategoryReply, error) {
	res := &types.AssignProductToProductCategoryReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/product/products/:id/actions/assign-to-product-categroy").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
