package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type AdminProductPricebook struct {
	*PowerX
}

func (c *AdminProductPricebook) ListPriceBooks(req *types.ListPriceBooksPageRequest) (*types.ListPriceBooksPageReply, error) {
	res := &types.ListPriceBooksPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/product/price-books/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductPricebook) GetPriceBook(req *types.GetPriceBookRequest) (*types.GetPriceBookReply, error) {
	res := &types.GetPriceBookReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/product/price-books/%v", req.PriceBook)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductPricebook) UpsertPriceBook(req *types.UpsertPriceBookRequest) (*types.UpsertPriceBookReply, error) {
	res := &types.UpsertPriceBookReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri("/api/v1/admin/product/price-books").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductPricebook) DeletePriceBook(req *types.DeletePriceBookRequest) (*types.DeletePriceBookReply, error) {
	res := &types.DeletePriceBookReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/product/price-books/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
