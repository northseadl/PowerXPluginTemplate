package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"
	"fmt"
	"net/http"
)

type AdminScrmCustomer struct {
	*PowerX
}

func (c *AdminScrmCustomer) GetWeWorkCustomer(ctx context.Context, req *powerxtypes.GetWeWorkCustomerRequest) (*powerxtypes.GetWeWorkCustomerReply, error) {
	res := &powerxtypes.GetWeWorkCustomerReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/scrm/customer/customers/%v", req.Id)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminScrmCustomer) ListWeWorkCustomers(ctx context.Context, req *powerxtypes.ListWeWorkCustomersRequest) (*powerxtypes.ListWeWorkCustomersReply, error) {
	res := &powerxtypes.ListWeWorkCustomersReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/admin/scrm/customer/customers").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminScrmCustomer) PatchWeWorkCustomer(ctx context.Context, req *powerxtypes.PatchWeWorkCustomerRequest) (*powerxtypes.PatchWeWorkCustomerReply, error) {
	res := &powerxtypes.PatchWeWorkCustomerReply{}
	err := c.H.Df().Method(http.MethodPatch).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/scrm/customer/customers/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminScrmCustomer) SyncWeWorkCustomer(ctx context.Context) (*powerxtypes.SyncWeWorkCustomerReply, error) {
	res := &powerxtypes.SyncWeWorkCustomerReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/admin/scrm/customer/customers/actions/sync").
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
