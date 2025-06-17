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
		PlaceholderFormat(squirrel.Question).
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
		PlaceholderFormat(squirrel.Question).
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
		PlaceholderFormat(squirrel.Question).
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

func (r *repositorySQL) GetById(id int) (models.Category, error) {
	query, args, err := squirrel.
		Select("Categ_Name", "Categ_Path").
		From("category").
		Where(squirrel.Eq{"Categ_Id": id}).
		PlaceholderFormat(squirrel.Question).
		ToSql()

	if err != nil {
		return models.Category{}, err
	}
	row := r.db.QueryRow(query, args...)

	var name, path string
	if err := row.Scan(&name, &path); err != nil {
		return models.Category{}, err
	}

	return models.Category{
		CategID:   id,
		CategName: name,
		CategPath: path,
	}, nil
}

func (r *repositorySQL) GetBySlug(slug string) ([]models.Category, error) {
	query, args, err := squirrel.
		Select("Categ_Id", "Categ_Name", "Categ_Path").
		From("category").
		Where(squirrel.Like{"Categ_Path": "%" + slug + "%"}).
		PlaceholderFormat(squirrel.Question).
		ToSql()

	if err != nil {
		return []models.Category{}, err
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return []models.Category{}, err
	}
	defer rows.Close()

	var categories []models.Category
	var id int
	var name, path string
	for rows.Next() {
		if err := rows.Scan(&id, &name, &path); err != nil {
			return []models.Category{}, err
		}
		categories = append(categories, models.Category{
			CategID:   id,
			CategName: name,
			CategPath: path,
		})
	}
	return categories, nil
}

func (r *repositorySQL) GetAll() ([]models.Category, error) {
	query, args, err := squirrel.
		Select("Categ_Id", "Categ_Name", "Categ_Path").
		From("category").
		PlaceholderFormat(squirrel.Question).
		ToSql()

	if err != nil {
		return []models.Category{}, err
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return []models.Category{}, err
	}
	defer rows.Close()

	var categories []models.Category
	var id int
	var name, path string
	for rows.Next() {
		if err := rows.Scan(&id, &name, &path); err != nil {
			return []models.Category{}, err
		}
		categories = append(categories, models.Category{
			CategID:   id,
			CategName: name,
			CategPath: path,
		})
	}
	return categories, nil
}
