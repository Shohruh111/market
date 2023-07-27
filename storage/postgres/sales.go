package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"market/models"
	"strconv"

	uuid "github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type salesRepo struct {
	db *pgxpool.Pool
}

func NewSalesRepo(db *pgxpool.Pool) *salesRepo {
	return &salesRepo{
		db: db,
	}
}

func (r *salesRepo) Create(ctx context.Context, req *models.CreateSales) (string, error) {

	trx, err := r.db.Begin(ctx)
	if err != nil {
		return "", nil
	}

	defer func() {
		if err != nil {
			trx.Rollback(ctx)
		} else {
			trx.Commit(ctx)
		}
	}()

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO sales(id, branch_id, shop_assistent_id, cashier_id, price,payment_type,client_name, updated_at)
		VALUES ($1, $2, $3,$4, $5,$6, $7, NOW())
	`

	_, err = trx.Exec(ctx, query,
		id,
		req.BranchId,
		req.AsistentId,
		req.CashierId,
		req.Price,
		req.PaymentType,
		req.ClientName,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *salesRepo) GetByID(ctx context.Context, req *models.SalesPrimaryKey) (*models.Sales, error) {

	var (
		query string

		id          sql.NullString
		branchId    sql.NullString
		asistenId   sql.NullString
		cashierId   sql.NullString
		price       int
		paymentType string
		clientName  string
		status      string
		createdAt   sql.NullString
		updatedAt   sql.NullString
	)

	query = `
		SELECT 
			id,
			branch_id,
			shop_assistent_id,
			cashier_id,
			price,
			payment_type,
			client_name,
			status,
			created_at,
			updated_at
		FROM sales 
		WHERE id = $1 and deleted = false
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&branchId,
		&asistenId,
		&cashierId,
		&price,
		&paymentType,
		&clientName,
		&status,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.Sales{
		Id:          id.String,
		BranchId:    branchId.String,
		AsistentId:  asistenId.String,
		CashierId:   cashierId.String,
		Price:       price,
		PaymentType: paymentType,
		ClientName:  clientName,
		Status:      status,
		CreatedAt:   createdAt.String,
		UpdatedAt:   updatedAt.String,
	}, nil
}

func (r *salesRepo) GetList(ctx context.Context, req *models.SalesGetListRequest) (*models.SalesGetListResponse, error) {

	var (
		resp   = &models.SalesGetListResponse{}
		query  string
		where  = " WHERE deleted = false "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		order  = " ORDER BY created_at DESC "
	)

	query = `
		SELECT 
			COUNT(*) OVER(),
			id,
			branch_id,
			shop_assistent_id,
			cashier_id,
			price,
			payment_type,
			client_name,
			status,
			created_at,
			updated_at
		FROM sales  
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.SearchByClientName != "" {
		where += ` AND client_name ILIKE '%' || '` + req.SearchByClientName + `' || '%'`
	}
	if req.SearchByBranchId != "" {
		where += ` AND branch_id = || '` + req.SearchByBranchId + `'|| `
	}
	if req.CashierId != "" {
		where += ` AND cashier_id = || '` + req.CashierId + `'|| `
	}
	if req.PaymentType != "" {
		where += ` AND payment_type = || '` + req.PaymentType + `'|| `
	}
	if req.Status != "" {
		where += ` AND status = || '` + req.Status + `'|| `
	}
	if req.ShopAsistentId != "" {
		where += ` AND shop_assistent_id = || '` + req.ShopAsistentId + `'|| `
	}
	if req.Price > 0 {
		price := strconv.Itoa(req.Price)
		where += ` AND price = || '` + price + `'|| `
	}

	query += where + order + offset + limit 

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id          sql.NullString
			branchId    sql.NullString
			asistenId   sql.NullString
			cashierId   sql.NullString
			price       int
			paymentType string
			clientName  string
			status      string
			createdAt   sql.NullString
			updatedAt   sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&branchId,
			&asistenId,
			&cashierId,
			&price,
			&paymentType,
			&clientName,
			&status,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Sales = append(resp.Sales, &models.Sales{
			Id:          id.String,
			BranchId:    branchId.String,
			AsistentId:  asistenId.String,
			CashierId:   cashierId.String,
			Price:       price,
			PaymentType: paymentType,
			ClientName:  clientName,
			Status:      status,
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		})
	}

	return resp, nil
}

func (r *salesRepo) Update(ctx context.Context, req *models.UpdateSales) (int64, error) {
	trx, err := r.db.Begin(ctx)
	if err != nil {
		return 0, nil
	}

	defer func() {
		if err != nil {
			trx.Rollback(ctx)
		} else {
			trx.Commit(ctx)
		}
	}()

	var (
		query string
		// params map[string]interface{}
	)

	query = `
		UPDATE
			sales
		SET
			shop_assistent_id = $1,
			cashier_id = $2,
			price = $3,
			payment_type = $4,
			status = $5,
			updated_at = NOW()
		WHERE id = $6
	`

	// params = map[string]interface{}{
	// 	"Sales_name":  req.SalesName,
	// 	"address":      req.Address,
	// 	"phone_number": req.PhoneNumber,
	// 	"id":           req.Id,
	// }

	// query, args := helper.ReplaceQueryParams(query, params)

	result, err := trx.Exec(ctx, query, req.AsistentId, req.CashierId, req.Price, req.PaymentType, req.Status, req.Id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *salesRepo) Delete(ctx context.Context, req *models.SalesPrimaryKey) error {

	_, err := r.db.Exec(ctx, "UPDATE sales set deleted = true, deleted_at = NOW() WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *salesRepo) SortBySalesAmount(ctx context.Context, req *models.SalesSortRequest) (*models.SalesSortResponse, error) {
	var (
		query string
		resp  = models.SalesSortResponse{}
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			s.created_at::DATE,
			b.name,
			SUM(s.price) AS total
		FROM sales AS s
		JOIN branch AS b ON s.branch_id = b.id
		WHERE s.created_at::DATE BETWEEN to_date('` + req.CreatedAtFrom + `' ,'YYYY-MM-DD') AND to_date('` + req.CreatedAtTo + `','YYYY-MM-DD') AND s.deleted = false
		GROUP BY s.created_at::DATE, b.name
		ORDER BY total DESC
		LIMIT 5
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			day   sql.NullString
			name  sql.NullString
			total int
		)

		err = rows.Scan(
			&resp.Count,
			&day,
			&name,
			&total,
		)
		if err != nil {
			return nil, err
		}
		resp.SortSales = append(resp.SortSales, &models.SalesSort{
			Day:   day.String,
			Name:  name.String,
			Total: total,
		})
	}

	return &resp, nil
}
