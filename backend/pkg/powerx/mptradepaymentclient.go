package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"
	"fmt"
	"net/http"
)

type MpTradePayment struct {
	*PowerX
}

func (c *MpTradePayment) ListPaymentsPage(ctx context.Context, req *powerxtypes.ListPaymentsPageRequest) (*powerxtypes.ListPaymentsPageReply, error) {
	res := &powerxtypes.ListPaymentsPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/mp/trade/payments/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradePayment) GetPayment(ctx context.Context, req *powerxtypes.GetPaymentRequest) (*powerxtypes.GetPaymentReply, error) {
	res := &powerxtypes.GetPaymentReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/mp/trade/payments/%v", req.PaymentId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradePayment) CreatePaymentFromOrder(ctx context.Context, req *powerxtypes.CreatePaymentFromOrderRequest) (*powerxtypes.CreatePaymentFromOrderRequestReply, error) {
	res := &powerxtypes.CreatePaymentFromOrderRequestReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/mp/trade/payments").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradePayment) UpdatePayment(ctx context.Context, req *powerxtypes.UpdatePaymentRequest) (*powerxtypes.UpdatePaymentReply, error) {
	res := &powerxtypes.UpdatePaymentReply{}
	err := c.H.Df().Method(http.MethodPut).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/mp/trade/payments/%v", req.PaymentId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
