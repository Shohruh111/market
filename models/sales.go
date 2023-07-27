package models

type Sales struct {
	Id          string `json:"id"`
	BranchId    string `json:"branch_id"`
	AsistentId  string `json:"shop_asistent_id"`
	CashierId   string `json:"cashier_id"`
	Price       int    `json:"price"`
	PaymentType string `json:"payment type"`
	ClientName  string `json:"client_name"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Deleted     bool   `json:"deleted"`
	DeletedAt   string `json:"deleted_at"`
}
type CreateSales struct {
	BranchId    string `json:"branch_id"`
	AsistentId  string `json:"shop_asistent_id"`
	CashierId   string `json:"cashier_id"`
	Price       int    `json:"price"`
	PaymentType string `json:"payment type"`
	ClientName  string `json:"client_name"`
}
type UpdateSales struct {
	Id          string `json:"id"`
	AsistentId  string `json:"shop_asistent_id"`
	CashierId   string `json:"cashier_id"`
	Price       int    `json:"price"`
	PaymentType string `json:"payment type"`
	Status      string `json:"status"`
}
type SalesPrimaryKey struct {
	Id string `json:"id"`
}
type SalesGetListRequest struct {
	Offset             int    `json:"offset"`
	Limit              int    `json:"limit"`
	SearchByClientName string `json:"search_by_client_name"`
	SearchByBranchId   string `json:"search_by_branch_id"`
	PaymentType        string `json:"payment_type"`
	ShopAsistentId     string `json:"shop_assistent_id"`
	CashierId          string `json:"cashier_id"`
	Status             string `json:"status"`
	Price              int    `json:"price"`
}

type SalesSort struct {
	Day   string `json:"day"`
	Name  string `json:"name"`
	Total int    `json:"total"`
}

type SalesSortRequest struct {
	CreatedAtFrom string `json:"created_at_from"`
	CreatedAtTo   string `json:"created_at_to"`
}
type SalesSortResponse struct {
	Count     int          `json:"count"`
	SortSales []*SalesSort `json:"sort_sales"`
}

type SalesGetListResponse struct {
	Count int      `json:"count"`
	Sales []*Sales `json:"sales"`
}
