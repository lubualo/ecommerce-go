package address

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

func (r *repositorySQL) Insert(a models.Address, userUUID string) (int64, error) {
	columns := []string{
		"Add_UserID",
		"Add_Name",
		"Add_Title",
		"Add_Address",
		"Add_City",
		"Add_State",
		"Add_PostalCode",
		"Add_Phone",
	}
	values := []interface{}{
		userUUID,
		a.Name,
		a.Title,
		a.Address,
		a.City,
		a.State,
		a.PostalCode,
		a.Phone,
	}

	query, args, err := squirrel.
		Insert("addresses").
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

func (r *repositorySQL) Exists(id int) bool {
	query, args, err := squirrel.
		Select("1").
		From("addresses").
		Where(squirrel.Eq{"Add_Id": id}).
		Limit(1).
		PlaceholderFormat(squirrel.Question).
		ToSql()
	if err != nil {
		return false
	}
	var exists int
	err = r.db.QueryRow(query, args...).Scan(&exists)
	return err == nil
}

func (r *repositorySQL) Update(a models.Address) error {
	query, args, err := squirrel.
		Update("addresses").
		Set("Add_Name", a.Name).
		Set("Add_Title", a.Title).
		Set("Add_Address", a.Address).
		Set("Add_City", a.City).
		Set("Add_State", a.State).
		Set("Add_PostalCode", a.PostalCode).
		Set("Add_Phone", a.Phone).
		Where(squirrel.Eq{"Add_Id": a.Id}).
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
		Delete("addresses").
		Where(squirrel.Eq{"Add_Id": id}).
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

func (r *repositorySQL) GetAllByUserUUID(userUUID string) ([]models.Address, error) {
	query, args, err := squirrel.
		Select("Add_Id", "Add_Name", "Add_Title", "Add_Address", "Add_City", "Add_State", "Add_PostalCode", "Add_Phone").
		From("addresses").
		Where(squirrel.Eq{"Add_UserID": userUUID}).
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

	var addresses []models.Address
	for rows.Next() {
		var a models.Address
		if err := rows.Scan(&a.Id, &a.Name, &a.Title, &a.Address, &a.City, &a.State, &a.PostalCode, &a.Phone); err != nil {
			return nil, err
		}
		addresses = append(addresses, a)
	}

	return addresses, nil

}
