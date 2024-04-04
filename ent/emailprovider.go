// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/toutmost/admin-message-center/ent/emailprovider"
)

// EmailProvider is the model entity for the EmailProvider schema.
type EmailProvider struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// The email provider name | 电子邮件服务的提供商
	Name string `json:"name,omitempty"`
	// The auth type, supported plain, CRAMMD5 | 鉴权类型, 支持 plain, CRAMMD5
	AuthType uint8 `json:"auth_type,omitempty"`
	// The email address | 邮箱地址
	EmailAddr string `json:"email_addr,omitempty"`
	// The email's password | 电子邮件的密码
	Password string `json:"password,omitempty"`
	// The host name is the email service's host address | 电子邮箱服务的服务器地址
	HostName string `json:"host_name,omitempty"`
	// The identify info, for CRAMMD5 | 身份信息, 支持 CRAMMD5
	Identify string `json:"identify,omitempty"`
	// The secret, for CRAMMD5 | 邮箱密钥, 用于 CRAMMD5
	Secret string `json:"secret,omitempty"`
	// The port of the host | 服务器端口
	Port uint32 `json:"port,omitempty"`
	// Whether to use TLS | 是否采用 tls 加密
	TLS bool `json:"tls,omitempty"`
	// Is it the default provider | 是否为默认提供商
	IsDefault    bool `json:"is_default,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EmailProvider) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case emailprovider.FieldTLS, emailprovider.FieldIsDefault:
			values[i] = new(sql.NullBool)
		case emailprovider.FieldID, emailprovider.FieldAuthType, emailprovider.FieldPort:
			values[i] = new(sql.NullInt64)
		case emailprovider.FieldName, emailprovider.FieldEmailAddr, emailprovider.FieldPassword, emailprovider.FieldHostName, emailprovider.FieldIdentify, emailprovider.FieldSecret:
			values[i] = new(sql.NullString)
		case emailprovider.FieldCreatedAt, emailprovider.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EmailProvider fields.
func (ep *EmailProvider) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case emailprovider.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ep.ID = uint64(value.Int64)
		case emailprovider.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ep.CreatedAt = value.Time
			}
		case emailprovider.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ep.UpdatedAt = value.Time
			}
		case emailprovider.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ep.Name = value.String
			}
		case emailprovider.FieldAuthType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field auth_type", values[i])
			} else if value.Valid {
				ep.AuthType = uint8(value.Int64)
			}
		case emailprovider.FieldEmailAddr:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email_addr", values[i])
			} else if value.Valid {
				ep.EmailAddr = value.String
			}
		case emailprovider.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				ep.Password = value.String
			}
		case emailprovider.FieldHostName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field host_name", values[i])
			} else if value.Valid {
				ep.HostName = value.String
			}
		case emailprovider.FieldIdentify:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field identify", values[i])
			} else if value.Valid {
				ep.Identify = value.String
			}
		case emailprovider.FieldSecret:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field secret", values[i])
			} else if value.Valid {
				ep.Secret = value.String
			}
		case emailprovider.FieldPort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field port", values[i])
			} else if value.Valid {
				ep.Port = uint32(value.Int64)
			}
		case emailprovider.FieldTLS:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field tls", values[i])
			} else if value.Valid {
				ep.TLS = value.Bool
			}
		case emailprovider.FieldIsDefault:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_default", values[i])
			} else if value.Valid {
				ep.IsDefault = value.Bool
			}
		default:
			ep.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the EmailProvider.
// This includes values selected through modifiers, order, etc.
func (ep *EmailProvider) Value(name string) (ent.Value, error) {
	return ep.selectValues.Get(name)
}

// Update returns a builder for updating this EmailProvider.
// Note that you need to call EmailProvider.Unwrap() before calling this method if this EmailProvider
// was returned from a transaction, and the transaction was committed or rolled back.
func (ep *EmailProvider) Update() *EmailProviderUpdateOne {
	return NewEmailProviderClient(ep.config).UpdateOne(ep)
}

// Unwrap unwraps the EmailProvider entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ep *EmailProvider) Unwrap() *EmailProvider {
	_tx, ok := ep.config.driver.(*txDriver)
	if !ok {
		panic("ent: EmailProvider is not a transactional entity")
	}
	ep.config.driver = _tx.drv
	return ep
}

// String implements the fmt.Stringer.
func (ep *EmailProvider) String() string {
	var builder strings.Builder
	builder.WriteString("EmailProvider(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ep.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ep.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ep.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ep.Name)
	builder.WriteString(", ")
	builder.WriteString("auth_type=")
	builder.WriteString(fmt.Sprintf("%v", ep.AuthType))
	builder.WriteString(", ")
	builder.WriteString("email_addr=")
	builder.WriteString(ep.EmailAddr)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(ep.Password)
	builder.WriteString(", ")
	builder.WriteString("host_name=")
	builder.WriteString(ep.HostName)
	builder.WriteString(", ")
	builder.WriteString("identify=")
	builder.WriteString(ep.Identify)
	builder.WriteString(", ")
	builder.WriteString("secret=")
	builder.WriteString(ep.Secret)
	builder.WriteString(", ")
	builder.WriteString("port=")
	builder.WriteString(fmt.Sprintf("%v", ep.Port))
	builder.WriteString(", ")
	builder.WriteString("tls=")
	builder.WriteString(fmt.Sprintf("%v", ep.TLS))
	builder.WriteString(", ")
	builder.WriteString("is_default=")
	builder.WriteString(fmt.Sprintf("%v", ep.IsDefault))
	builder.WriteByte(')')
	return builder.String()
}

// EmailProviders is a parsable slice of EmailProvider.
type EmailProviders []*EmailProvider
