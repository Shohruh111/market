package handler

import (
	"market/models"
	"market/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create StaffTransaction godoc
// @ID create_StaffTransaction
// @Router /staffTransaction [POST]
// @Summary Create StaffTransaction
// @Description Create StaffTransaction
// @Tags StaffTransaction
// @Accept json
// @Procedure json
// @Param StaffTransaction body models.CreateStaffTransaction true "CreateStaffTransactionRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateStaffTransaction(c *gin.Context) {

	var createStaffTransaction models.CreateStaffTransaction
	err := c.ShouldBindJSON(&createStaffTransaction)
	if err != nil {
		h.handlerResponse(c, "error StaffTransaction should bind json", http.StatusBadRequest, err.Error())
		return
	}

	_, err = h.strg.StaffTransaction().Create(c.Request.Context(), &createStaffTransaction)
	if err != nil {
		h.handlerResponse(c, "storage.StaffTransaction.create", http.StatusInternalServerError, err.Error())
		return
	}

	// resp, err := h.strg.StaffTransaction().GetByID(c.Request.Context(), &models.StaffTransactionPrimaryKey{SalesId: id})
	// if err != nil {
	// 	h.handlerResponse(c, "storage.StaffTransaction.getById", http.StatusInternalServerError, err.Error())
	// 	return
	// }

	h.handlerResponse(c, "create StaffTransaction resposne", http.StatusCreated, "created successfully")
}

// GetByID StaffTransaction godoc
// @ID get_by_id_StaffTransaction
// @Router /staffTransaction/{id} [GET]
// @Summary Get By ID StaffTransaction
// @Description Get By ID StaffTransaction
// @Tags StaffTransaction
// @Accept json
// @Procedure json
// @Param id path string true "id"
func (h *handler) GetByIdStaffTransaction(c *gin.Context) {

	var id string = c.Param("id")
	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.StaffTransaction().GetByID(c.Request.Context(), &models.StaffTransactionPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.StaffTransaction.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id StaffTransaction resposne", http.StatusOK, resp)
}

// GetList StaffTransaction godoc
// @ID get_list_StaffTransaction
// @Router /staffTransaction [GET]
// @Summary Get List StaffTransaction
// @Description Get List StaffTransaction
// @Tags StaffTransaction
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param salesid query string false "salesID"
// @Param staffid query string false "staffID"
// @Param type query string false "type"
// @Param fromBalance query int false "fromBalance"
// @Param toBalance query int false "toBalance"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListStaffTransaction(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list StaffTransaction offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list StaffTransaction limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.StaffTransaction().GetList(c.Request.Context(), &models.StaffTransactionGetListRequest{
		Offset:     offset,
		Limit:      limit,
		SalesId:    c.Query("salesID"),
		StaffId:    c.Query("staffID"),
		Type:       c.Query("type"),
		FromAmount: c.GetInt("fromBalance"),
		ToAmount:   c.GetInt("toBalance"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.StaffTransaction.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list StaffTransaction resposne", http.StatusOK, resp)
}

// Update StaffTransaction godoc
// @ID update_StaffTransaction
// @Router /staffTransaction/{id} [PUT]
// @Summary Update StaffTransaction
// @Description Update StaffTransaction
// @Tags StaffTransaction
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param StaffTransaction body models.UpdateStaffTransaction true "UpdateStaffTransactionRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateStaffTransaction(c *gin.Context) {

	var (
		id                     string = c.Param("id")
		updateStaffTransaction models.UpdateStaffTransaction
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateStaffTransaction)
	if err != nil {
		h.handlerResponse(c, "error StaffTransaction should bind json", http.StatusBadRequest, err.Error())
		return
	}

	updateStaffTransaction.Id = id
	rowsAffected, err := h.strg.StaffTransaction().Update(c.Request.Context(), &updateStaffTransaction)
	if err != nil {
		h.handlerResponse(c, "storage.StaffTransaction.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.StaffTransaction.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.StaffTransaction().GetByID(c.Request.Context(), &models.StaffTransactionPrimaryKey{Id: updateStaffTransaction.Id})
	if err != nil {
		h.handlerResponse(c, "storage.StaffTransaction.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create StaffTransaction resposne", http.StatusAccepted, resp)
}

// Delete StaffTransaction godoc
// @ID delete_StaffTransaction
// @Router /staffTransaction/{id} [DELETE]
// @Summary Delete StaffTransaction
// @Description Delete StaffTransaction
// @Tags StaffTransaction
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteStaffTransaction(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.StaffTransaction().Delete(c.Request.Context(), &models.StaffTransactionPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.StaffTransaction.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create StaffTransaction resposne", http.StatusNoContent, nil)
}
