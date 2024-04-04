// Code generated by ent, DO NOT EDIT.

package emaillog

import (
	"time"

	"entgo.io/ent/dialect/sql"
	uuid "github.com/gofrs/uuid/v5"
)

const (
	// Label holds the string label denoting the emaillog type in the database.
	Label = "email_log"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldTarget holds the string denoting the target field in the database.
	FieldTarget = "target"
	// FieldSubject holds the string denoting the subject field in the database.
	FieldSubject = "subject"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldSendStatus holds the string denoting the send_status field in the database.
	FieldSendStatus = "send_status"
	// FieldProvider holds the string denoting the provider field in the database.
	FieldProvider = "provider"
	// Table holds the table name of the emaillog in the database.
	Table = "mcms_email_logs"
)

// Columns holds all SQL columns for emaillog fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldTarget,
	FieldSubject,
	FieldContent,
	FieldSendStatus,
	FieldProvider,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the EmailLog queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByTarget orders the results by the target field.
func ByTarget(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTarget, opts...).ToFunc()
}

// BySubject orders the results by the subject field.
func BySubject(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubject, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
}

// BySendStatus orders the results by the send_status field.
func BySendStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSendStatus, opts...).ToFunc()
}

// ByProvider orders the results by the provider field.
func ByProvider(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProvider, opts...).ToFunc()
}
