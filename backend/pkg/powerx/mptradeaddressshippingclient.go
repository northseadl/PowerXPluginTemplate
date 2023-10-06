package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"
	"fmt"
	"net/http"
)

type MpTradeAddressShipping struct {
	*PowerX
}

func (c *MpTradeAddressShipping) ListShippingAddressesPage(ctx context.Context, req *powerxtypes.ListShippingAddressesPageRequest) (*powerxtypes.ListShippingAddressesPageReply, error) {
	res := &powerxtypes.ListShippingAddressesPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/mp/trade/address/shipping/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeAddressShipping) GetShippingAddress(ctx context.Context, req *powerxtypes.GetShippingAddressRequest) (*powerxtypes.GetShippingAddressReply, error) {
	res := &powerxtypes.GetShippingAddressReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/mp/trade/address/shipping/%v", req.ShippingAddressId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeAddressShipping) CreateShippingAddress(ctx context.Context, req *powerxtypes.CreateShippingAddressRequest) (*powerxtypes.CreateShippingAddressReply, error) {
	res := &powerxtypes.CreateShippingAddressReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/mp/trade/address/shipping").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeAddressShipping) PutShippingAddress(ctx context.Context, req *powerxtypes.PutShippingAddressRequest) (*powerxtypes.PutShippingAddressReply, error) {
	res := &powerxtypes.PutShippingAddressReply{}
	err := c.H.Df().Method(http.MethodPut).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/mp/trade/address/shipping/%v", req.ShippingAddressId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeAddressShipping) PatchShippingAddress(ctx context.Context, req *powerxtypes.PatchShippingAddressRequest) (*powerxtypes.PatchShippingAddressReply, error) {
	res := &powerxtypes.PatchShippingAddressReply{}
	err := c.H.Df().Method(http.MethodPatch).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/mp/trade/address/shipping/%v", req.ShippingAddressId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeAddressShipping) DeleteShippingAddress(ctx context.Context, req *powerxtypes.DeleteShippingAddressRequest) (*powerxtypes.DeleteShippingAddressReply, error) {
	res := &powerxtypes.DeleteShippingAddressReply{}
	err := c.H.Df().Method(http.MethodDelete).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/mp/trade/address/shipping/%v", req.ShippingAddressId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
