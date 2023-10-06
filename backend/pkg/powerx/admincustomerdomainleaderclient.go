package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"
	"fmt"
	"net/http"
)

type AdminCustomerdomainLeader struct {
	*PowerX
}

func (c *AdminCustomerdomainLeader) GetLead(ctx context.Context, req *powerxtypes.GetLeadReqeuest) (*powerxtypes.GetLeadReply, error) {
	res := &powerxtypes.GetLeadReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/customerdomain/leads/%v", req.Id)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCustomerdomainLeader) ListLeadsPage(ctx context.Context, req *powerxtypes.ListLeadsPageRequest) (*powerxtypes.ListLeadsPageReply, error) {
	res := &powerxtypes.ListLeadsPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/admin/customerdomain/leads/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCustomerdomainLeader) CreateLead(ctx context.Context, req *powerxtypes.CreateLeadRequest) (*powerxtypes.CreateLeadReply, error) {
	res := &powerxtypes.CreateLeadReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/admin/customerdomain/leads").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCustomerdomainLeader) PutLead(ctx context.Context, req *powerxtypes.PutLeadRequest) (*powerxtypes.PutLeadReply, error) {
	res := &powerxtypes.PutLeadReply{}
	err := c.H.Df().Method(http.MethodPut).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/customerdomain/leads/%v", req.LeadId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCustomerdomainLeader) PatchLead(ctx context.Context, req *powerxtypes.PatchLeadRequest) (*powerxtypes.PatchLeadReply, error) {
	res := &powerxtypes.PatchLeadReply{}
	err := c.H.Df().Method(http.MethodPatch).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/customerdomain/leads/%v", req.LeadId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCustomerdomainLeader) DeleteLead(ctx context.Context, req *powerxtypes.DeleteLeadRequest) (*powerxtypes.DeleteLeadReply, error) {
	res := &powerxtypes.DeleteLeadReply{}
	err := c.H.Df().Method(http.MethodDelete).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/customerdomain/leads/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCustomerdomainLeader) AssignLeadToEmployee(ctx context.Context, req *powerxtypes.AssignLeadToEmployeeRequest) (*powerxtypes.AssignLeadToEmployeeReply, error) {
	res := &powerxtypes.AssignLeadToEmployeeReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/customerdomain/leads/%v/actions/employees", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
