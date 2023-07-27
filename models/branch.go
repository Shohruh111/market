package models

type Branch struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Deleted   bool   `json:"deleted"`
	DeletedAt string `json:"deleted_at"`
}
type CreateBranch struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}
type UpdateBranch struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type BranchPrimaryKey struct {
	Id string `json:"id"`
}
type BranchGetListRequest struct {
	Offset          int    `json:"offset"`
	Limit           int    `json:"limit"`
	SearchByName    string `json:"search_by_name"`
	SearchByAddress string `json:"search_by_address"`
}

type BranchGetListResponse struct {
	Count   int       `json:"count"`
	Branchs []*Branch `json:"markets"`
}
