package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type AdminCustomerdomainCustomer struct {
	*PowerX
}

func (c *AdminCustomerdomainCustomer) GetCustomer(req *types.GetCustomerReqeuest) (*types.GetCustomerReply, error) {
	res := &types.GetCustomerReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/customerdomain/customers/%v", req.Id)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCustomerdomainCustomer) ListCustomersPage(req *types.ListCustomersPageRequest) (*types.ListCustomersPageReply, error) {
	res := &types.ListCustomersPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/customerdomain/customers/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCustomerdomainCustomer) CreateCustomer(req *types.CreateCustomerRequest) (*types.CreateCustomerReply, error) {
	res := &types.CreateCustomerReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/customerdomain/customers").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCustomerdomainCustomer) PutCustomer(req *types.PutCustomerRequest) (*types.PutCustomerReply, error) {
	res := &types.PutCustomerReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/customerdomain/customers/%v", req.CustomerId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCustomerdomainCustomer) PatchCustomer(req *types.PatchCustomerRequest) (*types.PatchCustomerReply, error) {
	res := &types.PatchCustomerReply{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri(fmt.Sprintf("/api/v1/admin/customerdomain/customers/%v", req.CustomerId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCustomerdomainCustomer) DeleteCustomer(req *types.DeleteCustomerRequest) (*types.DeleteCustomerReply, error) {
	res := &types.DeleteCustomerReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/customerdomain/customers/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminCustomerdomainCustomer) AssignCustomerToEmployee(req *types.AssignCustomerToEmployeeRequest) (*types.AssignCustomerToEmployeeReply, error) {
	res := &types.AssignCustomerToEmployeeReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri(fmt.Sprintf("/api/v1/admin/customerdomain/customers/%v/actions/employees", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
