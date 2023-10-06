package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"
	"fmt"
	"net/http"
)

type AdminPermission struct {
	*PowerX
}

func (c *AdminPermission) ListRoles(ctx context.Context) (*powerxtypes.ListRolesReply, error) {
	res := &powerxtypes.ListRolesReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/admin/permission/roles").
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminPermission) CreateRole(ctx context.Context, req *powerxtypes.CreateRoleRequest) (*powerxtypes.CreateRoleReply, error) {
	res := &powerxtypes.CreateRoleReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/admin/permission/roles").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminPermission) GetRole(ctx context.Context, req *powerxtypes.GetRoleRequest) (*powerxtypes.GetRoleReply, error) {
	res := &powerxtypes.GetRoleReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/permission/roles/%v", req.RoleCode)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminPermission) PatchRole(ctx context.Context, req *powerxtypes.PatchRoleReqeust) (*powerxtypes.PatchRoleReply, error) {
	res := &powerxtypes.PatchRoleReply{}
	err := c.H.Df().Method(http.MethodPatch).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/permission/roles/%v", req.RoleCode)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminPermission) GetRoleEmployees(ctx context.Context, req *powerxtypes.GetRoleEmployeesReqeust) (*powerxtypes.GetRoleEmployeesReply, error) {
	res := &powerxtypes.GetRoleEmployeesReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/permission/roles/%v/users", req.RoleCode)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminPermission) SetRolePermissions(ctx context.Context, req *powerxtypes.SetRolePermissionsRequest) (*powerxtypes.SetRolePermissionsReply, error) {
	res := &powerxtypes.SetRolePermissionsReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/permission/roles/%v/actions/set-permissions", req.RoleCode)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminPermission) ListAPI(ctx context.Context, req *powerxtypes.ListAPIRequest) (*powerxtypes.ListAPIReply, error) {
	res := &powerxtypes.ListAPIReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/admin/permission/api-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminPermission) SetRoleEmployees(ctx context.Context, req *powerxtypes.SetRoleEmployeesRequest) (*powerxtypes.SetRoleEmployeesReply, error) {
	res := &powerxtypes.SetRoleEmployeesReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/permission/roles/%v/actions/set-employees", req.RoleCode)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminPermission) SetUserRoles(ctx context.Context, req *powerxtypes.SetUserRolesRequest) (*powerxtypes.SetUserRolesReply, error) {
	res := &powerxtypes.SetUserRolesReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/permission/users/%v/actions/set-roles", req.UserId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
