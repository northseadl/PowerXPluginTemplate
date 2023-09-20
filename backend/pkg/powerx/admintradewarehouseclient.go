package powerx

import (
	"PluginTemplate/pkg/powerx/types"

	"net/http"
)

type AdminTradeWarehouse struct {
	*PowerX
}

func (c *AdminTradeWarehouse) ListWarehouses(req *types.ListWarehousesRequest) (*types.ListWarehousesResponse, error) {
	res := &types.ListWarehousesResponse{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/trade/warehouses").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeWarehouse) GetWarehouse(req *types.GetWarehouseRequest) (*types.GetWarehouseResponse, error) {
	res := &types.GetWarehouseResponse{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/trade/warehouses/:id").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeWarehouse) CreateWarehouse(req *types.CreateWarehouseRequest) (*types.CreateWarehouseResponse, error) {
	res := &types.CreateWarehouseResponse{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/trade/warehouses").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeWarehouse) UpdateWarehouse(req *types.UpdateWarehouseRequest) (*types.UpdateWarehouseResponse, error) {
	res := &types.UpdateWarehouseResponse{}
	err := c.H.Df().Method(http.MethodPut).
		Uri("/api/v1/admin/trade/warehouses/:id").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeWarehouse) PatchWarehouse(req *types.PatchWarehouseRequest) (*types.PatchWarehouseResponse, error) {
	res := &types.PatchWarehouseResponse{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri("/api/v1/admin/trade/warehouses/:id").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeWarehouse) DeleteWarehouse(req *types.DeleteWarehouseRequest) (*types.DeleteWarehouseResponse, error) {
	res := &types.DeleteWarehouseResponse{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri("/api/v1/admin/trade/warehouses/:id").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
