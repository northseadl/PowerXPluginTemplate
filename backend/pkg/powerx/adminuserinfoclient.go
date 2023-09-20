package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	
	"net/http"
)

type AdminUserinfo struct {
	*PowerX
}


func (c *AdminUserinfo) GetUserInfo() (*types.GetUserInfoReply, error) {
	res := &types.GetUserInfoReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/user-center/user-info").
		
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminUserinfo) GetMenuRoles() (*types.GetMenuRolesReply, error) {
	res := &types.GetMenuRolesReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/user-center/menu-roles").
		
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminUserinfo) ModifyUserPassword(req *types.ModifyPasswordReqeust) (*types., error) {
	res := &types.{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/user-center/users/actions/modify-password").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
