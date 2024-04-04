package emaillog

import (
	"context"
	"github.com/toutmost/admin-common/i18n"
	"github.com/toutmost/admin-common/utils/pointy"
	"github.com/toutmost/admin-common/utils/uuidx"
	"github.com/toutmost/admin-message-center/internal/utils/dberrorhandler"

	"github.com/toutmost/admin-message-center/internal/svc"
	"github.com/toutmost/admin-message-center/types/mcms"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateEmailLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateEmailLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEmailLogLogic {
	return &UpdateEmailLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateEmailLogLogic) UpdateEmailLog(in *mcms.EmailLogInfo) (*mcms.BaseResp, error) {
	query := l.svcCtx.DB.EmailLog.UpdateOneID(uuidx.ParseUUIDString(*in.Id)).
		SetNotNilTarget(in.Target).
		SetNotNilSubject(in.Subject).
		SetNotNilContent(in.Content)

	if in.SendStatus != nil {
		query.SetNotNilSendStatus(pointy.GetPointer(uint8(*in.SendStatus)))
	}

	err := query.Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &mcms.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
