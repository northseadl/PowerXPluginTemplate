package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"
	"fmt"
	"net/http"
)

type AdminMarketStore struct {
	*PowerX
}

func (c *AdminMarketStore) ListStoresPage(ctx context.Context, req *powerxtypes.ListStoresPageRequest) (*powerxtypes.ListStoresPageReply, error) {
	res := &powerxtypes.ListStoresPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/admin/market/stores/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminMarketStore) GetStore(ctx context.Context, req *powerxtypes.GetStoreRequest) (*powerxtypes.GetStoreReply, error) {
	res := &powerxtypes.GetStoreReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/market/stores/%v", req.StoreId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminMarketStore) CreateStore(ctx context.Context, req *powerxtypes.CreateStoreRequest) (*powerxtypes.CreateStoreReply, error) {
	res := &powerxtypes.CreateStoreReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/admin/market/stores").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminMarketStore) PutStore(ctx context.Context, req *powerxtypes.PutStoreRequest) (*powerxtypes.PutStoreReply, error) {
	res := &powerxtypes.PutStoreReply{}
	err := c.H.Df().Method(http.MethodPut).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/market/stores/%v", req.StoreId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminMarketStore) DeleteStore(ctx context.Context, req *powerxtypes.DeleteStoreRequest) (*powerxtypes.DeleteStoreReply, error) {
	res := &powerxtypes.DeleteStoreReply{}
	err := c.H.Df().Method(http.MethodDelete).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/market/stores/%v", req.StoreId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminMarketStore) AssignStoreToStoreManager(ctx context.Context, req *powerxtypes.AssignStoreManagerRequest) (*powerxtypes.AssignStoreManagerReply, error) {
	res := &powerxtypes.AssignStoreManagerReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/market/stores/%v/actions/assign-to-store-categroy", req.Id)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
