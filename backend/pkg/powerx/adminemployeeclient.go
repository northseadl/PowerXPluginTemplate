package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"fmt"
	"net/http"
)

type AdminEmployee struct {
	*PowerX
}

func (c *AdminEmployee) SyncEmployees(req *powerxtypes.SyncEmployeesRequest) (*powerxtypes.SyncEmployeesReply, error) {
	res := &powerxtypes.SyncEmployeesReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/employee/employees/actions/sync").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminEmployee) GetEmployee(req *powerxtypes.GetEmployeeRequest) (*powerxtypes.GetEmployeeReply, error) {
	res := &powerxtypes.GetEmployeeReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/employee/employees/%v", req.Id)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminEmployee) ListEmployees(req *powerxtypes.ListEmployeesRequest) (*powerxtypes.ListEmployeesReply, error) {
	res := &powerxtypes.ListEmployeesReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/employee/employees").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminEmployee) CreateEmployee(req *powerxtypes.CreateEmployeeRequest) (*powerxtypes.CreateEmployeeReply, error) {
	res := &powerxtypes.CreateEmployeeReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/employee/employees").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminEmployee) UpdateEmployee(req *powerxtypes.UpdateEmployeeRequest) (*powerxtypes.UpdateEmployeeReply, error) {
	res := &powerxtypes.UpdateEmployeeReply{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri(fmt.Sprintf("/api/v1/admin/employee/employees/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminEmployee) DeleteEmployee(req *powerxtypes.DeleteEmployeeRequest) (*powerxtypes.DeleteEmployeeReply, error) {
	res := &powerxtypes.DeleteEmployeeReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/employee/employees/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminEmployee) ResetPassword(req *powerxtypes.ResetPasswordRequest) (*powerxtypes.ResetPasswordReply, error) {
	res := &powerxtypes.ResetPasswordReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/employee/employees/actions/reset-password").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
