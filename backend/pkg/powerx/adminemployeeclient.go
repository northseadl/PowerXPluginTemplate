package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"
	"fmt"
	"net/http"
)

type AdminEmployee struct {
	*PowerX
}

func (c *AdminEmployee) SyncEmployees(ctx context.Context, req *powerxtypes.SyncEmployeesRequest) (*powerxtypes.SyncEmployeesReply, error) {
	res := &powerxtypes.SyncEmployeesReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/admin/employee/employees/actions/sync").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminEmployee) GetEmployee(ctx context.Context, req *powerxtypes.GetEmployeeRequest) (*powerxtypes.GetEmployeeReply, error) {
	res := &powerxtypes.GetEmployeeReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/employee/employees/%v", req.Id)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminEmployee) ListEmployees(ctx context.Context, req *powerxtypes.ListEmployeesRequest) (*powerxtypes.ListEmployeesReply, error) {
	res := &powerxtypes.ListEmployeesReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/admin/employee/employees").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminEmployee) CreateEmployee(ctx context.Context, req *powerxtypes.CreateEmployeeRequest) (*powerxtypes.CreateEmployeeReply, error) {
	res := &powerxtypes.CreateEmployeeReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/admin/employee/employees").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminEmployee) UpdateEmployee(ctx context.Context, req *powerxtypes.UpdateEmployeeRequest) (*powerxtypes.UpdateEmployeeReply, error) {
	res := &powerxtypes.UpdateEmployeeReply{}
	err := c.H.Df().Method(http.MethodPatch).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/employee/employees/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminEmployee) DeleteEmployee(ctx context.Context, req *powerxtypes.DeleteEmployeeRequest) (*powerxtypes.DeleteEmployeeReply, error) {
	res := &powerxtypes.DeleteEmployeeReply{}
	err := c.H.Df().Method(http.MethodDelete).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/employee/employees/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminEmployee) ResetPassword(ctx context.Context, req *powerxtypes.ResetPasswordRequest) (*powerxtypes.ResetPasswordReply, error) {
	res := &powerxtypes.ResetPasswordReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/admin/employee/employees/actions/reset-password").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
