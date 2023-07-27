package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"market/models"

	uuid "github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type branchRepo struct {
	db *pgxpool.Pool
}

func NewBranchRepo(db *pgxpool.Pool) *branchRepo {
	return &branchRepo{
		db: db,
	}
}

func (r *branchRepo) Create(ctx context.Context, req *models.CreateBranch) (string, error) {

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
		INSERT INTO branch(id, name, address, updated_at)
		VALUES ($1, $2, $3, NOW())
	`

	_, err = trx.Exec(ctx, query,
		id,
		req.Name,
		req.Address,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *branchRepo) GetByID(ctx context.Context, req *models.BranchPrimaryKey) (*models.Branch, error) {

	var (
		query string

		id        string
		name      string
		address   string
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	query = `
		SELECT 
			id,
			name,
			address,
			created_at,
			updated_at
		FROM branch 
		WHERE id = $1 and deleted = false
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&address,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.Branch{
		Id:        id,
		Name:      name,
		Address:   address,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}, nil
}

func (r *branchRepo) GetList(ctx context.Context, req *models.BranchGetListRequest) (*models.BranchGetListResponse, error) {

	var (
		resp   = &models.BranchGetListResponse{}
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
			address,
			created_at,
			updated_at
		FROM branch

	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.SearchByName != "" {
		where += ` AND name ILIKE '%' || '` + req.SearchByName + `' || '%'`
	}

	if req.SearchByAddress != "" {
		where += ` AND address ILIKE '%' || '` + req.SearchByAddress + `' || '%'`
	}

	query += where + order + offset + limit 

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id        sql.NullString
			name      sql.NullString
			address   sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&address,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Branchs = append(resp.Branchs, &models.Branch{
			Id:        id.String,
			Name:      name.String,
			Address:   address.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return resp, nil
}

func (r *branchRepo) Update(ctx context.Context, req *models.UpdateBranch) (int64, error) {
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
			branch
		SET
			name = $1,
			address = $2,
			updated_at = NOW()
		WHERE id = $3
	`

	// params = map[string]interface{}{
	// 	"branch_name":  req.BranchName,
	// 	"address":      req.Address,
	// 	"phone_number": req.PhoneNumber,
	// 	"id":           req.Id,
	// }

	// query, args := helper.ReplaceQueryParams(query, params)

	result, err := trx.Exec(ctx, query, req.Name, req.Address, req.Id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *branchRepo) Delete(ctx context.Context, req *models.BranchPrimaryKey) error {

	_, err := r.db.Exec(ctx, "UPDATE branch set deleted = true, deleted_at = NOW() WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
