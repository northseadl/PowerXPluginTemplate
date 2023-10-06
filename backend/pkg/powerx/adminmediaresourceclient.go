package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"
	"fmt"
	"net/http"
)

type AdminMediaresource struct {
	*PowerX
}

func (c *AdminMediaresource) ListMediaResources(ctx context.Context, req *powerxtypes.ListMediaResourcesPageRequest) (*powerxtypes.ListMediaResourcesPageReply, error) {
	res := &powerxtypes.ListMediaResourcesPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/admin/media/resources/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminMediaresource) CreateMediaResource(ctx context.Context) (*powerxtypes.CreateMediaResourceReply, error) {
	res := &powerxtypes.CreateMediaResourceReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/admin/media/resources").
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminMediaresource) GetMediaResource(ctx context.Context, req *powerxtypes.GetMediaResourceRequest) (*powerxtypes.GetMediaResourceReply, error) {
	res := &powerxtypes.GetMediaResourceReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/media/resources/%v", req.Id)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminMediaresource) DeleteMediaResource(ctx context.Context, req *powerxtypes.DeleteMediaResourceRequest) (*powerxtypes.DeleteMediaResourceReply, error) {
	res := &powerxtypes.DeleteMediaResourceReply{}
	err := c.H.Df().Method(http.MethodDelete).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/media/resources/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
