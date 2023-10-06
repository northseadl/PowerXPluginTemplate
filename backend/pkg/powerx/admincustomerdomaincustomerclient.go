package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"
	"fmt"
	"net/http"
)

type AdminCustomerdomainCustomer struct {
	*PowerX
}

func (c *AdminCustomerdomainCustomer) GetCustomer(ctx context.Context, req *powerxtypes.GetCustomerReqeuest) (*powerxtypes.GetCustomerReply, error) {
	res := &powerxtypes.GetCustomerReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/customerdomain/customers/%v", req.Id)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCustomerdomainCustomer) ListCustomersPage(ctx context.Context, req *powerxtypes.ListCustomersPageRequest) (*powerxtypes.ListCustomersPageReply, error) {
	res := &powerxtypes.ListCustomersPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/admin/customerdomain/customers/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCustomerdomainCustomer) CreateCustomer(ctx context.Context, req *powerxtypes.CreateCustomerRequest) (*powerxtypes.CreateCustomerReply, error) {
	res := &powerxtypes.CreateCustomerReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/admin/customerdomain/customers").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCustomerdomainCustomer) PutCustomer(ctx context.Context, req *powerxtypes.PutCustomerRequest) (*powerxtypes.PutCustomerReply, error) {
	res := &powerxtypes.PutCustomerReply{}
	err := c.H.Df().Method(http.MethodPut).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/customerdomain/customers/%v", req.CustomerId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCustomerdomainCustomer) PatchCustomer(ctx context.Context, req *powerxtypes.PatchCustomerRequest) (*powerxtypes.PatchCustomerReply, error) {
	res := &powerxtypes.PatchCustomerReply{}
	err := c.H.Df().Method(http.MethodPatch).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/customerdomain/customers/%v", req.CustomerId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCustomerdomainCustomer) DeleteCustomer(ctx context.Context, req *powerxtypes.DeleteCustomerRequest) (*powerxtypes.DeleteCustomerReply, error) {
	res := &powerxtypes.DeleteCustomerReply{}
	err := c.H.Df().Method(http.MethodDelete).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/customerdomain/customers/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCustomerdomainCustomer) AssignCustomerToEmployee(ctx context.Context, req *powerxtypes.AssignCustomerToEmployeeRequest) (*powerxtypes.AssignCustomerToEmployeeReply, error) {
	res := &powerxtypes.AssignCustomerToEmployeeReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/customerdomain/customers/%v/actions/employees", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
