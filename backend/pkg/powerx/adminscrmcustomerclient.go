package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type AdminScrmCustomer struct {
	*PowerX
}

func (c *AdminScrmCustomer) GetWeWorkCustomer(req *types.GetWeWorkCustomerRequest) (*types.GetWeWorkCustomerReply, error) {
	res := &types.GetWeWorkCustomerReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/scrm/customer/customers/%s", req.Id)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminScrmCustomer) ListWeWorkCustomers(req *types.ListWeWorkCustomersRequest) (*types.ListWeWorkCustomersReply, error) {
	res := &types.ListWeWorkCustomersReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/scrm/customer/customers").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminScrmCustomer) PatchWeWorkCustomer(req *types.PatchWeWorkCustomerRequest) (*types.PatchWeWorkCustomerReply, error) {
	res := &types.PatchWeWorkCustomerReply{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri(fmt.Sprintf("/api/v1/admin/scrm/customer/customers/%s", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminScrmCustomer) SyncWeWorkCustomer() (*types.SyncWeWorkCustomerReply, error) {
	res := &types.SyncWeWorkCustomerReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/scrm/customer/customers/actions/sync").
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
