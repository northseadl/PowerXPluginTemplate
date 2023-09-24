package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type WebCustomerAuth struct {
	*PowerX
}

func (c *WebCustomerAuth) Login(req *types.CustomerLoginRequest) (*types.CustomerLoginAuthReply, error) {
	res := &types.CustomerLoginAuthReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/web/customer/login").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *WebCustomerAuth) RegisterCustomerByPhone(req *types.CustomerRegisterByPhoneRequest) (*types.CustomerRegisterByPhoneReply, error) {
	res := &types.CustomerRegisterByPhoneReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/web/customer/registerByPhone").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *WebCustomerAuth) UpdateCustomerProfile(req *types.UpdateCustomerProfileRequest) (*types.UpdateCustomerProfileReply, error) {
	res := &types.UpdateCustomerProfileReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri(fmt.Sprintf("/api/v1/web/customer/updateCustomerProfile/%v", req.CustomerId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
