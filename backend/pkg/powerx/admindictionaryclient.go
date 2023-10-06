package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"
	"fmt"
	"net/http"
)

type AdminDictionary struct {
	*PowerX
}

func (c *AdminDictionary) ListDictionaryPageTypes(ctx context.Context, req *powerxtypes.ListDictionaryTypesPageRequest) (*powerxtypes.ListDictionaryTypesPageReply, error) {
	res := &powerxtypes.ListDictionaryTypesPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/admin/dictionary/types/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDictionary) ListDictionaryTypes(ctx context.Context, req *powerxtypes.ListDictionaryTypesPageRequest) (*powerxtypes.ListDictionaryTypesPageReply, error) {
	res := &powerxtypes.ListDictionaryTypesPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/admin/dictionary/types").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDictionary) GetDictionaryType(ctx context.Context, req *powerxtypes.GetDictionaryTypeRequest) (*powerxtypes.GetDictionaryTypeReply, error) {
	res := &powerxtypes.GetDictionaryTypeReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/dictionary/types/%v", req.DictionaryType)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDictionary) CreateDictionaryType(ctx context.Context, req *powerxtypes.CreateDictionaryTypeRequest) (*powerxtypes.CreateDictionaryTypeReply, error) {
	res := &powerxtypes.CreateDictionaryTypeReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/admin/dictionary/types").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDictionary) UpdateDictionaryType(ctx context.Context, req *powerxtypes.UpdateDictionaryTypeRequest) (*powerxtypes.UpdateDictionaryTypeReply, error) {
	res := &powerxtypes.UpdateDictionaryTypeReply{}
	err := c.H.Df().Method(http.MethodPut).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/dictionary/types/%v", req.Type)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDictionary) DeleteDictionaryType(ctx context.Context, req *powerxtypes.DeleteDictionaryTypeRequest) (*powerxtypes.DeleteDictionaryTypeReply, error) {
	res := &powerxtypes.DeleteDictionaryTypeReply{}
	err := c.H.Df().Method(http.MethodDelete).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/dictionary/types/%v", req.Type)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDictionary) ListDictionaryItems(ctx context.Context, req *powerxtypes.ListDictionaryItemsRequest) (*powerxtypes.ListDictionaryItemsReply, error) {
	res := &powerxtypes.ListDictionaryItemsReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/admin/dictionary/items").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDictionary) GetDictionaryItem(ctx context.Context, req *powerxtypes.GetDictionaryItemRequest) (*powerxtypes.GetDictionaryItemReply, error) {
	res := &powerxtypes.GetDictionaryItemReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/dictionary/items/%v/%v", req.DictionaryType, req.DictionaryItem)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDictionary) CreateDictionaryItem(ctx context.Context, req *powerxtypes.CreateDictionaryItemRequest) (*powerxtypes.CreateDictionaryItemReply, error) {
	res := &powerxtypes.CreateDictionaryItemReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/admin/dictionary/items").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDictionary) UpdateDictionaryItem(ctx context.Context, req *powerxtypes.UpdateDictionaryItemRequest) (*powerxtypes.UpdateDictionaryItemReply, error) {
	res := &powerxtypes.UpdateDictionaryItemReply{}
	err := c.H.Df().Method(http.MethodPut).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/dictionary/items/%v/%v", req.Key, req.Type)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminDictionary) DeleteDictionaryItem(ctx context.Context, req *powerxtypes.DeleteDictionaryItemRequest) (*powerxtypes.DeleteDictionaryItemReply, error) {
	res := &powerxtypes.DeleteDictionaryItemReply{}
	err := c.H.Df().Method(http.MethodDelete).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/dictionary/items/%v/%v", req.Key, req.Type)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
