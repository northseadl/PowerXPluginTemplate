package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"fmt"
	"net/http"
)

type MpTradeAddressBilling struct {
	*PowerX
}

func (c *MpTradeAddressBilling) ListBillingAddressesPage(req *powerxtypes.ListBillingAddressesPageRequest) (*powerxtypes.ListBillingAddressesPageReply, error) {
	res := &powerxtypes.ListBillingAddressesPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/mp/trade/address/billing/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeAddressBilling) GetBillingAddress(req *powerxtypes.GetBillingAddressRequest) (*powerxtypes.GetBillingAddressReply, error) {
	res := &powerxtypes.GetBillingAddressReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/mp/trade/address/billing/%v", req.BillingAddressId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeAddressBilling) CreateBillingAddress(req *powerxtypes.CreateBillingAddressRequest) (*powerxtypes.CreateBillingAddressReply, error) {
	res := &powerxtypes.CreateBillingAddressReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/mp/trade/address/billing").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeAddressBilling) PutBillingAddress(req *powerxtypes.PutBillingAddressRequest) (*powerxtypes.PutBillingAddressReply, error) {
	res := &powerxtypes.PutBillingAddressReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/mp/trade/address/billing/%v", req.BillingAddressId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeAddressBilling) PatchBillingAddress(req *powerxtypes.PatchBillingAddressRequest) (*powerxtypes.PatchBillingAddressReply, error) {
	res := &powerxtypes.PatchBillingAddressReply{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri(fmt.Sprintf("/api/v1/mp/trade/address/billing/%v", req.BillingAddressId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeAddressBilling) DeleteBillingAddress(req *powerxtypes.DeleteBillingAddressRequest) (*powerxtypes.DeleteBillingAddressReply, error) {
	res := &powerxtypes.DeleteBillingAddressReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/mp/trade/address/billing/%v", req.BillingAddressId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
