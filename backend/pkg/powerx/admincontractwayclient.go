package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"fmt"
	"net/http"
)

type AdminContractway struct {
	*PowerX
}

func (c *AdminContractway) GetContractWayGroupTree(req *powerxtypes.GetContractWayGroupTreeRequest) (*powerxtypes.GetContractWayGroupTreeReply, error) {
	res := &powerxtypes.GetContractWayGroupTreeReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/contract-way/group-tree").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminContractway) GetContractWayGroupList(req *powerxtypes.GetContractWayGroupListRequest) (*powerxtypes.GetContractWayGroupListReply, error) {
	res := &powerxtypes.GetContractWayGroupListReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/contract-way/groups").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminContractway) GetContractWays(req *powerxtypes.GetContractWaysRequest) (*powerxtypes.GetContractWaysReply, error) {
	res := &powerxtypes.GetContractWaysReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/contract-way").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminContractway) CreateContractWay(req *powerxtypes.CreateContractWayRequest) (*powerxtypes.CreateContractWayReply, error) {
	res := &powerxtypes.CreateContractWayReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/contract-way").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminContractway) UpdateContractWay(req *powerxtypes.UpdateContractWayRequest) (*powerxtypes.UpdateContractWayReply, error) {
	res := &powerxtypes.UpdateContractWayReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/contract-way/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminContractway) DeleteContractWay(req *powerxtypes.DeleteContractWayRequest) (*powerxtypes.DeleteContractWayReply, error) {
	res := &powerxtypes.DeleteContractWayReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/contract-way/%v", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
