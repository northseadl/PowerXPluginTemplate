package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"
	"fmt"
	"net/http"
)

type MpTradeCart struct {
	*PowerX
}

func (c *MpTradeCart) ListCartItemsPage(ctx context.Context, req *powerxtypes.ListCartItemsPageRequest) (*powerxtypes.ListCartItemsPageReply, error) {
	res := &powerxtypes.ListCartItemsPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/mp/trade/cart/items/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeCart) GetCart(ctx context.Context, req *powerxtypes.GetCartRequest) (*powerxtypes.GetCartReply, error) {
	res := &powerxtypes.GetCartReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/mp/trade/cart/:cartId").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeCart) AddToCart(ctx context.Context, req *powerxtypes.AddToCartRequest) (*powerxtypes.AddToCartReply, error) {
	res := &powerxtypes.AddToCartReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/mp/trade/cart/items").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeCart) UpdateCartItemQuantity(ctx context.Context, req *powerxtypes.UpdateCartItemQuantityRequest) (*powerxtypes.UpdateCartItemQuantityReply, error) {
	res := &powerxtypes.UpdateCartItemQuantityReply{}
	err := c.H.Df().Method(http.MethodPut).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/mp/trade/cart/items/%v", req.ItemId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeCart) RemoveCartItem(ctx context.Context, req *powerxtypes.RemoveCartItemRequest) (*powerxtypes.RemoveCartItemReply, error) {
	res := &powerxtypes.RemoveCartItemReply{}
	err := c.H.Df().Method(http.MethodDelete).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/mp/trade/cart/items/%v", req.ItemId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeCart) ClearCartItems(ctx context.Context, req *powerxtypes.ClearCartItemsRequest) (*powerxtypes.ClearCartItemsReply, error) {
	res := &powerxtypes.ClearCartItemsReply{}
	err := c.H.Df().Method(http.MethodDelete).
		WithContext(ctx).
		Uri("/api/v1/mp/trade/cart/items/clear").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
