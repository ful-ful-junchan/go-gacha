package repository

import (
	"context"
	"database/sql"
)

// Executor は *sql.DB と *sql.Tx の両方で満たされるインターフェース。
// トランザクション内・外のどちらからでも同じリポジトリ関数を呼べるようにするために使う。
type Executor interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}
