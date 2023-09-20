package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type AdminMediaresource struct {
	*PowerX
}

func (c *AdminMediaresource) ListMediaResources(req *types.ListMediaResourcesPageRequest) (*types.ListMediaResourcesPageReply, error) {
	res := &types.ListMediaResourcesPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/media/resources/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminMediaresource) CreateMediaResource() (*types.CreateMediaResourceReply, error) {
	res := &types.CreateMediaResourceReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/media/resources").
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminMediaresource) GetMediaResource(req *types.GetMediaResourceRequest) (*types.GetMediaResourceReply, error) {
	res := &types.GetMediaResourceReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/media/resources/%s", req.Id)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminMediaresource) DeleteMediaResource(req *types.DeleteMediaResourceRequest) (*types.DeleteMediaResourceReply, error) {
	res := &types.DeleteMediaResourceReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/media/resources/%s", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
