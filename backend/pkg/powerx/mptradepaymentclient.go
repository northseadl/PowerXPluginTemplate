package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type MpTradePayment struct {
	*PowerX
}

func (c *MpTradePayment) ListPaymentsPage(req *types.ListPaymentsPageRequest) (*types.ListPaymentsPageReply, error) {
	res := &types.ListPaymentsPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/mp/trade/payments/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradePayment) GetPayment(req *types.GetPaymentRequest) (*types.GetPaymentReply, error) {
	res := &types.GetPaymentReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/mp/trade/payments/%v", req.PaymentId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradePayment) CreatePaymentFromOrder(req *types.CreatePaymentFromOrderRequest) (*types.CreatePaymentFromOrderRequestReply, error) {
	res := &types.CreatePaymentFromOrderRequestReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/mp/trade/payments").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpTradePayment) UpdatePayment(req *types.UpdatePaymentRequest) (*types.UpdatePaymentReply, error) {
	res := &types.UpdatePaymentReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/mp/trade/payments/%v", req.PaymentId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
