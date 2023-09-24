package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type AdminEmployee struct {
	*PowerX
}

func (c *AdminEmployee) SyncEmployees(req *types.SyncEmployeesRequest) (*types.SyncEmployeesReply, error) {
	res := &types.SyncEmployeesReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/employee/employees/actions/sync").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminEmployee) GetEmployee(req *types.GetEmployeeRequest) (*types.GetEmployeeReply, error) {
	res := &types.GetEmployeeReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/employee/employees/%v", req.Id)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminEmployee) ListEmployees(req *types.ListEmployeesRequest) (*types.ListEmployeesReply, error) {
	res := &types.ListEmployeesReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/employee/employees").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminEmployee) CreateEmployee(req *types.CreateEmployeeRequest) (*types.CreateEmployeeReply, error) {
	res := &types.CreateEmployeeReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/employee/employees").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminEmployee) UpdateEmployee(req *types.UpdateEmployeeRequest) (*types.UpdateEmployeeReply, error) {
	res := &types.UpdateEmployeeReply{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri(fmt.Sprintf("/api/v1/admin/employee/employees/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminEmployee) DeleteEmployee(req *types.DeleteEmployeeRequest) (*types.DeleteEmployeeReply, error) {
	res := &types.DeleteEmployeeReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/employee/employees/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminEmployee) ResetPassword(req *types.ResetPasswordRequest) (*types.ResetPasswordReply, error) {
	res := &types.ResetPasswordReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/employee/employees/actions/reset-password").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
