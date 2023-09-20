package powerx

import (
	"PluginTemplate/pkg/powerx/types"

	"net/http"
)

type AdminCommon struct {
	*PowerX
}

func (c *AdminCommon) GetEmployeeOptions(req *types.GetEmployeeOptionsRequest) (*types.GetEmployeeOptionsReply, error) {
	res := &types.GetEmployeeOptionsReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/common/options/employees").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCommon) GetEmployeeQueryOptions() (*types.GetEmployeeQueryOptionsReply, error) {
	res := &types.GetEmployeeQueryOptionsReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/common/options/employee-query").
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCommon) GetDepartmentOptions(req *types.GetDepartmentOptionsRequest) (*types.GetDepartmentOptionsReply, error) {
	res := &types.GetDepartmentOptionsReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/common/options/departments").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
