package handler

import (
	"market/models"
	"market/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Staff godoc
// @ID create_Staff
// @Router /staff [POST]
// @Summary Create Staff
// @Description Create Staff
// @Tags Staff
// @Accept json
// @Procedure json
// @Param Staff body models.CreateStaff true "CreateStaffRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateStaff(c *gin.Context) {

	var createStaff models.CreateStaff
	err := c.ShouldBindJSON(&createStaff)
	if err != nil {
		h.handlerResponse(c, "error Staff should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Staff().Create(c.Request.Context(), &createStaff)
	if err != nil {
		h.handlerResponse(c, "storage.Staff.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.Staff().GetByID(c.Request.Context(), &models.StaffPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Staff.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Staff resposne", http.StatusCreated, resp)
}

// GetByID Staff godoc
// @ID get_by_id_Staff
// @Router /staff/{id} [GET]
// @Summary Get By ID Staff
// @Description Get By ID Staff
// @Tags Staff
// @Accept json
// @Procedure json
// @Param id path string true "id"
func (h *handler) GetByIdStaff(c *gin.Context) {

	var id string = c.Param("id")
	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Staff().GetByID(c.Request.Context(), &models.StaffPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Staff.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id Staff resposne", http.StatusOK, resp)
}

// GetList Staff godoc
// @ID get_list_Staff
// @Router /staff [GET]
// @Summary Get List Staff
// @Description Get List Staff
// @Tags Staff
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param name query string false "name"
// @Param branchid query string false "branchID"
// @Param tarifid query string false "tarifID"
// @Param type query string false "type"
// @Param fromBalance query string false "fromBalance"
// @Param toBalance query string false "toBalance"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListStaff(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list Staff offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list Staff limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Staff().GetList(c.Request.Context(), &models.StaffGetListRequest{
		Offset:      offset,
		Limit:       limit,
		Search:      c.Query("name"),
		BranchID:    c.Query("branchID"),
		TarifID:     c.Query("TarifID"),
		Type:        c.Query("type"),
		FromBalance: c.GetInt("frommBalance"),
		ToBalance:   c.GetInt("toBalance"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.Staff.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list Staff resposne", http.StatusOK, resp)
}

// Update Staff godoc
// @ID update_Staff
// @Router /staff/{id} [PUT]
// @Summary Update Staff
// @Description Update Staff
// @Tags Staff
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param Staff body models.UpdateStaff true "UpdateStaffRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateStaff(c *gin.Context) {

	var (
		id          string = c.Param("id")
		updateStaff models.UpdateStaff
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateStaff)
	if err != nil {
		h.handlerResponse(c, "error Staff should bind json", http.StatusBadRequest, err.Error())
		return
	}

	updateStaff.Id = id
	rowsAffected, err := h.strg.Staff().Update(c.Request.Context(), &updateStaff)
	if err != nil {
		h.handlerResponse(c, "storage.Staff.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.Staff.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Staff().GetByID(c.Request.Context(), &models.StaffPrimaryKey{Id: updateStaff.Id})
	if err != nil {
		h.handlerResponse(c, "storage.Staff.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Staff resposne", http.StatusAccepted, resp)
}

// Delete Staff godoc
// @ID delete_Staff
// @Router /staff/{id} [DELETE]
// @Summary Delete Staff
// @Description Delete Staff
// @Tags Staff
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteStaff(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Staff().Delete(c.Request.Context(), &models.StaffPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Staff.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Staff resposne", http.StatusNoContent, nil)
}

// GetListTop Staff godoc
// @ID get_list_Staff_Top
// @Router /topstaff [GET]
// @Summary Get List Staff Top
// @Description Get List Staff Top
// @Tags Staff
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Param type query string true "type staff"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListTop(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list Staff offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list Staff Top limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Staff().GetListTop(c.Request.Context(), &models.StaffGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
		Type:   c.Query("type"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.Staff.get_list_top", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list Staff Top resposne", http.StatusOK, resp)
}
