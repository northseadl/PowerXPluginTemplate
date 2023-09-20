package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type AdminDictionary struct {
	*PowerX
}

func (c *AdminDictionary) ListDictionaryPageTypes(req *types.ListDictionaryTypesPageRequest) (*types.ListDictionaryTypesPageReply, error) {
	res := &types.ListDictionaryTypesPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/dictionary/types/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDictionary) ListDictionaryTypes(req *types.ListDictionaryTypesPageRequest) (*types.ListDictionaryTypesPageReply, error) {
	res := &types.ListDictionaryTypesPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/dictionary/types").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDictionary) GetDictionaryType(req *types.GetDictionaryTypeRequest) (*types.GetDictionaryTypeReply, error) {
	res := &types.GetDictionaryTypeReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/dictionary/types/%s", req.DictionaryType)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDictionary) CreateDictionaryType(req *types.CreateDictionaryTypeRequest) (*types.CreateDictionaryTypeReply, error) {
	res := &types.CreateDictionaryTypeReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/dictionary/types").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDictionary) UpdateDictionaryType(req *types.UpdateDictionaryTypeRequest) (*types.UpdateDictionaryTypeReply, error) {
	res := &types.UpdateDictionaryTypeReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/dictionary/types/%s", req.Type)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDictionary) DeleteDictionaryType(req *types.DeleteDictionaryTypeRequest) (*types.DeleteDictionaryTypeReply, error) {
	res := &types.DeleteDictionaryTypeReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/dictionary/types/%s", req.Type)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDictionary) ListDictionaryItems(req *types.ListDictionaryItemsRequest) (*types.ListDictionaryItemsReply, error) {
	res := &types.ListDictionaryItemsReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/dictionary/items").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDictionary) GetDictionaryItem(req *types.GetDictionaryItemRequest) (*types.GetDictionaryItemReply, error) {
	res := &types.GetDictionaryItemReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/dictionary/items/%s/%s", req.DictionaryType, req.DictionaryItem)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDictionary) CreateDictionaryItem(req *types.CreateDictionaryItemRequest) (*types.CreateDictionaryItemReply, error) {
	res := &types.CreateDictionaryItemReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/dictionary/items").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDictionary) UpdateDictionaryItem(req *types.UpdateDictionaryItemRequest) (*types.UpdateDictionaryItemReply, error) {
	res := &types.UpdateDictionaryItemReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/dictionary/items/%s/%s", req.Key, req.Type)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDictionary) DeleteDictionaryItem(req *types.DeleteDictionaryItemRequest) (*types.DeleteDictionaryItemReply, error) {
	res := &types.DeleteDictionaryItemReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/dictionary/items/%s/%s", req.Key, req.Type)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
