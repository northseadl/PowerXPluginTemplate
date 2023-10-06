package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"fmt"
	"net/http"
)

type AdminDictionary struct {
	*PowerX
}

func (c *AdminDictionary) ListDictionaryPageTypes(req *powerxtypes.ListDictionaryTypesPageRequest) (*powerxtypes.ListDictionaryTypesPageReply, error) {
	res := &powerxtypes.ListDictionaryTypesPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/dictionary/types/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDictionary) ListDictionaryTypes(req *powerxtypes.ListDictionaryTypesPageRequest) (*powerxtypes.ListDictionaryTypesPageReply, error) {
	res := &powerxtypes.ListDictionaryTypesPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/dictionary/types").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDictionary) GetDictionaryType(req *powerxtypes.GetDictionaryTypeRequest) (*powerxtypes.GetDictionaryTypeReply, error) {
	res := &powerxtypes.GetDictionaryTypeReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/dictionary/types/%v", req.DictionaryType)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDictionary) CreateDictionaryType(req *powerxtypes.CreateDictionaryTypeRequest) (*powerxtypes.CreateDictionaryTypeReply, error) {
	res := &powerxtypes.CreateDictionaryTypeReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/dictionary/types").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDictionary) UpdateDictionaryType(req *powerxtypes.UpdateDictionaryTypeRequest) (*powerxtypes.UpdateDictionaryTypeReply, error) {
	res := &powerxtypes.UpdateDictionaryTypeReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/dictionary/types/%v", req.Type)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDictionary) DeleteDictionaryType(req *powerxtypes.DeleteDictionaryTypeRequest) (*powerxtypes.DeleteDictionaryTypeReply, error) {
	res := &powerxtypes.DeleteDictionaryTypeReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/dictionary/types/%v", req.Type)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDictionary) ListDictionaryItems(req *powerxtypes.ListDictionaryItemsRequest) (*powerxtypes.ListDictionaryItemsReply, error) {
	res := &powerxtypes.ListDictionaryItemsReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/admin/dictionary/items").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDictionary) GetDictionaryItem(req *powerxtypes.GetDictionaryItemRequest) (*powerxtypes.GetDictionaryItemReply, error) {
	res := &powerxtypes.GetDictionaryItemReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/admin/dictionary/items/%v/%v", req.DictionaryType, req.DictionaryItem)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDictionary) CreateDictionaryItem(req *powerxtypes.CreateDictionaryItemRequest) (*powerxtypes.CreateDictionaryItemReply, error) {
	res := &powerxtypes.CreateDictionaryItemReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/dictionary/items").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDictionary) UpdateDictionaryItem(req *powerxtypes.UpdateDictionaryItemRequest) (*powerxtypes.UpdateDictionaryItemReply, error) {
	res := &powerxtypes.UpdateDictionaryItemReply{}
	err := c.H.Df().Method(http.MethodPut).
		Uri(fmt.Sprintf("/api/v1/admin/dictionary/items/%v/%v", req.Key, req.Type)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDictionary) DeleteDictionaryItem(req *powerxtypes.DeleteDictionaryItemRequest) (*powerxtypes.DeleteDictionaryItemReply, error) {
	res := &powerxtypes.DeleteDictionaryItemReply{}
	err := c.H.Df().Method(http.MethodDelete).
		Uri(fmt.Sprintf("/api/v1/admin/dictionary/items/%v/%v", req.Key, req.Type)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
