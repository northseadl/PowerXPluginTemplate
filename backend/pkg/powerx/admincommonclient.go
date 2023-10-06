package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"

	"net/http"
)

type AdminCommon struct {
	*PowerX
}

func (c *AdminCommon) GetEmployeeOptions(ctx context.Context, req *powerxtypes.GetEmployeeOptionsRequest) (*powerxtypes.GetEmployeeOptionsReply, error) {
	res := &powerxtypes.GetEmployeeOptionsReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/admin/common/options/employees").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCommon) GetEmployeeQueryOptions(ctx context.Context) (*powerxtypes.GetEmployeeQueryOptionsReply, error) {
	res := &powerxtypes.GetEmployeeQueryOptionsReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/admin/common/options/employee-query").
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCommon) GetDepartmentOptions(ctx context.Context, req *powerxtypes.GetDepartmentOptionsRequest) (*powerxtypes.GetDepartmentOptionsReply, error) {
	res := &powerxtypes.GetDepartmentOptionsReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/admin/common/options/departments").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
