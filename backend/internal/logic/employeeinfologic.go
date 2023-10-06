package logic

import (
	"PluginTemplate/internal/svc"
	"PluginTemplate/internal/types"
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type EmployeeInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmployeeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmployeeInfoLogic {
	return &EmployeeInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmployeeInfoLogic) EmployeeInfo(req *types.EmployeeInfoRequest) (resp *types.EmployeeInfoReply, err error) {
	res, err := l.svcCtx.PowerX.AdminEmployee.GetEmployee(l.ctx, &powerxtypes.GetEmployeeRequest{
		Id: req.EmployeeId,
	})
	if err != nil {
		return
	}
	resp = &types.EmployeeInfoReply{
		EmployeeName: res.Name,
	}
	return
}
