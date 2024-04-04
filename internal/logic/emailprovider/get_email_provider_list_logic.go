package emailprovider

import (
	"context"
	"github.com/toutmost/admin-common/utils/pointy"
	"github.com/toutmost/admin-message-center/ent/emailprovider"
	"github.com/toutmost/admin-message-center/ent/predicate"
	"github.com/toutmost/admin-message-center/internal/utils/dberrorhandler"

	"github.com/toutmost/admin-message-center/internal/svc"
	"github.com/toutmost/admin-message-center/types/mcms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmailProviderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEmailProviderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmailProviderListLogic {
	return &GetEmailProviderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetEmailProviderListLogic) GetEmailProviderList(in *mcms.EmailProviderListReq) (*mcms.EmailProviderListResp, error) {
	var predicates []predicate.EmailProvider
	if in.Name != nil {
		predicates = append(predicates, emailprovider.NameContains(*in.Name))
	}
	if in.EmailAddr != nil {
		predicates = append(predicates, emailprovider.EmailAddrContains(*in.EmailAddr))
	}
	result, err := l.svcCtx.DB.EmailProvider.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &mcms.EmailProviderListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &mcms.EmailProviderInfo{
			Id:        &v.ID,
			CreatedAt: pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt: pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Name:      &v.Name,
			AuthType:  pointy.GetPointer(uint32(v.AuthType)),
			EmailAddr: &v.EmailAddr,
			Password:  &v.Password,
			HostName:  &v.HostName,
			Identify:  &v.Identify,
			Secret:    &v.Secret,
			Port:      &v.Port,
			Tls:       &v.TLS,
			IsDefault: &v.IsDefault,
		})
	}

	return resp, nil
}
