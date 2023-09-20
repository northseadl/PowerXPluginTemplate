package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type AdminTradeOrder struct {
	*PowerX
}

func (c *AdminTradeOrder) ListOrdersPage(req *types.ListOrdersPageRequest) (*types.ListOrdersPageReply, error) {
	res := &types.ListOrdersPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/trade/orders/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeOrder) GetOrder(req *types.GetOrderRequest) (*types.GetOrderReply, error) {
	res := &types.GetOrderReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/trade/orders/%s", req.orderId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeOrder) CreateOrder(req *types.CreateOrderRequest) (*types.CreateOrderReply, error) {
	res := &types.CreateOrderReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/trade/orders").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeOrder) PutOrder(req *types.PutOrderRequest) (*types.PutOrderReply, error) {
	res := &types.PutOrderReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/trade/orders/%s", req.OrderId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeOrder) PatchOrder(req *types.PatchOrderRequest) (*types.PatchOrderReply, error) {
	res := &types.PatchOrderReply{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri(fmt.Sprintf("/api/v1/admin/trade/orders/%s", req.OrderId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeOrder) DeleteOrder(req *types.DeleteOrderRequest) (*types.DeleteOrderReply, error) {
	res := &types.DeleteOrderReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/trade/orders/%s", req.OrderId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeOrder) ExportOrders(req *types.ExportOrdersRequest) (*types.ExportOrdersReply, error) {
	res := &types.ExportOrdersReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/trade/orders/export").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
