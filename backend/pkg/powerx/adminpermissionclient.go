package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type AdminPermission struct {
	*PowerX
}

func (c *AdminPermission) ListRoles() (*types.ListRolesReply, error) {
	res := &types.ListRolesReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/permission/roles").
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminPermission) CreateRole(req *types.CreateRoleRequest) (*types.CreateRoleReply, error) {
	res := &types.CreateRoleReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/permission/roles").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminPermission) GetRole(req *types.GetRoleRequest) (*types.GetRoleReply, error) {
	res := &types.GetRoleReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/permission/roles/%s", req.RoleCode)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminPermission) PatchRole(req *types.PatchRoleReqeust) (*types.PatchRoleReply, error) {
	res := &types.PatchRoleReply{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri(fmt.Sprintf("/api/v1/admin/permission/roles/%s", req.RoleCode)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminPermission) GetRoleEmployees(req *types.GetRoleEmployeesReqeust) (*types.GetRoleEmployeesReply, error) {
	res := &types.GetRoleEmployeesReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/permission/roles/%s/users", req.RoleCode)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminPermission) SetRolePermissions(req *types.SetRolePermissionsRequest) (*types.SetRolePermissionsReply, error) {
	res := &types.SetRolePermissionsReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri(fmt.Sprintf("/api/v1/admin/permission/roles/%s/actions/set-permissions", req.RoleCode)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminPermission) ListAPI(req *types.ListAPIRequest) (*types.ListAPIReply, error) {
	res := &types.ListAPIReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/permission/api-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminPermission) SetRoleEmployees(req *types.SetRoleEmployeesRequest) (*types.SetRoleEmployeesReply, error) {
	res := &types.SetRoleEmployeesReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri(fmt.Sprintf("/api/v1/admin/permission/roles/%s/actions/set-employees", req.RoleCode)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminPermission) SetUserRoles(req *types.SetUserRolesRequest) (*types.SetUserRolesReply, error) {
	res := &types.SetUserRolesReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri(fmt.Sprintf("/api/v1/admin/permission/users/%s/actions/set-roles", req.UserId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
