package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/toutmost/admin-message-center/ent"
	"github.com/toutmost/admin-message-center/internal/config"
	"github.com/toutmost/admin-message-center/internal/utils/smssdk"
	"github.com/zeromicro/go-zero/core/logx"
	"net/smtp"
)

type ServiceContext struct {
	Config           config.Config
	DB               *ent.Client
	Redis            redis.UniversalClient
	EmailAuth        *smtp.Auth
	SmsGroup         *smssdk.SmsGroup
	EmailClientGroup map[string]*smtp.Client
	EmailAddrGroup   map[string]string
}

func NewServiceContext(c config.Config) *ServiceContext {

	db := ent.NewClient(
		ent.Log(logx.Info), // logger
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
		ent.Debug(), // debug mode
	)

	return &ServiceContext{
		Config:           c,
		DB:               db,
		Redis:            c.RedisConf.MustNewUniversalRedis(),
		SmsGroup:         &smssdk.SmsGroup{},
		EmailAddrGroup:   map[string]string{},
		EmailClientGroup: map[string]*smtp.Client{},
	}
}
