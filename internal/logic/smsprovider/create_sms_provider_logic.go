package smsprovider

import (
	"context"

	"github.com/toutmost/admin-message-center/internal/svc"
	"github.com/toutmost/admin-message-center/types/mcms"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSmsProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSmsProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSmsProviderLogic {
	return &CreateSmsProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SmsProvider management
func (l *CreateSmsProviderLogic) CreateSmsProvider(in *mcms.SmsProviderInfo) (*mcms.BaseIDResp, error) {
	// todo: add your logic here and delete this line

	return &mcms.BaseIDResp{}, nil
}
