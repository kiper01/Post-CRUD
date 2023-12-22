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
