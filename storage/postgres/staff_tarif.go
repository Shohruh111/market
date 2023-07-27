package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"market/models"

	uuid "github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type staffTarifRepo struct {
	db *pgxpool.Pool
}

func NewStaffTarifRepo(db *pgxpool.Pool) *staffTarifRepo {
	return &staffTarifRepo{
		db: db,
	}
}

func (r *staffTarifRepo) Create(ctx context.Context, req *models.CreateStaffTarif) (string, error) {

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
		INSERT INTO staff_tarif(id, type, name,amount_for_cash,amount_for_card, updated_at)
		VALUES ($1, $2, $3,$4, $5, NOW())
	`

	_, err = trx.Exec(ctx, query,
		id,
		req.Type,
		req.Name,
		req.AmountForCash,
		req.AmountForCard,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *staffTarifRepo) GetByID(ctx context.Context, req *models.StaffTarifPrimaryKey) (*models.StaffTarif, error) {

	var (
		query string

		id            sql.NullString
		name          string
		Type          string
		amountForCash int
		amountForCard int
		// typeStaffTarif string
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	query = `
		SELECT 
			id,
			name,
			type,
			amount_for_cash,
			amount_for_card,

			created_at,
			updated_at
		FROM staff_tarif
		WHERE id = $1 and deleted = false
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&Type,
		&amountForCash,
		&amountForCard,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.StaffTarif{
		Id:            id.String,
		Name:          name,
		Type:          Type,
		AmountForCash: amountForCash,
		AmountForCard: amountForCard,
		Created_at:    createdAt.String,
		Updated_at:    updatedAt.String,
	}, nil
}

func (r *staffTarifRepo) GetList(ctx context.Context, req *models.StaffTarifGetListRequest) (*models.StaffTarifGetListResponse, error) {

	var (
		resp   = &models.StaffTarifGetListResponse{}
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
			name,
			type,
			amount_for_cash,
			amount_for_card,
			created_at,
			updated_at
		FROM staff_tarif
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND name ILIKE '%' || '` + req.Search + `' || '%'`
	}

	query += where + order + offset + limit 

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id            sql.NullString
			name          string
			Type          string
			amountForCash int
			amountForCard int
			createdAt     sql.NullString
			updatedAt     sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&Type,
			&amountForCash,
			&amountForCard,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.StaffTarif = append(resp.StaffTarif, &models.StaffTarif{
			Id:            id.String,
			Name:          name,
			Type:          Type,
			AmountForCash: amountForCash,
			AmountForCard: amountForCard,
			Created_at:    createdAt.String,
			Updated_at:    updatedAt.String,
		})
	}

	return resp, nil
}

func (r *staffTarifRepo) Update(ctx context.Context, req *models.UpdateStaffTarif) (int64, error) {
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
			staff_tarif
		SET
			name = $1,
			type = $2,
			amount_for_cash = $3,
			amount_for_card = $4,
			updated_at = NOW()
		WHERE id = $5 
	`

	// params = map[string]interface{}{
	// 	"StaffTarif_name":  req.StaffTarifName,
	// 	"address":      req.Address,
	// 	"phone_number": req.PhoneNumber,
	// 	"id":           req.Id,
	// }

	// query, args := helper.ReplaceQueryParams(query, params)

	result, err := trx.Exec(ctx, query, req.Name, req.Type, req.AmountForCash, req.AmountForCard, req.Id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *staffTarifRepo) Delete(ctx context.Context, req *models.StaffTarifPrimaryKey) error {

	_, err := r.db.Exec(ctx, "UPDATE staff_tarif set deleted = true, deleted_at = NOW() WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
