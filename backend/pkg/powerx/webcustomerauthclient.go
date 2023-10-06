package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"fmt"
	"net/http"
)

type WebCustomerAuth struct {
	*PowerX
}

func (c *WebCustomerAuth) Login(req *powerxtypes.CustomerLoginRequest) (*powerxtypes.CustomerLoginAuthReply, error) {
	res := &powerxtypes.CustomerLoginAuthReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/web/customer/login").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *WebCustomerAuth) RegisterCustomerByPhone(req *powerxtypes.CustomerRegisterByPhoneRequest) (*powerxtypes.CustomerRegisterByPhoneReply, error) {
	res := &powerxtypes.CustomerRegisterByPhoneReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/web/customer/registerByPhone").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *WebCustomerAuth) UpdateCustomerProfile(req *powerxtypes.UpdateCustomerProfileRequest) (*powerxtypes.UpdateCustomerProfileReply, error) {
	res := &powerxtypes.UpdateCustomerProfileReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri(fmt.Sprintf("/api/v1/web/customer/updateCustomerProfile/%v", req.CustomerId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
