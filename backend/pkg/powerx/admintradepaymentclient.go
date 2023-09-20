package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type AdminTradePayment struct {
	*PowerX
}

func (c *AdminTradePayment) ListPaymentsPage(req *types.ListPaymentsPageRequest) (*types.ListPaymentsPageReply, error) {
	res := &types.ListPaymentsPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/trade/payments/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradePayment) GetPayment(req *types.GetPaymentRequest) (*types.GetPaymentReply, error) {
	res := &types.GetPaymentReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/trade/payments/%s", req.PaymentId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradePayment) CreatePayment(req *types.CreatePaymentRequest) (*types.CreatePaymentReply, error) {
	res := &types.CreatePaymentReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/trade/payments").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradePayment) PutPayment(req *types.PutPaymentRequest) (*types.PutPaymentReply, error) {
	res := &types.PutPaymentReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/trade/payments/%s", req.PaymentId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradePayment) PatchPayment(req *types.PatchPaymentRequest) (*types.PatchPaymentReply, error) {
	res := &types.PatchPaymentReply{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri(fmt.Sprintf("/api/v1/admin/trade/payments/%s", req.PaymentId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradePayment) DeletePayment(req *types.DeletePaymentRequest) (*types.DeletePaymentReply, error) {
	res := &types.DeletePaymentReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/trade/payments/%s", req.PaymentId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
