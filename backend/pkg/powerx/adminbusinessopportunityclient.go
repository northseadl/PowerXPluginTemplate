package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type AdminBusinessOpportunity struct {
	*PowerX
}

func (c *AdminBusinessOpportunity) GetOpportunityList(req *types.GetOpportunityListRequest) (*types.GetOpportunityListReply, error) {
	res := &types.GetOpportunityListReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/business/opportunities").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminBusinessOpportunity) CreateOpportunity(req *types.CreateOpportunityRequest) (*types.CreateOpportunityReply, error) {
	res := &types.CreateOpportunityReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/business/opportunities").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminBusinessOpportunity) AssignEmployeeToOpportunity(req *types.AssignEmployeeToOpportunityRequest) (*types.AssignEmployeeToOpportunityReply, error) {
	res := &types.AssignEmployeeToOpportunityReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/business/opportunities/%s/assign-employee", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminBusinessOpportunity) UpdateOpportunity(req *types.UpdateOpportunityRequest) (*types.UpdateOpportunityReply, error) {
	res := &types.UpdateOpportunityReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/business/opportunities/%s", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminBusinessOpportunity) DeleteOpportunity(req *types.DeleteOpportunityRequest) (*types.DeleteOpportunityReply, error) {
	res := &types.DeleteOpportunityReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/business/opportunities/%s", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
