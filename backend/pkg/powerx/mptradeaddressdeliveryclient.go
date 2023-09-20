package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type MpTradeAddressDelivery struct {
	*PowerX
}

func (c *MpTradeAddressDelivery) ListDeliveryAddressesPage(req *types.ListDeliveryAddressesPageRequest) (*types.ListDeliveryAddressesPageReply, error) {
	res := &types.ListDeliveryAddressesPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/mp/trade/address/delivery/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeAddressDelivery) GetDeliveryAddress(req *types.GetDeliveryAddressRequest) (*types.GetDeliveryAddressReply, error) {
	res := &types.GetDeliveryAddressReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/mp/trade/address/delivery/%s", req.DeliveryAddressId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeAddressDelivery) CreateDeliveryAddress(req *types.CreateDeliveryAddressRequest) (*types.CreateDeliveryAddressReply, error) {
	res := &types.CreateDeliveryAddressReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/mp/trade/address/delivery").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeAddressDelivery) PutDeliveryAddress(req *types.PutDeliveryAddressRequest) (*types.PutDeliveryAddressReply, error) {
	res := &types.PutDeliveryAddressReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/mp/trade/address/delivery/%s", req.DeliveryAddressId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeAddressDelivery) PatchDeliveryAddress(req *types.PatchDeliveryAddressRequest) (*types.PatchDeliveryAddressReply, error) {
	res := &types.PatchDeliveryAddressReply{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri(fmt.Sprintf("/api/v1/mp/trade/address/delivery/%s", req.DeliveryAddressId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeAddressDelivery) DeleteDeliveryAddress(req *types.DeleteDeliveryAddressRequest) (*types.DeleteDeliveryAddressReply, error) {
	res := &types.DeleteDeliveryAddressReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/mp/trade/address/delivery/%s", req.DeliveryAddressId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
