package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"
	"fmt"
	"net/http"
)

type AdminProductPricebook struct {
	*PowerX
}

func (c *AdminProductPricebook) ListPriceBooks(ctx context.Context, req *powerxtypes.ListPriceBooksPageRequest) (*powerxtypes.ListPriceBooksPageReply, error) {
	res := &powerxtypes.ListPriceBooksPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/admin/product/price-books/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductPricebook) GetPriceBook(ctx context.Context, req *powerxtypes.GetPriceBookRequest) (*powerxtypes.GetPriceBookReply, error) {
	res := &powerxtypes.GetPriceBookReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/product/price-books/%v", req.PriceBook)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductPricebook) UpsertPriceBook(ctx context.Context, req *powerxtypes.UpsertPriceBookRequest) (*powerxtypes.UpsertPriceBookReply, error) {
	res := &powerxtypes.UpsertPriceBookReply{}
	err := c.H.Df().Method(http.MethodPut).
		WithContext(ctx).
		Uri("/api/v1/admin/product/price-books").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductPricebook) DeletePriceBook(ctx context.Context, req *powerxtypes.DeletePriceBookRequest) (*powerxtypes.DeletePriceBookReply, error) {
	res := &powerxtypes.DeletePriceBookReply{}
	err := c.H.Df().Method(http.MethodDelete).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/product/price-books/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
