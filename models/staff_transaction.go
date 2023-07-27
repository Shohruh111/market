package models

type StaffTransaction struct {
	Id         string `json:"id"`
	SalesId    string `json:"sales_id"`
	Type       string `json:"type"`
	Text       string `json:"text"`
	Amount     int    `json:"amount"`
	StaffId    string `json:"staff_id"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Deleted    bool   `json:"deleted"`
	Deleted_at string `json:"deleted_at"`
}
type CreateStaffTransaction struct {
	SalesId string `json:"sales_id"`
	Type    string `json:"type"`
	Text    string `json:"text"`
	Amount  int    `json:"amount"`
	StaffId string `json:"staff_id"`
}
type UpdateStaffTransaction struct {
	Id      string `json:"id"`
	Type    string `json:"type"`
	Text    string `json:"text"`
	Amount  int    `json:"amount"`
	StaffId string `json:"staff_id"`
}

type StaffTransactionPrimaryKey struct {
	Id string `json:"id"`
}
type PrimaryKey struct {
	SalesId string `json:"sales_id"`
}
type StaffTransactionGetListRequest struct {
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
	SalesId    string `json:"sales_id"`
	FromAmount int    `json:"from_amount"`
	ToAmount   int    `json:"to_amount"`
	StaffId    string `json:"staff_id"`
	Type       string `json:"type"`
}

type StaffTransactionGetListResponse struct {
	Count            int                 `json:"count"`
	StaffTransaction []*StaffTransaction `json:"staff_transaction"`
}
