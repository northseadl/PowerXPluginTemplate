package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"fmt"
	"net/http"
)

type AdminTradePayment struct {
	*PowerX
}

func (c *AdminTradePayment) ListPaymentsPage(req *powerxtypes.ListPaymentsPageRequest) (*powerxtypes.ListPaymentsPageReply, error) {
	res := &powerxtypes.ListPaymentsPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/trade/payments/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradePayment) GetPayment(req *powerxtypes.GetPaymentRequest) (*powerxtypes.GetPaymentReply, error) {
	res := &powerxtypes.GetPaymentReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/trade/payments/%v", req.PaymentId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradePayment) CreatePayment(req *powerxtypes.CreatePaymentRequest) (*powerxtypes.CreatePaymentReply, error) {
	res := &powerxtypes.CreatePaymentReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/trade/payments").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradePayment) PutPayment(req *powerxtypes.PutPaymentRequest) (*powerxtypes.PutPaymentReply, error) {
	res := &powerxtypes.PutPaymentReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/trade/payments/%v", req.PaymentId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradePayment) PatchPayment(req *powerxtypes.PatchPaymentRequest) (*powerxtypes.PatchPaymentReply, error) {
	res := &powerxtypes.PatchPaymentReply{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri(fmt.Sprintf("/api/v1/admin/trade/payments/%v", req.PaymentId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradePayment) DeletePayment(req *powerxtypes.DeletePaymentRequest) (*powerxtypes.DeletePaymentReply, error) {
	res := &powerxtypes.DeletePaymentReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/trade/payments/%v", req.PaymentId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
