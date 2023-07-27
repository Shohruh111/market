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

type staffRepo struct {
	db *pgxpool.Pool
}

func NewStaffRepo(db *pgxpool.Pool) *staffRepo {
	return &staffRepo{
		db: db,
	}
}

func (r *staffRepo) Create(ctx context.Context, req *models.CreateStaff) (string, error) {

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
		INSERT INTO staff(id, branch_id, tarif_id, type, name, updated_at)
		VALUES ($1, $2, $3,$4, $5, NOW())
	`

	_, err = trx.Exec(ctx, query,
		id,
		req.BranchId,
		req.TarifId,
		req.Type,
		req.Name,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *staffRepo) GetByID(ctx context.Context, req *models.StaffPrimaryKey) (*models.Staff, error) {

	var (
		query string

		id        sql.NullString
		name      string
		branchId  sql.NullString
		tarifId   sql.NullString
		typeStaff string
		balance   int
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	query = `
		SELECT 
			id,
			branch_id,
			tarif_id,
			type,
			name,
			balance,
			created_at,
			updated_at
		FROM staff 
		WHERE id = $1 and deleted = false
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&branchId,
		&tarifId,
		&typeStaff,
		&name,
		&balance,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.Staff{
		Id:         id.String,
		Name:       name,
		BranchId:   branchId.String,
		TarifId:    tarifId.String,
		Type:       typeStaff,
		Balance:    balance,
		Created_at: createdAt.String,
		Updated_at: updatedAt.String,
	}, nil
}

func (r *staffRepo) GetList(ctx context.Context, req *models.StaffGetListRequest) (*models.StaffGetListResponse, error) {

	var (
		resp   = &models.StaffGetListResponse{}
		query  string
		where         = " WHERE deleted = false "
		offset        = " OFFSET 0"
		limit         = " LIMIT 10"
		from   string = strconv.Itoa(req.FromBalance)
		to     string = strconv.Itoa(req.ToBalance)
		order         = " ORDER BY created_at DESC"
	)

	query = `
		SELECT 
			COUNT(*) OVER(),
			id,
			branch_id,
			tarif_id,
			type,
			name,
			balance,
			created_at,
			updated_at
		FROM staff 
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
	if req.BranchID != "" {
		where += ` AND branch_id ILIKE '%' || '` + req.BranchID + `' || '%'`
	}
	if req.TarifID != "" {
		where += ` AND tarif_id ILIKE '%' || '` + req.TarifID + `' || '%'`
	}
	if req.FromBalance > 0 || req.ToBalance > 0 {
		where += ` AND balance between  '` + from + `' AND '` + to + `' `
	}
	if req.Type != "" {
		where += ` AND type ILIKE '%' || '` + req.Type + `' || '%'`
	}

	query += where + order + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id        sql.NullString
			name      string
			branchId  sql.NullString
			tarifId   sql.NullString
			typeStaff string
			balance   int
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&branchId,
			&tarifId,
			&typeStaff,
			&name,
			&balance,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Staff = append(resp.Staff, &models.Staff{
			Id:         id.String,
			BranchId:   branchId.String,
			TarifId:    tarifId.String,
			Type:       typeStaff,
			Name:       name,
			Balance:    balance,
			Created_at: createdAt.String,
			Updated_at: updatedAt.String,
		})
	}

	return resp, nil
}

func (r *staffRepo) Update(ctx context.Context, req *models.UpdateStaff) (int64, error) {
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
			staff
		SET
			name = $1,
			type = $2,
			balance = $3,
			updated_at = NOW()
		WHERE id = $4
	`

	// params = map[string]interface{}{
	// 	"Staff_name":  req.StaffName,
	// 	"address":      req.Address,
	// 	"phone_number": req.PhoneNumber,
	// 	"id":           req.Id,
	// }

	// query, args := helper.ReplaceQueryParams(query, params)

	result, err := trx.Exec(ctx, query, req.Name, req.Type, req.Balance, req.Id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *staffRepo) Delete(ctx context.Context, req *models.StaffPrimaryKey) error {

	_, err := r.db.Exec(ctx, "UPDATE staff set deleted = true, deleted_at = NOW() WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *staffRepo) GetListTop(ctx context.Context, req *models.StaffGetListRequest) (*models.StaffTopResponse, error) {

	var (
		resp      = &models.StaffTopResponse{}
		query     string
		where     = " WHERE s.deleted = false "
		offset    = " OFFSET 0"
		limit     = " LIMIT 10"
		order     = " ORDER BY s.balance DESC"
		typeStaff = " s.type=$1"
	)

	query = `
		SELECT 
			COUNT(*) OVER(),
			s.name as name,
			b.name as branch,
			s.balance as balance
		FROM staff as s
		JOIN branch AS b ON b.id = s.branch_id
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

	query += where + typeStaff + order + offset + limit

	rows, err := r.db.Query(ctx, query, req.Type)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			name    string
			branch  string
			balance int
		)

		err := rows.Scan(
			&resp.Count,
			&name,
			&branch,
			&balance,
		)

		if err != nil {
			return nil, err
		}

		resp.StaffTop = append(resp.StaffTop, &models.StaffTop{
			Name:    name,
			Branch:  branch,
			Balance: balance,
		})
	}

	return resp, nil
}
