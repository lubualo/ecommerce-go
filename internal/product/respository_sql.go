package product

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

func (r *repositorySQL) Insert(p models.Product) (int64, error) {
	columns := []string{}
	values := []interface{}{}

	columns = append(columns, "Prod_Title")
	values = append(values, p.Title)

	if p.Description != "" {
		columns = append(columns, "Prod_Description")
		values = append(values, p.Description)
	}
	if p.Price > 0 {
		columns = append(columns, "Prod_Price")
		values = append(values, p.Price)
	}
	if p.CategoryId > 0 {
		columns = append(columns, "Prod_CategId")
		values = append(values, p.CategoryId)
	}
	if p.Stock > 0 {
		columns = append(columns, "Prod_Stock")
		values = append(values, p.Stock)
	}
	if p.Path != "" {
		columns = append(columns, "Prod_Path")
		values = append(values, p.Description)
	}

	query, args, err := squirrel.
		Insert("products").
		Columns(columns...).
		Values(values...).
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

func (r *repositorySQL) Update(p models.Product) error {
	update := squirrel.
		Update("products").
		PlaceholderFormat(squirrel.Question).
		Set("Prod_Updated", squirrel.Expr("NOW()"))

	if p.Title != "" {
		update.Set("Prod_Title", p.Title)
	}
	if p.Description != "" {
		update.Set("Prod_Description", p.Description)
	}
	if p.Price > 0 {
		update.Set("Prod_Price", p.Price)
	}
	if p.CategoryId > 0 {
		update.Set("Prod_CategId", p.CategoryId)
	}
	if p.Stock > 0 {
		update.Set("Prod_Stock", p.Stock)
	}
	if p.Path != "" {
		update.Set("Prod_Path", p.Path)
	}

	query, args, err := update.
		Where(squirrel.Eq{"Prod_Id": p.Id}).
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
	return nil
	// query, args, err := squirrel.
	// 	Delete("category").
	// 	Where(squirrel.Eq{"Categ_Id": id}).
	// 	ToSql()
	// if err != nil {
	// 	return err
	// }
	// _, err = r.db.Exec(query, args...)
	// if err != nil {
	// 	return err
	// }
	// return nil
}

func (r *repositorySQL) GetById(id int) (models.Product, error) {
	return models.Product{}, nil
	// query, args, err := squirrel.
	// 	Select("Categ_Name", "Categ_Path").
	// 	From("category").
	// 	Where(squirrel.Eq{"Categ_Id": id}).
	// 	ToSql()

	// if err != nil {
	// 	return models.Category{}, err
	// }
	// row := r.db.QueryRow(query, args...)

	// var name, path string
	// if err := row.Scan(&name, &path); err != nil {
	// 	return models.Category{}, err
	// }

	// return models.Category{
	// 	CategID:   id,
	// 	CategName: name,
	// 	CategPath: path,
	// }, nil
}

func (r *repositorySQL) GetBySlug(slug string) ([]models.Product, error) {
	return []models.Product{}, nil
	// query, args, err := squirrel.
	// 	Select("Categ_Id", "Categ_Name", "Categ_Path").
	// 	From("category").
	// 	Where(squirrel.Like{"Categ_Path": "%" + slug + "%"}).
	// 	ToSql()

	// if err != nil {
	// 	return []models.Category{}, err
	// }

	// rows, err := r.db.Query(query, args...)
	// if err != nil {
	// 	return []models.Category{}, err
	// }
	// defer rows.Close()

	// var categories []models.Category
	// var id int
	// var name, path string
	// for rows.Next() {
	// 	if err := rows.Scan(&id, &name, &path); err != nil {
	// 		return []models.Category{}, err
	// 	}
	// 	categories = append(categories, models.Category{
	// 		CategID:   id,
	// 		CategName: name,
	// 		CategPath: path,
	// 	})
	// }
	// return categories, nil
}

func (r *repositorySQL) GetAll() ([]models.Product, error) {
	return []models.Product{}, nil
	// query, args, err := squirrel.
	// 	Select("Categ_Id", "Categ_Name", "Categ_Path").
	// 	From("category").
	// 	ToSql()

	// if err != nil {
	// 	return []models.Category{}, err
	// }

	// rows, err := r.db.Query(query, args...)
	// if err != nil {
	// 	return []models.Category{}, err
	// }
	// defer rows.Close()

	// var categories []models.Category
	// var id int
	// var name, path string
	// for rows.Next() {
	// 	if err := rows.Scan(&id, &name, &path); err != nil {
	// 		return []models.Category{}, err
	// 	}
	// 	categories = append(categories, models.Category{
	// 		CategID:   id,
	// 		CategName: name,
	// 		CategPath: path,
	// 	})
	// }
	// return categories, nil
}
