package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"fmt"
	"net/http"
)

type AdminDepartment struct {
	*PowerX
}

func (c *AdminDepartment) GetDepartmentTree(req *powerxtypes.GetDepartmentTreeRequest) (*powerxtypes.GetDepartmentTreeReply, error) {
	res := &powerxtypes.GetDepartmentTreeReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/department/department-tree/%v", req.DepId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDepartment) GetDepartment(req *powerxtypes.GetDepartmentRequest) (*powerxtypes.GetDepartmentReply, error) {
	res := &powerxtypes.GetDepartmentReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/department/departments/%v", req.Id)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDepartment) CreateDepartment(req *powerxtypes.CreateDepartmentRequest) (*powerxtypes.CreateDepartmentReply, error) {
	res := &powerxtypes.CreateDepartmentReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/department/departments").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDepartment) PatchDepartment(req *powerxtypes.PatchDepartmentRequest) (*powerxtypes.PatchDepartmentReply, error) {
	res := &powerxtypes.PatchDepartmentReply{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri(fmt.Sprintf("/api/v1/admin/department/departments/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDepartment) DeleteDepartment(req *powerxtypes.DeleteDepartmentRequest) (*powerxtypes.DeleteDepartmentReply, error) {
	res := &powerxtypes.DeleteDepartmentReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/department/departments/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
