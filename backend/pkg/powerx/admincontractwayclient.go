package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type AdminContractway struct {
	*PowerX
}

func (c *AdminContractway) GetContractWayGroupTree(req *types.GetContractWayGroupTreeRequest) (*types.GetContractWayGroupTreeReply, error) {
	res := &types.GetContractWayGroupTreeReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/contract-way/group-tree").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminContractway) GetContractWayGroupList(req *types.GetContractWayGroupListRequest) (*types.GetContractWayGroupListReply, error) {
	res := &types.GetContractWayGroupListReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/contract-way/groups").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminContractway) GetContractWays(req *types.GetContractWaysRequest) (*types.GetContractWaysReply, error) {
	res := &types.GetContractWaysReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/contract-way").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminContractway) CreateContractWay(req *types.CreateContractWayRequest) (*types.CreateContractWayReply, error) {
	res := &types.CreateContractWayReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/contract-way").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminContractway) UpdateContractWay(req *types.UpdateContractWayRequest) (*types.UpdateContractWayReply, error) {
	res := &types.UpdateContractWayReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/contract-way/%s", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminContractway) DeleteContractWay(req *types.DeleteContractWayRequest) (*types.DeleteContractWayReply, error) {
	res := &types.DeleteContractWayReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/contract-way/%s", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
