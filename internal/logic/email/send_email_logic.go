package email

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/toutmost/admin-common/i18n"
	emailprovider2 "github.com/toutmost/admin-message-center/ent/emailprovider"
	"github.com/toutmost/admin-message-center/internal/config"
	"github.com/toutmost/admin-message-center/internal/utils/dberrorhandler"
	"github.com/toutmost/admin-message-center/internal/utils/email"
	"github.com/zeromicro/go-zero/core/errorx"
	"net/smtp"
	"strings"

	"github.com/toutmost/admin-message-center/internal/svc"
	"github.com/toutmost/admin-message-center/types/mcms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailLogic {
	return &SendEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendEmailLogic) SendEmail(in *mcms.EmailInfo) (*mcms.BaseUUIDResp, error) {
	// If the provider is nil, use default
	if in.Provider == nil {
		defaultProvider, err := l.svcCtx.DB.EmailProvider.Query().Where(emailprovider2.IsDefaultEQ(true)).First(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
		}
		in.Provider = &defaultProvider.Name
	}

	var client *smtp.Client
	var ok bool

	// client cache
	if client, ok = l.svcCtx.EmailClientGroup[*in.Provider]; !ok {
		providerData, err := l.svcCtx.DB.EmailProvider.Query().Where(emailprovider2.NameEQ(*in.Provider)).First(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
		}

		emailProviderConfig := &config.EmailConf{
			AuthType:  email.ConvertAuthTypeToString(providerData.AuthType),
			EmailAddr: providerData.EmailAddr,
			Password:  providerData.Password,
			HostName:  providerData.HostName,
			Port:      int(providerData.Port),
			TLS:       providerData.TLS,
		}
		l.svcCtx.EmailClientGroup[*in.Provider] = emailProviderConfig.NewClient()
		l.svcCtx.EmailAddrGroup[*in.Provider] = emailProviderConfig.EmailAddr

		client = l.svcCtx.EmailClientGroup[*in.Provider]
	}

	// error handler
	emailErrHandler := func(err error) (*mcms.BaseUUIDResp, error) {
		l.Logger.Errorw("failed to send email", logx.Field("detail", err.Error()), logx.Field("data", in))

		dberr := l.svcCtx.DB.EmailLog.Create().
			SetTarget(strings.Join(in.Target, ",")).
			SetContent(in.Content).
			SetSubject(in.Subject).
			SetSendStatus(2).
			SetProvider(*in.Provider).
			Exec(context.Background())

		if dberr != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, dberr, in)
		}

		return nil, errorx.NewInternalError(i18n.Failed)
	}

	fromEmailAddress := l.svcCtx.EmailAddrGroup[*in.Provider]

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = fromEmailAddress
	headers["To"] = strings.Join(in.Target, ",")
	headers["Subject"] = in.Subject
	headers["Content-Type"] = "text/html; charset=UTF-8"

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + in.Content

	err := client.Mail(fromEmailAddress)
	if err != nil {
		l.Logger.Errorw("failed to set the from address in email", logx.Field("detail", err), logx.Field("data", in))
		return emailErrHandler(errors.Wrap(err, "failed to set the from address in email"))
	}

	for _, v := range in.Target {
		err := client.Rcpt(v)
		if err != nil {
			l.Logger.Errorw("failed to set the to address in email", logx.Field("detail", err), logx.Field("data", in))
			return emailErrHandler(errors.Wrap(err, "failed to set the from address in email"))
		}
	}

	w, err := client.Data()
	if err != nil {
		l.Logger.Errorw("failed to create the writer for email", logx.Field("detail", err), logx.Field("data", in))
		return emailErrHandler(errors.Wrap(err, "failed to create the writer for email"))
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		l.Logger.Errorw("failed to write the data to writer for email", logx.Field("detail", err), logx.Field("data", in))
		return emailErrHandler(errors.Wrap(err, "failed to write the data to writer for email"))
	}

	err = w.Close()
	if err != nil {
		l.Logger.Errorw("failed to close the writer for email", logx.Field("detail", err), logx.Field("data", in))
		return emailErrHandler(errors.Wrap(err, "failed to close the writer for email"))
	}

	logData, dberr := l.svcCtx.DB.EmailLog.Create().
		SetTarget(strings.Join(in.Target, ",")).
		SetContent(in.Content).
		SetSubject(in.Subject).
		SetSendStatus(1).
		SetProvider(*in.Provider).
		Save(context.Background())

	if dberr != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, dberr, in)
	}

	return &mcms.BaseUUIDResp{Msg: i18n.Success, Id: logData.ID.String()}, nil
}
