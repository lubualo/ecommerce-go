package order

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

func (r *repositorySQL) Insert(o models.Order) (int64, error) {
	columns := []string{}
	values := []interface{}{}

	columns = append(columns, "Order_UserUUID")
	values = append(values, o.UserUUID)
	columns = append(columns, "Order_AddressId")
	values = append(values, o.AddressId)
	columns = append(columns, "Order_Total")
	values = append(values, o.Total)

	query, args, err := squirrel.
		Insert("orders").
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

	orderId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// Insert order details
	for _, detail := range o.Details {
		_, err := squirrel.Insert("orders_detail").
			Columns("OD_OrderId", "OD_ProdId", "OD_Quantity", "OD_Price").
			Values(orderId, detail.ProductId, detail.Quantity, detail.Price).
			PlaceholderFormat(squirrel.Question).
			RunWith(r.db).
			Exec()
		if err != nil {
			return 0, err
		}
	}

	return orderId, nil
}

func (r *repositorySQL) GetById(id int) (models.Order, error) {
	query, args, err := squirrel.
		Select("Order_Id", "Order_UserUUID", "Order_AddressId", "Order_Total").
		From("orders").
		Where(squirrel.Eq{"Order_Id": id}).
		PlaceholderFormat(squirrel.Question).
		ToSql()

	if err != nil {
		return models.Order{}, err
	}

	row := r.db.QueryRow(query, args...)
	var o models.Order
	err = row.Scan(&o.Id, &o.UserUUID, &o.AddressId, &o.Total)
	if err != nil {
		return models.Order{}, err
	}

	// Fetch order details
	detailsQuery, detailsArgs, err := squirrel.
		Select("OD_OrderId", "OD_ProdId", "OD_Quantity", "OD_Price").
		From("orders_detail").
		Where(squirrel.Eq{"OD_OrderId": id}).
		PlaceholderFormat(squirrel.Question).
		ToSql()
	if err != nil {
		return models.Order{}, err
	}

	rows, err := r.db.Query(detailsQuery, detailsArgs...)
	if err != nil {
		return models.Order{}, err
	}
	defer rows.Close()

	var details []models.OrdersDetails
	for rows.Next() {
		var d models.OrdersDetails
		if err := rows.Scan(&d.OrderId, &d.ProductId, &d.Quantity, &d.Price); err != nil {
			return models.Order{}, err
		}
		details = append(details, d)
	}
	o.Details = details

	return o, nil
}

func (r *repositorySQL) GetAllByUserUUID(userUUID string, offset, limit int, from_date, to_date string) ([]models.Order, error) {
	var orders []models.Order

	queryBuilder := squirrel.
		Select("Order_Id", "Order_UserUUID", "Order_AddressId", "Order_Total", "Order_Date").
		From("orders").
		Where(squirrel.Eq{"Order_UserUUID": userUUID}).
		PlaceholderFormat(squirrel.Question)

	if from_date != "" {
		queryBuilder = queryBuilder.Where(squirrel.GtOrEq{"Order_Date": from_date})
	}
	if to_date != "" {
		queryBuilder = queryBuilder.Where(squirrel.LtOrEq{"Order_Date": to_date})
	}

	queryBuilder = queryBuilder.Offset(uint64(offset)).Limit(uint64(limit))

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var o models.Order
		err := rows.Scan(&o.Id, &o.UserUUID, &o.AddressId, &o.Total, &o.Date)
		if err != nil {
			return nil, err
		}
		// Fetch order details
		detailsQuery, detailsArgs, err := squirrel.
			Select("OD_OrderId", "OD_ProdId", "OD_Quantity", "OD_Price").
			From("orders_detail").
			Where(squirrel.Eq{"OD_OrderId": o.Id}).
			PlaceholderFormat(squirrel.Question).
			ToSql()
		if err != nil {
			return nil, err
		}
		detailsRows, err := r.db.Query(detailsQuery, detailsArgs...)
		if err != nil {
			return nil, err
		}
		var details []models.OrdersDetails
		for detailsRows.Next() {
			var d models.OrdersDetails
			if err := detailsRows.Scan(&d.OrderId, &d.ProductId, &d.Quantity, &d.Price); err != nil {
				detailsRows.Close()
				return nil, err
			}
			details = append(details, d)
		}
		detailsRows.Close()
		o.Details = details
		orders = append(orders, o)
	}
	return orders, nil
}
