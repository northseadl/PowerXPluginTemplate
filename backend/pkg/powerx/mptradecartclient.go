package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type MpTradeCart struct {
	*PowerX
}

func (c *MpTradeCart) ListCartItemsPage(req *types.ListCartItemsPageRequest) (*types.ListCartItemsPageReply, error) {
	res := &types.ListCartItemsPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/mp/trade/cart/items/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeCart) GetCart(req *types.GetCartRequest) (*types.GetCartReply, error) {
	res := &types.GetCartReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/mp/trade/cart/:cartId").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeCart) AddToCart(req *types.AddToCartRequest) (*types.AddToCartReply, error) {
	res := &types.AddToCartReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/mp/trade/cart/items").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeCart) UpdateCartItemQuantity(req *types.UpdateCartItemQuantityRequest) (*types.UpdateCartItemQuantityReply, error) {
	res := &types.UpdateCartItemQuantityReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/mp/trade/cart/items/%v", req.ItemId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeCart) RemoveCartItem(req *types.RemoveCartItemRequest) (*types.RemoveCartItemReply, error) {
	res := &types.RemoveCartItemReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/mp/trade/cart/items/%v", req.ItemId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeCart) ClearCartItems(req *types.ClearCartItemsRequest) (*types.ClearCartItemsReply, error) {
	res := &types.ClearCartItemsReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri("/api/v1/mp/trade/cart/items/clear").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
