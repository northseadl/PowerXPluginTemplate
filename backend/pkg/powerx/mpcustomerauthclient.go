package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"

	"net/http"
)

type MpCustomerAuth struct {
	*PowerX
}

func (c *MpCustomerAuth) Login(ctx context.Context, req *powerxtypes.MPCustomerLoginRequest) (*powerxtypes.MPCustomerLoginAuthReply, error) {
	res := &powerxtypes.MPCustomerLoginAuthReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/mp/customer/login").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpCustomerAuth) AuthByPhone(ctx context.Context, req *powerxtypes.MPCustomerAuthRequest) (*powerxtypes.MPCustomerLoginAuthReply, error) {
	res := &powerxtypes.MPCustomerLoginAuthReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/mp/customer/authByPhone").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpCustomerAuth) AuthByProfile(ctx context.Context) (*powerxtypes.MPCustomerLoginAuthReply, error) {
	res := &powerxtypes.MPCustomerLoginAuthReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/mp/customer/authByProfile").
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
