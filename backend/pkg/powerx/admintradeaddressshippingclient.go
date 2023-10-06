package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"fmt"
	"net/http"
)

type AdminTradeAddressShipping struct {
	*PowerX
}

func (c *AdminTradeAddressShipping) ListShippingAddressesPage(req *powerxtypes.ListShippingAddressesPageRequest) (*powerxtypes.ListShippingAddressesPageReply, error) {
	res := &powerxtypes.ListShippingAddressesPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/trade/address/shipping/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressShipping) GetShippingAddress(req *powerxtypes.GetShippingAddressRequest) (*powerxtypes.GetShippingAddressReply, error) {
	res := &powerxtypes.GetShippingAddressReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/trade/address/shipping/%v", req.ShippingAddressId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressShipping) CreateShippingAddress(req *powerxtypes.CreateShippingAddressRequest) (*powerxtypes.CreateShippingAddressReply, error) {
	res := &powerxtypes.CreateShippingAddressReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/trade/address/shipping").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressShipping) PutShippingAddress(req *powerxtypes.PutShippingAddressRequest) (*powerxtypes.PutShippingAddressReply, error) {
	res := &powerxtypes.PutShippingAddressReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/trade/address/shipping/%v", req.ShippingAddressId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressShipping) PatchShippingAddress(req *powerxtypes.PatchShippingAddressRequest) (*powerxtypes.PatchShippingAddressReply, error) {
	res := &powerxtypes.PatchShippingAddressReply{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri(fmt.Sprintf("/api/v1/admin/trade/address/shipping/%v", req.ShippingAddressId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressShipping) DeleteShippingAddress(req *powerxtypes.DeleteShippingAddressRequest) (*powerxtypes.DeleteShippingAddressReply, error) {
	res := &powerxtypes.DeleteShippingAddressReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/trade/address/shipping/%v", req.ShippingAddressId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
