package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type AdminProductPricebookentry struct {
	*PowerX
}

func (c *AdminProductPricebookentry) ListPriceBookEntries(req *types.ListPriceBookEntriesPageRequest) (*types.ListPriceBookEntriesPageReply, error) {
	res := &types.ListPriceBookEntriesPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/product/price-book-entries/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductPricebookentry) GetPriceBookEntry(req *types.GetPriceBookEntryRequest) (*types.GetPriceBookEntryReply, error) {
	res := &types.GetPriceBookEntryReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/product/price-book-entries/%v", req.PriceBookEntry)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductPricebookentry) ConfigPriceBookEntry(req *types.ConfigPriceBookEntryRequest) (*types.ConfigPriceBookEntryReply, error) {
	res := &types.ConfigPriceBookEntryReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/product/price-book-entries/config").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductPricebookentry) UpdatePriceBookEntry(req *types.UpdatePriceBookEntryRequest) (*types.UpdatePriceBookEntryReply, error) {
	res := &types.UpdatePriceBookEntryReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri(fmt.Sprintf("/api/v1/admin/product/price-book-entries/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductPricebookentry) DeletePriceBookEntry(req *types.DeletePriceBookEntryRequest) (*types.DeletePriceBookEntryReply, error) {
	res := &types.DeletePriceBookEntryReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/product/price-book-entries/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
