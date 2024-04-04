package emailprovider

import (
	"context"
	"github.com/toutmost/admin-common/utils/pointy"
	"github.com/toutmost/admin-message-center/internal/utils/dberrorhandler"

	"github.com/toutmost/admin-message-center/internal/svc"
	"github.com/toutmost/admin-message-center/types/mcms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmailProviderByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEmailProviderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmailProviderByIdLogic {
	return &GetEmailProviderByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetEmailProviderByIdLogic) GetEmailProviderById(in *mcms.IDReq) (*mcms.EmailProviderInfo, error) {
	result, err := l.svcCtx.DB.EmailProvider.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &mcms.EmailProviderInfo{
		Id:        &result.ID,
		CreatedAt: pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt: pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		Name:      &result.Name,
		AuthType:  pointy.GetPointer(uint32(result.AuthType)),
		EmailAddr: &result.EmailAddr,
		Password:  &result.Password,
		HostName:  &result.HostName,
		Identify:  &result.Identify,
		Secret:    &result.Secret,
		Port:      &result.Port,
		Tls:       &result.TLS,
		IsDefault: &result.IsDefault,
	}, nil
}
