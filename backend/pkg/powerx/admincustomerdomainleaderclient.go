package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type AdminCustomerdomainLeader struct {
	*PowerX
}

func (c *AdminCustomerdomainLeader) GetLead(req *types.GetLeadReqeuest) (*types.GetLeadReply, error) {
	res := &types.GetLeadReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/customerdomain/leads/%v", req.Id)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCustomerdomainLeader) ListLeadsPage(req *types.ListLeadsPageRequest) (*types.ListLeadsPageReply, error) {
	res := &types.ListLeadsPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/customerdomain/leads/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCustomerdomainLeader) CreateLead(req *types.CreateLeadRequest) (*types.CreateLeadReply, error) {
	res := &types.CreateLeadReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/customerdomain/leads").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCustomerdomainLeader) PutLead(req *types.PutLeadRequest) (*types.PutLeadReply, error) {
	res := &types.PutLeadReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/customerdomain/leads/%v", req.LeadId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCustomerdomainLeader) PatchLead(req *types.PatchLeadRequest) (*types.PatchLeadReply, error) {
	res := &types.PatchLeadReply{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri(fmt.Sprintf("/api/v1/admin/customerdomain/leads/%v", req.LeadId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCustomerdomainLeader) DeleteLead(req *types.DeleteLeadRequest) (*types.DeleteLeadReply, error) {
	res := &types.DeleteLeadReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/customerdomain/leads/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCustomerdomainLeader) AssignLeadToEmployee(req *types.AssignLeadToEmployeeRequest) (*types.AssignLeadToEmployeeReply, error) {
	res := &types.AssignLeadToEmployeeReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri(fmt.Sprintf("/api/v1/admin/customerdomain/leads/%v/actions/employees", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
