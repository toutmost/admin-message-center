package emaillog

import (
	"context"
	"github.com/toutmost/admin-common/utils/pointy"
	"github.com/toutmost/admin-message-center/ent/emaillog"
	"github.com/toutmost/admin-message-center/ent/predicate"
	"github.com/toutmost/admin-message-center/internal/utils/dberrorhandler"

	"github.com/toutmost/admin-message-center/internal/svc"
	"github.com/toutmost/admin-message-center/types/mcms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmailLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEmailLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmailLogListLogic {
	return &GetEmailLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetEmailLogListLogic) GetEmailLogList(in *mcms.EmailLogListReq) (*mcms.EmailLogListResp, error) {
	var predicates []predicate.EmailLog
	if in.Target != nil {
		predicates = append(predicates, emaillog.TargetContains(*in.Target))
	}
	if in.Subject != nil {
		predicates = append(predicates, emaillog.SubjectContains(*in.Subject))
	}
	if in.Provider != nil {
		predicates = append(predicates, emaillog.ContentContains(*in.Provider))
	}
	if in.SendStatus != nil && *in.SendStatus != 0 {
		predicates = append(predicates, emaillog.SendStatusEQ(uint8(*in.SendStatus)))
	}
	result, err := l.svcCtx.DB.EmailLog.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &mcms.EmailLogListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &mcms.EmailLogInfo{
			Id:         pointy.GetPointer(v.ID.String()),
			CreatedAt:  pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:  pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Target:     &v.Target,
			Subject:    &v.Subject,
			Content:    &v.Content,
			SendStatus: pointy.GetPointer(uint32(v.SendStatus)),
		})
	}

	return resp, nil
}
