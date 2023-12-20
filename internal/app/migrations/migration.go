package migrations

// import (
// 	"database/sql"
// 	"fmt"
// )

const CreateTablePost = `
CREATE TABLE IF NOT EXISTS post (
    id TEXT PRIMARY KEY,
    code TEXT NOT NULL,
    name TEXT NOT NULL,
    river TEXT NOT NULL
);

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM pg_constraint
        WHERE conname = 'post_unique_constraint'
    ) THEN
        ALTER TABLE post
        ADD CONSTRAINT post_unique_constraint UNIQUE (code, name);
    END IF;
END
$$;

CREATE INDEX IF NOT EXISTS idx_post_code ON post(code);
`

// // ApplyMigration применяет миграцию в базе данных
// func ApplyMigration(db *sql.DB) error {
// 	_, err := db.Exec(CreateTablePost)
// 	if err != nil {
// 		return fmt.Errorf("ошибка при применении миграции: %v", err)
// 	}

// 	fmt.Println("Миграция успешно применена")
// 	return nil
// }
