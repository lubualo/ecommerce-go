package product

import (
	"database/sql"
	"fmt"

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
		columns = append(columns, "Prod_CategoryId")
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
		update = update.Set("Prod_Title", p.Title)
	}
	if p.Description != "" {
		update = update.Set("Prod_Description", p.Description)
	}
	if p.Price > 0 {
		update = update.Set("Prod_Price", p.Price)
	}
	if p.CategoryId > 0 {
		update = update.Set("Prod_CategoryId", p.CategoryId)
	}
	if p.Stock > 0 {
		update = update.Set("Prod_Stock", p.Stock)
	}
	if p.Path != "" {
		update = update.Set("Prod_Path", p.Path)
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
	query, args, err := squirrel.
		Delete("products").
		Where(squirrel.Eq{"Prod_Id": id}).
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

func (r *repositorySQL) GetById(id int) (models.Product, error) {
	query, args, err := squirrel.
		Select("p.Prod_Id", "p.Prod_Title", "p.Prod_Description", "p.Prod_CreatedAt", "p.Prod_Updated", "p.Prod_Price", "p.Prod_Path", "p.Prod_CategoryId", "p.Prod_Stock", "c.Categ_Path").
		From("products p").
		Join("category c ON p.Prod_CategoryId = c.Categ_Id").
		Where(squirrel.Eq{"p.Prod_Id": id}).
		PlaceholderFormat(squirrel.Question).
		ToSql()

	fmt.Println(query)
	if err != nil {
		return models.Product{}, err
	}

	row := r.db.QueryRow(query, args...)
	var p models.Product
	err = row.Scan(&p.Id, &p.Title, &p.Description, &p.CreatedAt, &p.Updated, &p.Price, &p.Path, &p.CategoryId, &p.Stock, &p.CategoryPath)
	if err != nil {
		return models.Product{}, err
	}

	return p, nil
}

func (r *repositorySQL) GetBySlug(slug string) (models.Product, error) {
	query, args, err := squirrel.
		Select("Prod_Id", "Prod_Title", "Prod_Description", "Prod_CreatedAt", "Prod_Updated", "Prod_Price", "Prod_Path", "Prod_CategoryId", "Prod_Stock", "Categ_Path").
		From("products").
		Join("category ON products.Prod_CategoryId = category.Categ_Id").
		Where(squirrel.Eq{"Prod_Path": slug}).
		PlaceholderFormat(squirrel.Question).
		ToSql()

	if err != nil {
		return models.Product{}, err
	}

	row := r.db.QueryRow(query, args...)
	var p models.Product
	err = row.Scan(&p.Id, &p.Title, &p.Description, &p.CreatedAt, &p.Updated, &p.Price, &p.Path, &p.CategoryId, &p.Stock, &p.CategoryPath)
	if err != nil {
		return models.Product{}, err
	}

	return p, nil
}

func (r *repositorySQL) GetByCategoryId(catId int) ([]models.Product, error) {
	query, args, err := squirrel.
		Select("Prod_Id", "Prod_Title", "Prod_Description", "Prod_CreatedAt", "Prod_Updated", "Prod_Price", "Prod_Path", "Prod_CategoryId", "Prod_Stock", "Categ_Path").
		From("products").
		Join("category ON products.Prod_CategoryId = category.Categ_Id").
		Where(squirrel.Eq{"Prod_CategoryId": catId}).
		PlaceholderFormat(squirrel.Question).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.Id, &p.Title, &p.Description, &p.CreatedAt, &p.Updated, &p.Price, &p.Path, &p.CategoryId, &p.Stock, &p.CategoryPath); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (r *repositorySQL) GetByCategorySlug(slug string) ([]models.Product, error) {
	query, args, err := squirrel.
		Select("Prod_Id", "Prod_Title", "Prod_Description", "Prod_CreatedAt", "Prod_Updated", "Prod_Price", "Prod_Path", "Prod_CategoryId", "Prod_Stock", "Categ_Path").
		From("products").
		Join("category ON products.Prod_CategoryId = category.Categ_Id").
		Where(squirrel.Eq{"Categ_Path": slug}).
		PlaceholderFormat(squirrel.Question).
		ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.Id, &p.Title, &p.Description, &p.CreatedAt, &p.Updated, &p.Price, &p.Path, &p.CategoryId, &p.Stock, &p.CategoryPath); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (r *repositorySQL) SearchByText(search string, offset, limit int, sortBy, order string) ([]models.Product, error) {
	allowedSorts := map[string]string{
		"id":          "Prod_Id",
		"title":       "Prod_Title",
		"description": "Prod_Description",
		"price":       "Prod_Price",
		"category_id": "Prod_CategoryId",
		"stock":       "Prod_Stock",
		"created_at":  "Prod_CreatedAt",
	}
	dbSortBy, ok := allowedSorts[sortBy]
	if !ok {
		dbSortBy = "Prod_Title"
	}

	queryBuilder := squirrel.
		Select("Prod_Id", "Prod_Title", "Prod_Description", "Prod_CreatedAt", "Prod_Updated", "Prod_Price", "Prod_Path", "Prod_CategoryId", "Prod_Stock", "Categ_Path").
		From("products").
		Join("category ON products.Prod_CategoryId = category.Categ_Id").
		Where(squirrel.Or{
			squirrel.Like{"Prod_Title": "%" + search + "%"},
			squirrel.Like{"Prod_Description": "%" + search + "%"},
		}).
		OrderBy(fmt.Sprintf("%s %s", dbSortBy, order)).
		Offset(uint64(offset)).
		Limit(uint64(limit)).
		PlaceholderFormat(squirrel.Question)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.Id, &p.Title, &p.Description, &p.CreatedAt, &p.Updated, &p.Price, &p.Path, &p.CategoryId, &p.Stock, &p.CategoryPath); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (r *repositorySQL) GetAll(offset, limit int, sortBy, order string) ([]models.Product, error) {
	allowedSorts := map[string]string{
		"id":          "p.Prod_Id",
		"title":       "p.Prod_Title",
		"description": "p.Prod_Description",
		"price":       "p.Prod_Price",
		"category_id": "p.Prod_CategoryId",
		"stock":       "p.Prod_Stock",
		"created_at":  "p.Prod_CreatedAt",
	}

	dbSortBy, ok := allowedSorts[sortBy]
	if !ok {
		dbSortBy = "p.Prod_Title"
	}

	queryBuilder := squirrel.
		Select("Prod_Id", "Prod_Title", "Prod_Description", "Prod_CreatedAt", "Prod_Updated", "Prod_Price", "Prod_Path", "Prod_CategoryId", "Prod_Stock", "Categ_Path").
		From("products").
		Join("category ON products.Prod_CategoryId = category.Categ_Id").
		OrderBy(fmt.Sprintf("%s %s", dbSortBy, order)).
		Offset(uint64(offset)).
		Limit(uint64(limit)).
		PlaceholderFormat(squirrel.Question)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.Id, &p.Title, &p.Description, &p.CreatedAt, &p.Updated, &p.Price, &p.Path, &p.CategoryId, &p.Stock, &p.CategoryPath); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
