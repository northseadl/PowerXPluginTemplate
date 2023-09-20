package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type MpTradeOrder struct {
	*PowerX
}

func (c *MpTradeOrder) ListOrdersPage(req *types.ListOrdersPageRequest) (*types.ListOrdersPageReply, error) {
	res := &types.ListOrdersPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/mp/trade/orders/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeOrder) GetOrder(req *types.GetOrderRequest) (*types.GetOrderReply, error) {
	res := &types.GetOrderReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/mp/trade/orders/%s", req.orderId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeOrder) CreateOrderByProducts(req *types.CreateOrderByProductsRequest) (*types.CreateOrderByProductsReply, error) {
	res := &types.CreateOrderByProductsReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/mp/trade/orders/products").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeOrder) CreateOrderByCartItems(req *types.CreateOrderByCartItemsRequest) (*types.CreateOrderByCartItemsReply, error) {
	res := &types.CreateOrderByCartItemsReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/mp/trade/orders/cart-items").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeOrder) CancelOrder(req *types.CancelOrderRequest) (*types.CancelOrderReply, error) {
	res := &types.CancelOrderReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/mp/trade/orders/cancel/%s", req.OrderId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
