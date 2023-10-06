package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"fmt"
	"net/http"
)

type MpTradePayment struct {
	*PowerX
}

func (c *MpTradePayment) ListPaymentsPage(req *powerxtypes.ListPaymentsPageRequest) (*powerxtypes.ListPaymentsPageReply, error) {
	res := &powerxtypes.ListPaymentsPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/mp/trade/payments/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradePayment) GetPayment(req *powerxtypes.GetPaymentRequest) (*powerxtypes.GetPaymentReply, error) {
	res := &powerxtypes.GetPaymentReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/mp/trade/payments/%v", req.PaymentId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradePayment) CreatePaymentFromOrder(req *powerxtypes.CreatePaymentFromOrderRequest) (*powerxtypes.CreatePaymentFromOrderRequestReply, error) {
	res := &powerxtypes.CreatePaymentFromOrderRequestReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/mp/trade/payments").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradePayment) UpdatePayment(req *powerxtypes.UpdatePaymentRequest) (*powerxtypes.UpdatePaymentReply, error) {
	res := &powerxtypes.UpdatePaymentReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/mp/trade/payments/%v", req.PaymentId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
