package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type MpDictionary struct {
	*PowerX
}

func (c *MpDictionary) ListDictionaryPageTypes(req *types.ListDictionaryTypesPageRequest) (*types.ListDictionaryTypesPageReply, error) {
	res := &types.ListDictionaryTypesPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/web/dictionary/types/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpDictionary) GetDictionaryType(req *types.GetDictionaryTypeRequest) (*types.GetDictionaryTypeReply, error) {
	res := &types.GetDictionaryTypeReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/web/dictionary/types/%s", req.DictionaryType)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpDictionary) ListDictionaryItems(req *types.ListDictionaryItemsRequest) (*types.ListDictionaryItemsReply, error) {
	res := &types.ListDictionaryItemsReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/web/dictionary/items").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpDictionary) GetDictionaryItem(req *types.GetDictionaryItemRequest) (*types.GetDictionaryItemReply, error) {
	res := &types.GetDictionaryItemReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri(fmt.Sprintf("/api/v1/web/dictionary/items/%s/%s", req.DictionaryType, req.DictionaryItem)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
