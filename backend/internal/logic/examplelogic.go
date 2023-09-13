package logic

import (
	"context"

	"PluginTemplate/internal/svc"
	"PluginTemplate/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExampleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExampleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExampleLogic {
	return &ExampleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExampleLogic) Example(req *types.ExampleRequest) (resp *types.ExampleReply, err error) {
	// todo: add your logic here and delete this line

	return
}
