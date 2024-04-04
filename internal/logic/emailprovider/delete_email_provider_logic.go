package emailprovider

import (
	"context"
	"github.com/toutmost/admin-common/i18n"
	"github.com/toutmost/admin-message-center/ent/emailprovider"
	"github.com/toutmost/admin-message-center/internal/utils/dberrorhandler"
	"net/smtp"

	"github.com/toutmost/admin-message-center/internal/svc"
	"github.com/toutmost/admin-message-center/types/mcms"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteEmailProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteEmailProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteEmailProviderLogic {
	return &DeleteEmailProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteEmailProviderLogic) DeleteEmailProvider(in *mcms.IDsReq) (*mcms.BaseResp, error) {
	_, err := l.svcCtx.DB.EmailProvider.Delete().Where(emailprovider.IDIn(in.Ids...)).Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	l.svcCtx.EmailAddrGroup = map[string]string{}
	l.svcCtx.EmailClientGroup = map[string]*smtp.Client{}

	return &mcms.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
