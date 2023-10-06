package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"
	"fmt"
	"net/http"
)

type MpTradeOrder struct {
	*PowerX
}

func (c *MpTradeOrder) ListOrdersPage(ctx context.Context, req *powerxtypes.ListOrdersPageRequest) (*powerxtypes.ListOrdersPageReply, error) {
	res := &powerxtypes.ListOrdersPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/mp/trade/orders/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeOrder) CreateOrderByProducts(ctx context.Context, req *powerxtypes.CreateOrderByProductsRequest) (*powerxtypes.CreateOrderByProductsReply, error) {
	res := &powerxtypes.CreateOrderByProductsReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/mp/trade/orders/products").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeOrder) CreateOrderByCartItems(ctx context.Context, req *powerxtypes.CreateOrderByCartItemsRequest) (*powerxtypes.CreateOrderByCartItemsReply, error) {
	res := &powerxtypes.CreateOrderByCartItemsReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/mp/trade/orders/cart-items").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradeOrder) CancelOrder(ctx context.Context, req *powerxtypes.CancelOrderRequest) (*powerxtypes.CancelOrderReply, error) {
	res := &powerxtypes.CancelOrderReply{}
	err := c.H.Df().Method(http.MethodPut).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/mp/trade/orders/cancel/%v", req.OrderId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
