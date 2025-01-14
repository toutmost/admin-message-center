package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/toutmost/admin-common/orm/ent/mixins"
)

// EmailLog holds the schema definition for the EmailLog entity.
type EmailLog struct {
	ent.Schema
}

// Fields of the EmailLog.
func (EmailLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("target").Comment("The target email address | 目标邮箱地址").
			Annotations(entsql.WithComments(true)),
		field.String("subject").Comment("The subject | 发送的标题").
			Annotations(entsql.WithComments(true)),
		field.String("content").Comment("The content | 发送的内容").
			Annotations(entsql.WithComments(true)),
		field.Uint8("send_status").Comment("The send status, 0 unknown 1 success 2 failed | 发送的状态, 0 未知， 1 成功， 2 失败").
			Annotations(entsql.WithComments(true)),
		field.String("provider").Comment("The sms service provider | 短信服务提供商").
			Annotations(entsql.WithComments(true)),
	}
}

func (EmailLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
	}
}

// Edges of the EmailLog.
func (EmailLog) Edges() []ent.Edge {
	return nil
}

func (EmailLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "mcms_email_logs"},
	}
}
