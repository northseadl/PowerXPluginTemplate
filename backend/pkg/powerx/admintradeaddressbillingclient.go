package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type AdminTradeAddressBilling struct {
	*PowerX
}

func (c *AdminTradeAddressBilling) ListBillingAddressesPage(req *types.ListBillingAddressesPageRequest) (*types.ListBillingAddressesPageReply, error) {
	res := &types.ListBillingAddressesPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/trade/address/billing/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressBilling) GetBillingAddress(req *types.GetBillingAddressRequest) (*types.GetBillingAddressReply, error) {
	res := &types.GetBillingAddressReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/trade/address/billing/%s", req.BillingAddressId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressBilling) CreateBillingAddress(req *types.CreateBillingAddressRequest) (*types.CreateBillingAddressReply, error) {
	res := &types.CreateBillingAddressReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/trade/address/billing").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressBilling) PutBillingAddress(req *types.PutBillingAddressRequest) (*types.PutBillingAddressReply, error) {
	res := &types.PutBillingAddressReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/trade/address/billing/%s", req.BillingAddressId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressBilling) PatchBillingAddress(req *types.PatchBillingAddressRequest) (*types.PatchBillingAddressReply, error) {
	res := &types.PatchBillingAddressReply{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri(fmt.Sprintf("/api/v1/admin/trade/address/billing/%s", req.BillingAddressId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressBilling) DeleteBillingAddress(req *types.DeleteBillingAddressRequest) (*types.DeleteBillingAddressReply, error) {
	res := &types.DeleteBillingAddressReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/trade/address/billing/%s", req.BillingAddressId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
