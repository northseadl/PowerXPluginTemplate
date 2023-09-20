package powerx

import (
	"PluginTemplate/pkg/powerx/types"

	"net/http"
)

type WebCustomerAuthOa struct {
	*PowerX
}

func (c *WebCustomerAuthOa) OALogin(req *types.OACustomerLoginRequest) (*types.OACustomerLoginAuthReply, error) {
	res := &types.OACustomerLoginAuthReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/web/customer/oa/login").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *WebCustomerAuthOa) AuthByPhone(req *types.OACustomerAuthRequest) (*types.OACustomerLoginAuthReply, error) {
	res := &types.OACustomerLoginAuthReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/web/customer/oa/authByPhone").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *WebCustomerAuthOa) AuthByProfile() (*types.OACustomerLoginAuthReply, error) {
	res := &types.OACustomerLoginAuthReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/web/customer/oa/authByProfile").
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
