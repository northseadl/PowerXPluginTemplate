package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"fmt"
	"net/http"
)

type AdminProductPricebook struct {
	*PowerX
}

func (c *AdminProductPricebook) ListPriceBooks(req *powerxtypes.ListPriceBooksPageRequest) (*powerxtypes.ListPriceBooksPageReply, error) {
	res := &powerxtypes.ListPriceBooksPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/product/price-books/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductPricebook) GetPriceBook(req *powerxtypes.GetPriceBookRequest) (*powerxtypes.GetPriceBookReply, error) {
	res := &powerxtypes.GetPriceBookReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/product/price-books/%v", req.PriceBook)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductPricebook) UpsertPriceBook(req *powerxtypes.UpsertPriceBookRequest) (*powerxtypes.UpsertPriceBookReply, error) {
	res := &powerxtypes.UpsertPriceBookReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri("/api/v1/admin/product/price-books").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductPricebook) DeletePriceBook(req *powerxtypes.DeletePriceBookRequest) (*powerxtypes.DeletePriceBookReply, error) {
	res := &powerxtypes.DeletePriceBookReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/product/price-books/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
