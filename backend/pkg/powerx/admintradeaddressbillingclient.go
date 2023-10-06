package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"fmt"
	"net/http"
)

type AdminTradeAddressBilling struct {
	*PowerX
}

func (c *AdminTradeAddressBilling) ListBillingAddressesPage(req *powerxtypes.ListBillingAddressesPageRequest) (*powerxtypes.ListBillingAddressesPageReply, error) {
	res := &powerxtypes.ListBillingAddressesPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/trade/address/billing/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressBilling) GetBillingAddress(req *powerxtypes.GetBillingAddressRequest) (*powerxtypes.GetBillingAddressReply, error) {
	res := &powerxtypes.GetBillingAddressReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/trade/address/billing/%v", req.BillingAddressId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressBilling) CreateBillingAddress(req *powerxtypes.CreateBillingAddressRequest) (*powerxtypes.CreateBillingAddressReply, error) {
	res := &powerxtypes.CreateBillingAddressReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/trade/address/billing").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressBilling) PutBillingAddress(req *powerxtypes.PutBillingAddressRequest) (*powerxtypes.PutBillingAddressReply, error) {
	res := &powerxtypes.PutBillingAddressReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/trade/address/billing/%v", req.BillingAddressId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressBilling) PatchBillingAddress(req *powerxtypes.PatchBillingAddressRequest) (*powerxtypes.PatchBillingAddressReply, error) {
	res := &powerxtypes.PatchBillingAddressReply{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri(fmt.Sprintf("/api/v1/admin/trade/address/billing/%v", req.BillingAddressId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressBilling) DeleteBillingAddress(req *powerxtypes.DeleteBillingAddressRequest) (*powerxtypes.DeleteBillingAddressReply, error) {
	res := &powerxtypes.DeleteBillingAddressReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/trade/address/billing/%v", req.BillingAddressId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
