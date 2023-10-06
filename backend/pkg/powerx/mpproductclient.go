package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"

	"net/http"
)

type MpProduct struct {
	*PowerX
}

func (c *MpProduct) ListProductCategoryTree(ctx context.Context, req *powerxtypes.ListProductCategoryTreeRequest) (*powerxtypes.ListProductCategoryTreeReply, error) {
	res := &powerxtypes.ListProductCategoryTreeReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/mp/product/product-category-tree").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MpProduct) ListProductCategories(ctx context.Context, req *powerxtypes.ListProductCategoriesRequest) (*powerxtypes.ListProductCategoriesReply, error) {
	res := &powerxtypes.ListProductCategoriesReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/mp/product/product-categories").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
