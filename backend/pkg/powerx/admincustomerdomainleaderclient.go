package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"fmt"
	"net/http"
)

type AdminCustomerdomainLeader struct {
	*PowerX
}

func (c *AdminCustomerdomainLeader) GetLead(req *powerxtypes.GetLeadReqeuest) (*powerxtypes.GetLeadReply, error) {
	res := &powerxtypes.GetLeadReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/customerdomain/leads/%v", req.Id)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCustomerdomainLeader) ListLeadsPage(req *powerxtypes.ListLeadsPageRequest) (*powerxtypes.ListLeadsPageReply, error) {
	res := &powerxtypes.ListLeadsPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/customerdomain/leads/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCustomerdomainLeader) CreateLead(req *powerxtypes.CreateLeadRequest) (*powerxtypes.CreateLeadReply, error) {
	res := &powerxtypes.CreateLeadReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/customerdomain/leads").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCustomerdomainLeader) PutLead(req *powerxtypes.PutLeadRequest) (*powerxtypes.PutLeadReply, error) {
	res := &powerxtypes.PutLeadReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/customerdomain/leads/%v", req.LeadId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCustomerdomainLeader) PatchLead(req *powerxtypes.PatchLeadRequest) (*powerxtypes.PatchLeadReply, error) {
	res := &powerxtypes.PatchLeadReply{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri(fmt.Sprintf("/api/v1/admin/customerdomain/leads/%v", req.LeadId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCustomerdomainLeader) DeleteLead(req *powerxtypes.DeleteLeadRequest) (*powerxtypes.DeleteLeadReply, error) {
	res := &powerxtypes.DeleteLeadReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/customerdomain/leads/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCustomerdomainLeader) AssignLeadToEmployee(req *powerxtypes.AssignLeadToEmployeeRequest) (*powerxtypes.AssignLeadToEmployeeReply, error) {
	res := &powerxtypes.AssignLeadToEmployeeReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri(fmt.Sprintf("/api/v1/admin/customerdomain/leads/%v/actions/employees", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
