package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"
	"fmt"
	"net/http"
)

type AdminDepartment struct {
	*PowerX
}

func (c *AdminDepartment) GetDepartmentTree(ctx context.Context, req *powerxtypes.GetDepartmentTreeRequest) (*powerxtypes.GetDepartmentTreeReply, error) {
	res := &powerxtypes.GetDepartmentTreeReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/department/department-tree/%v", req.DepId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDepartment) GetDepartment(ctx context.Context, req *powerxtypes.GetDepartmentRequest) (*powerxtypes.GetDepartmentReply, error) {
	res := &powerxtypes.GetDepartmentReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/department/departments/%v", req.Id)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDepartment) CreateDepartment(ctx context.Context, req *powerxtypes.CreateDepartmentRequest) (*powerxtypes.CreateDepartmentReply, error) {
	res := &powerxtypes.CreateDepartmentReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/admin/department/departments").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDepartment) PatchDepartment(ctx context.Context, req *powerxtypes.PatchDepartmentRequest) (*powerxtypes.PatchDepartmentReply, error) {
	res := &powerxtypes.PatchDepartmentReply{}
	err := c.H.Df().Method(http.MethodPatch).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/department/departments/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDepartment) DeleteDepartment(ctx context.Context, req *powerxtypes.DeleteDepartmentRequest) (*powerxtypes.DeleteDepartmentReply, error) {
	res := &powerxtypes.DeleteDepartmentReply{}
	err := c.H.Df().Method(http.MethodDelete).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/department/departments/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
