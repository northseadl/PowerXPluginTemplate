package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type AdminTradeAddressShipping struct {
	*PowerX
}

func (c *AdminTradeAddressShipping) ListShippingAddressesPage(req *types.ListShippingAddressesPageRequest) (*types.ListShippingAddressesPageReply, error) {
	res := &types.ListShippingAddressesPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/trade/address/shipping/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressShipping) GetShippingAddress(req *types.GetShippingAddressRequest) (*types.GetShippingAddressReply, error) {
	res := &types.GetShippingAddressReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/trade/address/shipping/%v", req.ShippingAddressId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressShipping) CreateShippingAddress(req *types.CreateShippingAddressRequest) (*types.CreateShippingAddressReply, error) {
	res := &types.CreateShippingAddressReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/trade/address/shipping").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressShipping) PutShippingAddress(req *types.PutShippingAddressRequest) (*types.PutShippingAddressReply, error) {
	res := &types.PutShippingAddressReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/trade/address/shipping/%v", req.ShippingAddressId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressShipping) PatchShippingAddress(req *types.PatchShippingAddressRequest) (*types.PatchShippingAddressReply, error) {
	res := &types.PatchShippingAddressReply{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri(fmt.Sprintf("/api/v1/admin/trade/address/shipping/%v", req.ShippingAddressId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressShipping) DeleteShippingAddress(req *types.DeleteShippingAddressRequest) (*types.DeleteShippingAddressReply, error) {
	res := &types.DeleteShippingAddressReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/trade/address/shipping/%v", req.ShippingAddressId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
