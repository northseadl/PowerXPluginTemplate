package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"fmt"
	"net/http"
)

type AdminMediaresource struct {
	*PowerX
}

func (c *AdminMediaresource) ListMediaResources(req *powerxtypes.ListMediaResourcesPageRequest) (*powerxtypes.ListMediaResourcesPageReply, error) {
	res := &powerxtypes.ListMediaResourcesPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/media/resources/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminMediaresource) CreateMediaResource() (*powerxtypes.CreateMediaResourceReply, error) {
	res := &powerxtypes.CreateMediaResourceReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/media/resources").
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminMediaresource) GetMediaResource(req *powerxtypes.GetMediaResourceRequest) (*powerxtypes.GetMediaResourceReply, error) {
	res := &powerxtypes.GetMediaResourceReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/media/resources/%v", req.Id)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminMediaresource) DeleteMediaResource(req *powerxtypes.DeleteMediaResourceRequest) (*powerxtypes.DeleteMediaResourceReply, error) {
	res := &powerxtypes.DeleteMediaResourceReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/media/resources/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
