package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"fmt"
	"net/http"
)

type AdminProduct struct {
	*PowerX
}

func (c *AdminProduct) ListProductsPage(req *powerxtypes.ListProductsPageRequest) (*powerxtypes.ListProductsPageReply, error) {
	res := &powerxtypes.ListProductsPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/product/products/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProduct) GetProduct(req *powerxtypes.GetProductRequest) (*powerxtypes.GetProductReply, error) {
	res := &powerxtypes.GetProductReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/product/products/%v", req.ProductId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProduct) CreateProduct(req *powerxtypes.CreateProductRequest) (*powerxtypes.CreateProductReply, error) {
	res := &powerxtypes.CreateProductReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/product/products").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProduct) PutProduct(req *powerxtypes.PutProductRequest) (*powerxtypes.PutProductReply, error) {
	res := &powerxtypes.PutProductReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/product/products/%v", req.ProductId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProduct) PatchProduct(req *powerxtypes.PatchProductRequest) (*powerxtypes.PatchProductReply, error) {
	res := &powerxtypes.PatchProductReply{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri(fmt.Sprintf("/api/v1/admin/product/products/%v", req.ProductId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProduct) DeleteProduct(req *powerxtypes.DeleteProductRequest) (*powerxtypes.DeleteProductReply, error) {
	res := &powerxtypes.DeleteProductReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/product/products/%v", req.ProductId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProduct) AssignProductToProductCategory(req *powerxtypes.AssignProductToProductCategoryRequest) (*powerxtypes.AssignProductToProductCategoryReply, error) {
	res := &powerxtypes.AssignProductToProductCategoryReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/product/products/:id/actions/assign-to-product-categroy").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
