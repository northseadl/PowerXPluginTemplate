package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"fmt"
	"net/http"
)

type AdminScrmCustomer struct {
	*PowerX
}

func (c *AdminScrmCustomer) GetWeWorkCustomer(req *powerxtypes.GetWeWorkCustomerRequest) (*powerxtypes.GetWeWorkCustomerReply, error) {
	res := &powerxtypes.GetWeWorkCustomerReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/scrm/customer/customers/%v", req.Id)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminScrmCustomer) ListWeWorkCustomers(req *powerxtypes.ListWeWorkCustomersRequest) (*powerxtypes.ListWeWorkCustomersReply, error) {
	res := &powerxtypes.ListWeWorkCustomersReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/scrm/customer/customers").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminScrmCustomer) PatchWeWorkCustomer(req *powerxtypes.PatchWeWorkCustomerRequest) (*powerxtypes.PatchWeWorkCustomerReply, error) {
	res := &powerxtypes.PatchWeWorkCustomerReply{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri(fmt.Sprintf("/api/v1/admin/scrm/customer/customers/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminScrmCustomer) SyncWeWorkCustomer() (*powerxtypes.SyncWeWorkCustomerReply, error) {
	res := &powerxtypes.SyncWeWorkCustomerReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/scrm/customer/customers/actions/sync").
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
