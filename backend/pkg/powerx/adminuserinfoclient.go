package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"

	"net/http"
)

type AdminUserinfo struct {
	*PowerX
}

func (c *AdminUserinfo) GetUserInfo(ctx context.Context) (*powerxtypes.GetUserInfoReply, error) {
	res := &powerxtypes.GetUserInfoReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/admin/user-center/user-info").
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminUserinfo) GetMenuRoles(ctx context.Context) (*powerxtypes.GetMenuRolesReply, error) {
	res := &powerxtypes.GetMenuRolesReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/admin/user-center/menu-roles").
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
