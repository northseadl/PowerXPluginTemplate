package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"fmt"
	"net/http"
)

type AdminProductPricebookentry struct {
	*PowerX
}

func (c *AdminProductPricebookentry) ListPriceBookEntries(req *powerxtypes.ListPriceBookEntriesPageRequest) (*powerxtypes.ListPriceBookEntriesPageReply, error) {
	res := &powerxtypes.ListPriceBookEntriesPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/product/price-book-entries/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductPricebookentry) GetPriceBookEntry(req *powerxtypes.GetPriceBookEntryRequest) (*powerxtypes.GetPriceBookEntryReply, error) {
	res := &powerxtypes.GetPriceBookEntryReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/product/price-book-entries/%v", req.PriceBookEntry)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductPricebookentry) ConfigPriceBookEntry(req *powerxtypes.ConfigPriceBookEntryRequest) (*powerxtypes.ConfigPriceBookEntryReply, error) {
	res := &powerxtypes.ConfigPriceBookEntryReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/product/price-book-entries/config").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductPricebookentry) UpdatePriceBookEntry(req *powerxtypes.UpdatePriceBookEntryRequest) (*powerxtypes.UpdatePriceBookEntryReply, error) {
	res := &powerxtypes.UpdatePriceBookEntryReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri(fmt.Sprintf("/api/v1/admin/product/price-book-entries/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminProductPricebookentry) DeletePriceBookEntry(req *powerxtypes.DeletePriceBookEntryRequest) (*powerxtypes.DeletePriceBookEntryReply, error) {
	res := &powerxtypes.DeletePriceBookEntryReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/product/price-book-entries/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
