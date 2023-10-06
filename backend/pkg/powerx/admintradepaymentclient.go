package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"
	"fmt"
	"net/http"
)

type AdminTradePayment struct {
	*PowerX
}

func (c *AdminTradePayment) ListPaymentsPage(ctx context.Context, req *powerxtypes.ListPaymentsPageRequest) (*powerxtypes.ListPaymentsPageReply, error) {
	res := &powerxtypes.ListPaymentsPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/admin/trade/payments/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradePayment) GetPayment(ctx context.Context, req *powerxtypes.GetPaymentRequest) (*powerxtypes.GetPaymentReply, error) {
	res := &powerxtypes.GetPaymentReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/trade/payments/%v", req.PaymentId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradePayment) CreatePayment(ctx context.Context, req *powerxtypes.CreatePaymentRequest) (*powerxtypes.CreatePaymentReply, error) {
	res := &powerxtypes.CreatePaymentReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/admin/trade/payments").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradePayment) PutPayment(ctx context.Context, req *powerxtypes.PutPaymentRequest) (*powerxtypes.PutPaymentReply, error) {
	res := &powerxtypes.PutPaymentReply{}
	err := c.H.Df().Method(http.MethodPut).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/trade/payments/%v", req.PaymentId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradePayment) PatchPayment(ctx context.Context, req *powerxtypes.PatchPaymentRequest) (*powerxtypes.PatchPaymentReply, error) {
	res := &powerxtypes.PatchPaymentReply{}
	err := c.H.Df().Method(http.MethodPatch).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/trade/payments/%v", req.PaymentId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradePayment) DeletePayment(ctx context.Context, req *powerxtypes.DeletePaymentRequest) (*powerxtypes.DeletePaymentReply, error) {
	res := &powerxtypes.DeletePaymentReply{}
	err := c.H.Df().Method(http.MethodDelete).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/trade/payments/%v", req.PaymentId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
