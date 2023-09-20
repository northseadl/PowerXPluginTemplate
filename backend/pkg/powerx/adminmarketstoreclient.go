package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type AdminMarketStore struct {
	*PowerX
}

func (c *AdminMarketStore) ListStoresPage(req *types.ListStoresPageRequest) (*types.ListStoresPageReply, error) {
	res := &types.ListStoresPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/market/stores/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminMarketStore) GetStore(req *types.GetStoreRequest) (*types.GetStoreReply, error) {
	res := &types.GetStoreReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/market/stores/%s", req.StoreId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminMarketStore) CreateStore(req *types.CreateStoreRequest) (*types.CreateStoreReply, error) {
	res := &types.CreateStoreReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/market/stores").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminMarketStore) PutStore(req *types.PutStoreRequest) (*types.PutStoreReply, error) {
	res := &types.PutStoreReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/market/stores/%s", req.StoreId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminMarketStore) DeleteStore(req *types.DeleteStoreRequest) (*types.DeleteStoreReply, error) {
	res := &types.DeleteStoreReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/market/stores/%s", req.StoreId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminMarketStore) AssignStoreToStoreManager(req *types.AssignStoreManagerRequest) (*types.AssignStoreManagerReply, error) {
	res := &types.AssignStoreManagerReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri(fmt.Sprintf("/api/v1/admin/market/stores/%s/actions/assign-to-store-categroy", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
