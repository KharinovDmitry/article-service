package adapter

import (
	"context"
)

type DBAdapter interface {
	Execute(requestCtx context.Context, sql string, args ...interface{}) error
	ExecuteAndGet(requestCtx context.Context, destination interface{}, sql string, args ...interface{}) error
	Query(requestCtx context.Context, destination interface{}, query string, args ...interface{}) error
}
