package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type AdminDepartment struct {
	*PowerX
}

func (c *AdminDepartment) GetDepartmentTree(req *types.GetDepartmentTreeRequest) (*types.GetDepartmentTreeReply, error) {
	res := &types.GetDepartmentTreeReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/department/department-tree/%s", req.DepId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDepartment) GetDepartment(req *types.GetDepartmentRequest) (*types.GetDepartmentReply, error) {
	res := &types.GetDepartmentReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/department/departments/%s", req.Id)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDepartment) CreateDepartment(req *types.CreateDepartmentRequest) (*types.CreateDepartmentReply, error) {
	res := &types.CreateDepartmentReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/department/departments").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDepartment) PatchDepartment(req *types.PatchDepartmentRequest) (*types.PatchDepartmentReply, error) {
	res := &types.PatchDepartmentReply{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri(fmt.Sprintf("/api/v1/admin/department/departments/%s", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDepartment) DeleteDepartment(req *types.DeleteDepartmentRequest) (*types.DeleteDepartmentReply, error) {
	res := &types.DeleteDepartmentReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/department/departments/%s", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
