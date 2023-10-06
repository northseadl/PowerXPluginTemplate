package powerx

import (
	powerxtypes "PluginTemplate/pkg/powerx/powerxtypes"
	"net/http"
)

type AdminTradeWarehouse struct {
	*PowerX
}

func (c *AdminTradeWarehouse) ListWarehouses(req *powerxtypes.ListWarehousesRequest) (*powerxtypes.ListWarehousesResponse, error) {
	res := &powerxtypes.ListWarehousesResponse{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/trade/warehouses").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeWarehouse) GetWarehouse(req *powerxtypes.GetWarehouseRequest) (*powerxtypes.GetWarehouseResponse, error) {
	res := &powerxtypes.GetWarehouseResponse{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/trade/warehouses/:id").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeWarehouse) CreateWarehouse(req *powerxtypes.CreateWarehouseRequest) (*powerxtypes.CreateWarehouseResponse, error) {
	res := &powerxtypes.CreateWarehouseResponse{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/trade/warehouses").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeWarehouse) UpdateWarehouse(req *powerxtypes.UpdateWarehouseRequest) (*powerxtypes.UpdateWarehouseResponse, error) {
	res := &powerxtypes.UpdateWarehouseResponse{}
	err := c.H.Df().Method(http.MethodPut).
		Uri("/api/v1/admin/trade/warehouses/:id").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeWarehouse) PatchWarehouse(req *powerxtypes.PatchWarehouseRequest) (*powerxtypes.PatchWarehouseResponse, error) {
	res := &powerxtypes.PatchWarehouseResponse{}
	err := c.H.Df().Method(http.MethodPatch).
		Uri("/api/v1/admin/trade/warehouses/:id").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeWarehouse) DeleteWarehouse(req *powerxtypes.DeleteWarehouseRequest) (*powerxtypes.DeleteWarehouseResponse, error) {
	res := &powerxtypes.DeleteWarehouseResponse{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri("/api/v1/admin/trade/warehouses/:id").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
