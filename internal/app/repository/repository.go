package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kiper01/Post-CRUD/internal/app/model"
)

type PostInfoRepository struct {
	dbPool *pgxpool.Pool
}

func NewPostInfoRepository(pool *pgxpool.Pool) *PostInfoRepository {
	return &PostInfoRepository{dbPool: pool}
}

func (r *PostInfoRepository) AddPostValue(ctx context.Context, value model.PostValue) error {

	sql := `INSERT INTO post (id, code, name, river)
            VALUES ($1, $2, $3, $4)
            ON CONFLICT (id) DO NOTHING;`

	commandTag, err := r.dbPool.Exec(ctx, sql, value.ID, value.Code, value.Name, value.River)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("post already exists")
	}

	return nil
}

func (r *PostInfoRepository) RemovePostValue(ctx context.Context, id string) error {

	sql := `DELETE FROM post WHERE id = $1;`

	commandTag, err := r.dbPool.Exec(ctx, sql, id)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("no post value found with id: %s", id)
	}

	return nil
}

func (r *PostInfoRepository) UpdatePostValues(ctx context.Context, values []model.PostValue) error {

	tx, err := r.dbPool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	for _, value := range values {
		sql := `UPDATE post SET code = $1, name = $2, river = $3 WHERE id = $4 AND NOT EXISTS (
                    SELECT 1 FROM post WHERE code = $1 AND name = $2 AND river = $3 AND id != $4
                );`

		commandTag, err := tx.Exec(ctx, sql, value.Code, value.Name, value.River, value.ID)
		if err != nil {
			return err
		}

		if commandTag.RowsAffected() == 0 {
			return fmt.Errorf("update failed for post value with id: %s", value.ID)
		}
	}

	return tx.Commit(ctx)
}

func (r *PostInfoRepository) GetPostValues(ctx context.Context, code int, page, pageSize int) ([]model.PostValue, int, error) {

	var postValues []model.PostValue
	var totalCount int

	query := `SELECT id, code, name, river FROM post WHERE code = $1 ORDER BY code DESC LIMIT $2 OFFSET $3`
	rows, err := r.dbPool.Query(ctx, query, code, pageSize, (page-1)*pageSize)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	for rows.Next() {
		var pv model.PostValue
		if err := rows.Scan(&pv.ID, &pv.Code, &pv.Name, &pv.River); err != nil {
			return nil, 0, err
		}
		postValues = append(postValues, pv)
	}

	countQuery := `SELECT COUNT(*) FROM post WHERE code = $1`
	err = r.dbPool.QueryRow(ctx, countQuery, code).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	return postValues, totalCount, nil
}

func (r *PostInfoRepository) GetPostValuesByCodeOrId(ctx context.Context, id string, code int) ([]model.PostValue, error) {

	var postValues []model.PostValue

	query := `SELECT id, code, name, river FROM post WHERE id = $1 OR code = $2`
	rows, err := r.dbPool.Query(ctx, query, id, code)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	count := 0
	for rows.Next() {
		var pv model.PostValue
		if count > 1 {
			return nil, fmt.Errorf("returned more than one post value")
		}
		if err := rows.Scan(&pv.ID, &pv.Code, &pv.Name, &pv.River); err != nil {
			return nil, err
		}
		postValues = append(postValues, pv)
		count++
	}

	if count == 0 {
		return nil, fmt.Errorf("no post value found with id: %s or code: %d", id, code)
	}

	return postValues, nil
}
