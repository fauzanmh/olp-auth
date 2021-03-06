package mysql

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.checkUserStmt, err = db.PrepareContext(ctx, checkUser); err != nil {
		return nil, fmt.Errorf("error preparing query CheckUser: %w", err)
	}
	if q.checkUsernameStmt, err = db.PrepareContext(ctx, checkUsername); err != nil {
		return nil, fmt.Errorf("error preparing query CheckUsername: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.deleteUserStmt, err = db.PrepareContext(ctx, deleteUser); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUser: %w", err)
	}
	if q.getAdminByUsernameStmt, err = db.PrepareContext(ctx, getAdminByUsername); err != nil {
		return nil, fmt.Errorf("error preparing query GetAdminByUsername: %w", err)
	}
	if q.getUserByUsernameStmt, err = db.PrepareContext(ctx, getUserByUsername); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByUsername: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.checkUserStmt != nil {
		if cerr := q.checkUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing checkUserStmt: %w", cerr)
		}
	}
	if q.checkUsernameStmt != nil {
		if cerr := q.checkUsernameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing checkUsernameStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.deleteUserStmt != nil {
		if cerr := q.deleteUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserStmt: %w", cerr)
		}
	}
	if q.getAdminByUsernameStmt != nil {
		if cerr := q.getAdminByUsernameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAdminByUsernameStmt: %w", cerr)
		}
	}
	if q.getUserByUsernameStmt != nil {
		if cerr := q.getUserByUsernameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByUsernameStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                     DBTX
	tx                     *sql.Tx
	checkUserStmt          *sql.Stmt
	checkUsernameStmt      *sql.Stmt
	createUserStmt         *sql.Stmt
	deleteUserStmt         *sql.Stmt
	getAdminByUsernameStmt *sql.Stmt
	getUserByUsernameStmt  *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                     tx,
		tx:                     tx,
		checkUserStmt:          q.checkUserStmt,
		checkUsernameStmt:      q.checkUsernameStmt,
		createUserStmt:         q.createUserStmt,
		deleteUserStmt:         q.deleteUserStmt,
		getAdminByUsernameStmt: q.getAdminByUsernameStmt,
		getUserByUsernameStmt:  q.getUserByUsernameStmt,
	}
}
