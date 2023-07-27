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

type staffTransactionRepo struct {
	db *pgxpool.Pool
}

func NewStaffTransactionRepo(db *pgxpool.Pool) *staffTransactionRepo {
	return &staffTransactionRepo{
		db: db,
	}
}

func (r *staffTransactionRepo) Create(ctx context.Context, req *models.CreateStaffTransaction) (string, error) {

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
		INSERT INTO staff_transaction(id,sales_id, type,text, amount,staff_id, updated_at)
		VALUES ($1, $2, $3,$4, $5, $6, NOW())
	`

	_, err = trx.Exec(ctx, query,
		id,
		req.SalesId,
		req.Type,
		req.Text,
		req.Amount,
		req.StaffId,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *staffTransactionRepo) GetByID(ctx context.Context, req *models.StaffTransactionPrimaryKey) (*models.StaffTransaction, error) {

	var (
		query string

		id        sql.NullString
		salesId   sql.NullString
		text      sql.NullString
		Type      string
		amount    int
		staffId   sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	query = `
		SELECT 
			id,
			sales_id,
			type,
			text,
			amount,
			staff_id,
			created_at,
			updated_at
		FROM staff_transaction
		WHERE id = $1 and deleted = false
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&salesId,
		&text,
		&Type,
		&amount,
		&staffId,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.StaffTransaction{
		Id:         id.String,
		SalesId:    salesId.String,
		Text:       text.String,
		Type:       Type,
		Amount:     amount,
		StaffId:    staffId.String,
		Created_at: createdAt.String,
		Updated_at: updatedAt.String,
	}, nil
}

func (r *staffTransactionRepo) GetList(ctx context.Context, req *models.StaffTransactionGetListRequest) (*models.StaffTransactionGetListResponse, error) {

	var (
		resp   = &models.StaffTransactionGetListResponse{}
		query  string
		where  = " WHERE deleted = false "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		from   = strconv.Itoa(req.FromAmount)
		to     = strconv.Itoa(req.ToAmount)
		order  = " ORDER BY created_at DESC"
	)

	query = `
		SELECT 
			COUNT(*) OVER(),
			id,
			sales_id,
			type,
			text,
			amount,
			staff_id,
			created_at,
			updated_at
		FROM staff_transaction
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.SalesId != "" {
		where += ` AND sales_id = '` + req.SalesId + `' `
	}

	if req.StaffId != "" {
		where += ` AND staff_id = '` + req.StaffId + `' `
	}
	if req.Type != "" {
		where += ` AND type ='` + req.Type + `' `
	}
	if req.FromAmount > 0 || req.ToAmount > 0 {
		where += ` AND amount BETWEEN  '` + from + `' AND  '` + to
	}

	query += where + order + offset + limit 

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id        sql.NullString
			salesId   sql.NullString
			text      sql.NullString
			Type      string
			amount    int
			staffId   sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&salesId,
			&text,
			&Type,
			&amount,
			&staffId,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.StaffTransaction = append(resp.StaffTransaction, &models.StaffTransaction{
			Id:         id.String,
			SalesId:    salesId.String,
			Text:       text.String,
			Type:       Type,
			Amount:     amount,
			StaffId:    staffId.String,
			Created_at: createdAt.String,
			Updated_at: updatedAt.String,
		})
	}

	return resp, nil
}

func (r *staffTransactionRepo) Update(ctx context.Context, req *models.UpdateStaffTransaction) (int64, error) {
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
			staff_transaction
		SET
			type = $1,
			text =$2,
			amount= $3,
			staff_id = $4,
			updated_at = NOW()
		WHERE id = $5
	`

	// params = map[string]interface{}{
	// 	"StaffTransaction_name":  req.StaffTransactionName,
	// 	"address":      req.Address,
	// 	"phone_number": req.PhoneNumber,
	// 	"id":           req.Id,
	// }

	// query, args := helper.ReplaceQueryParams(query, params)

	result, err := trx.Exec(ctx, query, req.Type, req.Text, req.Amount, req.StaffId, req.Id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *staffTransactionRepo) Delete(ctx context.Context, req *models.StaffTransactionPrimaryKey) error {
	var (
		query string
	)

	query = `
	"UPDATE 
		staff_transaction
		set 
		deleted = true, 
		deleted_at = NOW() 
		WHERE id = $1"
	`

	_, err := r.db.Exec(ctx, query, req.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *staffTransactionRepo) GetBySalesID(ctx context.Context, req *models.PrimaryKey) (*models.StaffTransactionGetListResponse, error) {

	var (
		query string
		resp  = models.StaffTransactionGetListResponse{}
	)

	query = `
		SELECT 
			id,
			sales_id,
			type,
			text,
			amount,
			staff_id,
			created_at,
			updated_at
		FROM staff_transaction
		WHERE sales_id = $1 and deleted = false
	`

	rows, err := r.db.Query(ctx, query, req.SalesId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id        sql.NullString
			salesId   sql.NullString
			text      sql.NullString
			Type      string
			amount    int
			staffId   sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&salesId,
			&text,
			&Type,
			&amount,
			&staffId,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.StaffTransaction = append(resp.StaffTransaction, &models.StaffTransaction{
			Id:         id.String,
			SalesId:    salesId.String,
			Text:       text.String,
			Type:       Type,
			Amount:     amount,
			StaffId:    staffId.String,
			Created_at: createdAt.String,
			Updated_at: updatedAt.String,
		})
	}

	return &resp, nil
}
