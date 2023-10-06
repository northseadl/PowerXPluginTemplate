package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"net/http"
)

type AdminCommon struct {
	*PowerX
}

func (c *AdminCommon) GetEmployeeOptions(req *powerxtypes.GetEmployeeOptionsRequest) (*powerxtypes.GetEmployeeOptionsReply, error) {
	res := &powerxtypes.GetEmployeeOptionsReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/common/options/employees").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCommon) GetEmployeeQueryOptions() (*powerxtypes.GetEmployeeQueryOptionsReply, error) {
	res := &powerxtypes.GetEmployeeQueryOptionsReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/common/options/employee-query").
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCommon) GetDepartmentOptions(req *powerxtypes.GetDepartmentOptionsRequest) (*powerxtypes.GetDepartmentOptionsReply, error) {
	res := &powerxtypes.GetDepartmentOptionsReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/common/options/departments").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
