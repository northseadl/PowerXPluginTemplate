package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"

	"net/http"
)

type AdminScrmContact struct {
	*PowerX
}

func (c *AdminScrmContact) SyncWeWorkContact(ctx context.Context) (*powerxtypes.SyncWeWorkContactReply, error) {
	res := &powerxtypes.SyncWeWorkContactReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/admin/scrm/contact/contacts/actions/sync").
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminScrmContact) ListWeWorkEmployee(ctx context.Context, req *powerxtypes.ListWeWorkEmployeeReqeust) (*powerxtypes.ListWeWorkEmployeeReply, error) {
	res := &powerxtypes.ListWeWorkEmployeeReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/admin/scrm/contact/employees").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
