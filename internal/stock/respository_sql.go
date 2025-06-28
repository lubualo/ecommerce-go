package stock

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
)

type repositorySQL struct {
	db *sql.DB
}

func NewSQLRepository(db *sql.DB) Storage {
	return &repositorySQL{db: db}
}

func (r *repositorySQL) UpdateStock(productId, delta int) error {
	query, args, err := squirrel.
		Update("products").
		PlaceholderFormat(squirrel.Question).
		Set("Prod_Updated", squirrel.Expr("NOW()")).
		Set("Prod_Stock", squirrel.Expr("Prod_Stock + ?", delta)).
		Where(squirrel.Eq{"Prod_Id": productId}).
		ToSql()

	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}
