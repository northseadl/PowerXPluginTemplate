package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"

	"net/http"
)

type AdminTradeWarehouse struct {
	*PowerX
}

func (c *AdminTradeWarehouse) ListWarehouses(ctx context.Context, req *powerxtypes.ListWarehousesRequest) (*powerxtypes.ListWarehousesResponse, error) {
	res := &powerxtypes.ListWarehousesResponse{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/admin/trade/warehouses").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeWarehouse) GetWarehouse(ctx context.Context, req *powerxtypes.GetWarehouseRequest) (*powerxtypes.GetWarehouseResponse, error) {
	res := &powerxtypes.GetWarehouseResponse{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/admin/trade/warehouses/:id").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeWarehouse) CreateWarehouse(ctx context.Context, req *powerxtypes.CreateWarehouseRequest) (*powerxtypes.CreateWarehouseResponse, error) {
	res := &powerxtypes.CreateWarehouseResponse{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/admin/trade/warehouses").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeWarehouse) UpdateWarehouse(ctx context.Context, req *powerxtypes.UpdateWarehouseRequest) (*powerxtypes.UpdateWarehouseResponse, error) {
	res := &powerxtypes.UpdateWarehouseResponse{}
	err := c.H.Df().Method(http.MethodPut).
		WithContext(ctx).
		Uri("/api/v1/admin/trade/warehouses/:id").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeWarehouse) PatchWarehouse(ctx context.Context, req *powerxtypes.PatchWarehouseRequest) (*powerxtypes.PatchWarehouseResponse, error) {
	res := &powerxtypes.PatchWarehouseResponse{}
	err := c.H.Df().Method(http.MethodPatch).
		WithContext(ctx).
		Uri("/api/v1/admin/trade/warehouses/:id").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeWarehouse) DeleteWarehouse(ctx context.Context, req *powerxtypes.DeleteWarehouseRequest) (*powerxtypes.DeleteWarehouseResponse, error) {
	res := &powerxtypes.DeleteWarehouseResponse{}
	err := c.H.Df().Method(http.MethodDelete).
		WithContext(ctx).
		Uri("/api/v1/admin/trade/warehouses/:id").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
