package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"fmt"
	"net/http"
)

type AdminPermission struct {
	*PowerX
}

func (c *AdminPermission) ListRoles() (*powerxtypes.ListRolesReply, error) {
	res := &powerxtypes.ListRolesReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/permission/roles").
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminPermission) CreateRole(req *powerxtypes.CreateRoleRequest) (*powerxtypes.CreateRoleReply, error) {
	res := &powerxtypes.CreateRoleReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/permission/roles").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminPermission) GetRole(req *powerxtypes.GetRoleRequest) (*powerxtypes.GetRoleReply, error) {
	res := &powerxtypes.GetRoleReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/permission/roles/%v", req.RoleCode)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminPermission) PatchRole(req *powerxtypes.PatchRoleReqeust) (*powerxtypes.PatchRoleReply, error) {
	res := &powerxtypes.PatchRoleReply{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri(fmt.Sprintf("/api/v1/admin/permission/roles/%v", req.RoleCode)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminPermission) GetRoleEmployees(req *powerxtypes.GetRoleEmployeesReqeust) (*powerxtypes.GetRoleEmployeesReply, error) {
	res := &powerxtypes.GetRoleEmployeesReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/permission/roles/%v/users", req.RoleCode)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminPermission) SetRolePermissions(req *powerxtypes.SetRolePermissionsRequest) (*powerxtypes.SetRolePermissionsReply, error) {
	res := &powerxtypes.SetRolePermissionsReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri(fmt.Sprintf("/api/v1/admin/permission/roles/%v/actions/set-permissions", req.RoleCode)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminPermission) ListAPI(req *powerxtypes.ListAPIRequest) (*powerxtypes.ListAPIReply, error) {
	res := &powerxtypes.ListAPIReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/permission/api-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminPermission) SetRoleEmployees(req *powerxtypes.SetRoleEmployeesRequest) (*powerxtypes.SetRoleEmployeesReply, error) {
	res := &powerxtypes.SetRoleEmployeesReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri(fmt.Sprintf("/api/v1/admin/permission/roles/%v/actions/set-employees", req.RoleCode)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminPermission) SetUserRoles(req *powerxtypes.SetUserRolesRequest) (*powerxtypes.SetUserRolesReply, error) {
	res := &powerxtypes.SetUserRolesReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri(fmt.Sprintf("/api/v1/admin/permission/users/%v/actions/set-roles", req.UserId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
