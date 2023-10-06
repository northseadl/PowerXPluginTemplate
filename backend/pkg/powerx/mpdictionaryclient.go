package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"
	"fmt"
	"net/http"
)

type MpDictionary struct {
	*PowerX
}

func (c *MpDictionary) ListDictionaryPageTypes(ctx context.Context, req *powerxtypes.ListDictionaryTypesPageRequest) (*powerxtypes.ListDictionaryTypesPageReply, error) {
	res := &powerxtypes.ListDictionaryTypesPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/web/dictionary/types/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpDictionary) GetDictionaryType(ctx context.Context, req *powerxtypes.GetDictionaryTypeRequest) (*powerxtypes.GetDictionaryTypeReply, error) {
	res := &powerxtypes.GetDictionaryTypeReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/web/dictionary/types/%v", req.DictionaryType)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpDictionary) ListDictionaryItems(ctx context.Context, req *powerxtypes.ListDictionaryItemsRequest) (*powerxtypes.ListDictionaryItemsReply, error) {
	res := &powerxtypes.ListDictionaryItemsReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/web/dictionary/items").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpDictionary) GetDictionaryItem(ctx context.Context, req *powerxtypes.GetDictionaryItemRequest) (*powerxtypes.GetDictionaryItemReply, error) {
	res := &powerxtypes.GetDictionaryItemReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/web/dictionary/items/%v/%v", req.DictionaryType, req.DictionaryItem)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
