package models

type StaffTarif struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	AmountForCash int    `json:"amount_for_cash"`
	AmountForCard int    `json:"amount_for card"`
	Created_at    string `json:"created_at"`
	Updated_at    string `json:"updated_at"`
	Deleted       bool   `json:"deleted"`
	Deleted_at    string `json:"deleted_at"`
}
type CreateStaffTarif struct {
	Name          string `json:"name"`
	Type          string `json:"type"`
	AmountForCash int    `json:"amount_for_cash"`
	AmountForCard int    `json:"amount_for card"`
}
type UpdateStaffTarif struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	AmountForCash int    `json:"amount_for_cash"`
	AmountForCard int    `json:"amount_for card"`
}

type StaffTarifPrimaryKey struct {
	Id string `json:"id"`
}
type StaffTarifGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type StaffTarifGetListResponse struct {
	Count      int           `json:"count"`
	StaffTarif []*StaffTarif `json:"staff_tarif"`
}
