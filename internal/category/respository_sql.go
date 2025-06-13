package category

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/lubualo/ecommerce-go/models"
)

type repositorySQL struct {
	db *sql.DB
}

func NewSQLRepository(db *sql.DB) Storage {
	return &repositorySQL{db: db}
}

func (r *repositorySQL) Insert(c models.Category) (int64, error) {
	query, args, err := squirrel.
		Insert("category").
		Columns("Categ_Name", "Categ_Path").
		Values(c.CategName, c.CategPath).
		ToSql()

	if err != nil {
		return 0, err
	}

	result, err := r.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (r *repositorySQL) Update(c models.Category) error {
	query, args, err := squirrel.
		Update("category").
		Set("Categ_Name", c.CategName).
		Set("Categ_Path", c.CategPath).
		Where(squirrel.Eq{"Categ_Id": c.CategID}).
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

func (r *repositorySQL) Delete(id int) error {
	query, args, err := squirrel.
		Delete("category").
		Where(squirrel.Eq{"Categ_Id": id}).
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
