package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"net/http"
)

type MpCustomerAuth struct {
	*PowerX
}

func (c *MpCustomerAuth) Login(req *powerxtypes.MPCustomerLoginRequest) (*powerxtypes.MPCustomerLoginAuthReply, error) {
	res := &powerxtypes.MPCustomerLoginAuthReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/mp/customer/login").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpCustomerAuth) AuthByPhone(req *powerxtypes.MPCustomerAuthRequest) (*powerxtypes.MPCustomerLoginAuthReply, error) {
	res := &powerxtypes.MPCustomerLoginAuthReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/mp/customer/authByPhone").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpCustomerAuth) AuthByProfile() (*powerxtypes.MPCustomerLoginAuthReply, error) {
	res := &powerxtypes.MPCustomerLoginAuthReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/mp/customer/authByProfile").
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
