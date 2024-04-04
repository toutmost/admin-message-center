package smslog

import (
	"context"

	"github.com/toutmost/admin-message-center/internal/svc"
	"github.com/toutmost/admin-message-center/types/mcms"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSmsLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSmsLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSmsLogLogic {
	return &CreateSmsLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SmsLog management
func (l *CreateSmsLogLogic) CreateSmsLog(in *mcms.SmsLogInfo) (*mcms.BaseUUIDResp, error) {
	// todo: add your logic here and delete this line

	return &mcms.BaseUUIDResp{}, nil
}
