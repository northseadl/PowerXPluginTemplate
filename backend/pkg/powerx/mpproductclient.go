package powerx

import (
	"PluginTemplate/pkg/powerx/types"

	"net/http"
)

type MpProduct struct {
	*PowerX
}

func (c *MpProduct) ListProductCategoryTree(req *types.ListProductCategoryTreeRequest) (*types.ListProductCategoryTreeReply, error) {
	res := &types.ListProductCategoryTreeReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/mp/product/product-category-tree").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpProduct) ListProductCategories(req *types.ListProductCategoriesRequest) (*types.ListProductCategoriesReply, error) {
	res := &types.ListProductCategoriesReply{}
	err := c.H.Df().Method(http.MethodGet).
		Uri("/api/v1/mp/product/product-categories").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
