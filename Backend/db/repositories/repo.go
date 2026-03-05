package repositories

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"subscriptions/Backend/core/db"
)

type BaseRepository[T any] struct {
	pool  *pgxpool.Pool
	table string
}

func Construct[T any](pool *pgxpool.Pool, table string) *BaseRepository[T] {
	return &BaseRepository[T]{pool: pool, table: table}
}

func (baseRepository *BaseRepository[T]) Create(context *gin.Context, entity T) (T, error) {
	query, args, err := db.BuildInsertAutoReturning(baseRepository.table, entity)
	if err != nil {
		var zero T
		return zero, err
	}

	row := baseRepository.pool.QueryRow(context, query, args...)

	var out T
	targets, err := db.ScanTargetsByDBRet(&out)
	if err != nil {
		var zero T
		return zero, err
	}

	if err := row.Scan(targets...); err != nil {
		var zero T
		return zero, err
	}

	return out, nil
}

func (baseRepository *BaseRepository[T]) GetByID(context *gin.Context, id uuid.UUID) (T, error) {
	cols, err := db.DBRetColumnsFromType[T]()
	if err != nil {
		var zero T
		return zero, err
	}

	q := fmt.Sprintf(
		`SELECT %s FROM %s WHERE id = $1`,
		strings.Join(cols, ", "),
		baseRepository.table,
	)

	row := baseRepository.pool.QueryRow(context, q, id)

	var out T
	targets, err := db.ScanTargetsByDBRet(&out)
	if err != nil {
		var zero T
		return zero, err
	}

	if err := row.Scan(targets...); err != nil {
		var zero T
		return zero, err
	}

	return out, nil
}

func (baseRepository *BaseRepository[T]) UpdateByID(context *gin.Context, id uuid.UUID, entity T) (T, error) {
	cols, args, err := db.ExtractDBColumnsAndArgs(entity)
	if err != nil {
		var zero T
		return zero, err
	}

	cols, args = db.FilterOutColumn(cols, args, "id")

	if len(cols) == 0 {
		var zero T
		return zero, errors.New("no updatable columns (after filtering id)")
	}

	setParts := make([]string, 0, len(cols))
	for i, col := range cols {
		setParts = append(setParts, fmt.Sprintf("%s = $%d", col, i+1))
	}

	args = append(args, id)

	retCols, err := db.DBRetColumnsFromType[T]()
	if err != nil {
		var zero T
		return zero, err
	}

	q := fmt.Sprintf(
		`UPDATE %s SET %s WHERE id = $%d RETURNING %s`,
		baseRepository.table,
		strings.Join(setParts, ", "),
		len(args),
		strings.Join(retCols, ", "),
	)

	row := baseRepository.pool.QueryRow(context, q, args...)

	var out T
	targets, err := db.ScanTargetsByDBRet(&out)
	if err != nil {
		var zero T
		return zero, err
	}

	if err := row.Scan(targets...); err != nil {
		var zero T
		return zero, err
	}

	return out, nil
}

func (baseRepository *BaseRepository[T]) DeleteByID(context *gin.Context, id uuid.UUID) error {
	q := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, baseRepository.table)
	tag, err := baseRepository.pool.Exec(context, q, id)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (baseRepository *BaseRepository[T]) List(context *gin.Context, limit int, offset int) ([]T, error) {
	if limit <= 0 {
		limit = 50
	}
	if offset < 0 {
		offset = 0
	}

	cols, err := db.DBRetColumnsFromType[T]()
	if err != nil {
		return nil, err
	}

	orderBy := ""
	if db.ContainsString(cols, "id") {
		orderBy = " ORDER BY id"
	}

	q := fmt.Sprintf(
		`SELECT %s FROM %s%s LIMIT $1 OFFSET $2`,
		strings.Join(cols, ", "),
		baseRepository.table,
		orderBy,
	)

	rows, err := baseRepository.pool.Query(context, q, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]T, 0)

	for rows.Next() {
		var item T
		targets, err := db.ScanTargetsByDBRet(&item)
		if err != nil {
			return nil, err
		}
		if err := rows.Scan(targets...); err != nil {
			return nil, err
		}
		out = append(out, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return out, nil
}
