package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"fmt"
	"net/http"
)

type AdminTradeAddressDelivery struct {
	*PowerX
}

func (c *AdminTradeAddressDelivery) ListDeliveryAddressesPage(req *powerxtypes.ListDeliveryAddressesPageRequest) (*powerxtypes.ListDeliveryAddressesPageReply, error) {
	res := &powerxtypes.ListDeliveryAddressesPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/trade/address/delivery/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressDelivery) GetDeliveryAddress(req *powerxtypes.GetDeliveryAddressRequest) (*powerxtypes.GetDeliveryAddressReply, error) {
	res := &powerxtypes.GetDeliveryAddressReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/trade/address/delivery/%v", req.DeliveryAddressId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressDelivery) CreateDeliveryAddress(req *powerxtypes.CreateDeliveryAddressRequest) (*powerxtypes.CreateDeliveryAddressReply, error) {
	res := &powerxtypes.CreateDeliveryAddressReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/trade/address/delivery").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressDelivery) PutDeliveryAddress(req *powerxtypes.PutDeliveryAddressRequest) (*powerxtypes.PutDeliveryAddressReply, error) {
	res := &powerxtypes.PutDeliveryAddressReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/trade/address/delivery/%v", req.DeliveryAddressId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressDelivery) PatchDeliveryAddress(req *powerxtypes.PatchDeliveryAddressRequest) (*powerxtypes.PatchDeliveryAddressReply, error) {
	res := &powerxtypes.PatchDeliveryAddressReply{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri(fmt.Sprintf("/api/v1/admin/trade/address/delivery/%v", req.DeliveryAddressId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressDelivery) DeleteDeliveryAddress(req *powerxtypes.DeleteDeliveryAddressRequest) (*powerxtypes.DeleteDeliveryAddressReply, error) {
	res := &powerxtypes.DeleteDeliveryAddressReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/trade/address/delivery/%v", req.DeliveryAddressId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
