package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"
	"fmt"
	"net/http"
)

type AdminTradeAddressDelivery struct {
	*PowerX
}

func (c *AdminTradeAddressDelivery) ListDeliveryAddressesPage(ctx context.Context, req *powerxtypes.ListDeliveryAddressesPageRequest) (*powerxtypes.ListDeliveryAddressesPageReply, error) {
	res := &powerxtypes.ListDeliveryAddressesPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/admin/trade/address/delivery/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressDelivery) GetDeliveryAddress(ctx context.Context, req *powerxtypes.GetDeliveryAddressRequest) (*powerxtypes.GetDeliveryAddressReply, error) {
	res := &powerxtypes.GetDeliveryAddressReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/trade/address/delivery/%v", req.DeliveryAddressId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressDelivery) CreateDeliveryAddress(ctx context.Context, req *powerxtypes.CreateDeliveryAddressRequest) (*powerxtypes.CreateDeliveryAddressReply, error) {
	res := &powerxtypes.CreateDeliveryAddressReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/admin/trade/address/delivery").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressDelivery) PutDeliveryAddress(ctx context.Context, req *powerxtypes.PutDeliveryAddressRequest) (*powerxtypes.PutDeliveryAddressReply, error) {
	res := &powerxtypes.PutDeliveryAddressReply{}
	err := c.H.Df().Method(http.MethodPut).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/trade/address/delivery/%v", req.DeliveryAddressId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressDelivery) PatchDeliveryAddress(ctx context.Context, req *powerxtypes.PatchDeliveryAddressRequest) (*powerxtypes.PatchDeliveryAddressReply, error) {
	res := &powerxtypes.PatchDeliveryAddressReply{}
	err := c.H.Df().Method(http.MethodPatch).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/trade/address/delivery/%v", req.DeliveryAddressId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressDelivery) DeleteDeliveryAddress(ctx context.Context, req *powerxtypes.DeleteDeliveryAddressRequest) (*powerxtypes.DeleteDeliveryAddressReply, error) {
	res := &powerxtypes.DeleteDeliveryAddressReply{}
	err := c.H.Df().Method(http.MethodDelete).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/trade/address/delivery/%v", req.DeliveryAddressId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
