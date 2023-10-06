package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"
	"fmt"
	"net/http"
)

type AdminBusinessOpportunity struct {
	*PowerX
}

func (c *AdminBusinessOpportunity) GetOpportunityList(ctx context.Context, req *powerxtypes.GetOpportunityListRequest) (*powerxtypes.GetOpportunityListReply, error) {
	res := &powerxtypes.GetOpportunityListReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/admin/business/opportunities").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminBusinessOpportunity) CreateOpportunity(ctx context.Context, req *powerxtypes.CreateOpportunityRequest) (*powerxtypes.CreateOpportunityReply, error) {
	res := &powerxtypes.CreateOpportunityReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/admin/business/opportunities").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminBusinessOpportunity) AssignEmployeeToOpportunity(ctx context.Context, req *powerxtypes.AssignEmployeeToOpportunityRequest) (*powerxtypes.AssignEmployeeToOpportunityReply, error) {
	res := &powerxtypes.AssignEmployeeToOpportunityReply{}
	err := c.H.Df().Method(http.MethodPut).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/business/opportunities/%v/assign-employee", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminBusinessOpportunity) UpdateOpportunity(ctx context.Context, req *powerxtypes.UpdateOpportunityRequest) (*powerxtypes.UpdateOpportunityReply, error) {
	res := &powerxtypes.UpdateOpportunityReply{}
	err := c.H.Df().Method(http.MethodPut).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/business/opportunities/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminBusinessOpportunity) DeleteOpportunity(ctx context.Context, req *powerxtypes.DeleteOpportunityRequest) (*powerxtypes.DeleteOpportunityReply, error) {
	res := &powerxtypes.DeleteOpportunityReply{}
	err := c.H.Df().Method(http.MethodDelete).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/business/opportunities/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
