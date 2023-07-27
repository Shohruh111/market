package models

type Staff struct {
	Id         string `json:"id"`
	TarifId    string `json:"tarif_id"`
	BranchId   string `json:"branch_id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Balance    int    `json:"balance"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Deleted    bool   `json:"deleted"`
	Deleted_at string `json:"deleted_at"`
}
type CreateStaff struct {
	TarifId  string `json:"tarif_id"`
	BranchId string `json:"branch_id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}
type UpdateStaff struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Balance int    `json:"balance"`
}

type StaffPrimaryKey struct {
	Id string `json:"id"`
}
type StaffGetListRequest struct {
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	Search   string `json:"search"`
	BranchID string `json:"branch_id"`
	TarifID  string `json:"tarif_id"`
	Type     string `json:"type"`
	FromBalance int `json:"from"`
	ToBalance   int `json:"to"`
}
type StaffTop struct {
	Branch  string `json:"branch_name"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}
type StaffTopResponse struct {
	Count    int         `json:"count"`
	StaffTop []*StaffTop `json:"staff_top"`
}

type StaffGetListResponse struct {
	Count int      `json:"count"`
	Staff []*Staff `json:"staff"`
}
