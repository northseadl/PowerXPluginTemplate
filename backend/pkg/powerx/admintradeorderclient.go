package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"
	"fmt"
	"net/http"
)

type AdminTradeOrder struct {
	*PowerX
}

func (c *AdminTradeOrder) ListOrdersPage(ctx context.Context, req *powerxtypes.ListOrdersPageRequest) (*powerxtypes.ListOrdersPageReply, error) {
	res := &powerxtypes.ListOrdersPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/admin/trade/orders/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeOrder) CreateOrder(ctx context.Context, req *powerxtypes.CreateOrderRequest) (*powerxtypes.CreateOrderReply, error) {
	res := &powerxtypes.CreateOrderReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/admin/trade/orders").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeOrder) PutOrder(ctx context.Context, req *powerxtypes.PutOrderRequest) (*powerxtypes.PutOrderReply, error) {
	res := &powerxtypes.PutOrderReply{}
	err := c.H.Df().Method(http.MethodPut).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/trade/orders/%v", req.OrderId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeOrder) PatchOrder(ctx context.Context, req *powerxtypes.PatchOrderRequest) (*powerxtypes.PatchOrderReply, error) {
	res := &powerxtypes.PatchOrderReply{}
	err := c.H.Df().Method(http.MethodPatch).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/trade/orders/%v", req.OrderId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeOrder) DeleteOrder(ctx context.Context, req *powerxtypes.DeleteOrderRequest) (*powerxtypes.DeleteOrderReply, error) {
	res := &powerxtypes.DeleteOrderReply{}
	err := c.H.Df().Method(http.MethodDelete).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/trade/orders/%v", req.OrderId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeOrder) ExportOrders(ctx context.Context, req *powerxtypes.ExportOrdersRequest) (*powerxtypes.ExportOrdersReply, error) {
	res := &powerxtypes.ExportOrdersReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/admin/trade/orders/export").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
