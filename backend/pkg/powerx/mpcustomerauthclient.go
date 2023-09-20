package powerx

import (
	"PluginTemplate/pkg/powerx/types"

	"net/http"
)

type MpCustomerAuth struct {
	*PowerX
}

func (c *MpCustomerAuth) Login(req *types.MPCustomerLoginRequest) (*types.MPCustomerLoginAuthReply, error) {
	res := &types.MPCustomerLoginAuthReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/mp/customer/login").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpCustomerAuth) AuthByPhone(req *types.MPCustomerAuthRequest) (*types.MPCustomerLoginAuthReply, error) {
	res := &types.MPCustomerLoginAuthReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/mp/customer/authByPhone").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpCustomerAuth) AuthByProfile() (*types.MPCustomerLoginAuthReply, error) {
	res := &types.MPCustomerLoginAuthReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/mp/customer/authByProfile").
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
