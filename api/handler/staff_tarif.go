package handler

import (
	"market/models"
	"market/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create StaffTarif godoc
// @ID create_StaffTarif
// @Router /staffTarif [POST]
// @Summary Create StaffTarif
// @Description Create StaffTarif
// @Tags StaffTarif
// @Accept json
// @Procedure json
// @Param StaffTarif body models.CreateStaffTarif true "CreateStaffTarifRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateStaffTarif(c *gin.Context) {

	var createStaffTarif models.CreateStaffTarif
	err := c.ShouldBindJSON(&createStaffTarif)
	if err != nil {
		h.handlerResponse(c, "error StaffTarif should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.StaffTarif().Create(c.Request.Context(), &createStaffTarif)
	if err != nil {
		h.handlerResponse(c, "storage.StaffTarif.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.StaffTarif().GetByID(c.Request.Context(), &models.StaffTarifPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.StaffTarif.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create StaffTarif resposne", http.StatusCreated, resp)
}

// GetByID StaffTarif godoc
// @ID get_by_id_StaffTarif
// @Router /staffTarif/{id} [GET]
// @Summary Get By ID StaffTarif
// @Description Get By ID StaffTarif
// @Tags StaffTarif
// @Accept json
// @Procedure json
// @Param id path string true "id"
func (h *handler) GetByIdStaffTarif(c *gin.Context) {

	var id string = c.Param("id")
	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.StaffTarif().GetByID(c.Request.Context(), &models.StaffTarifPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.StaffTarif.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id StaffTarif resposne", http.StatusOK, resp)
}

// GetList StaffTarif godoc
// @ID get_list_StaffTarif
// @Router /staffTarif [GET]
// @Summary Get List StaffTarif
// @Description Get List StaffTarif
// @Tags StaffTarif
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListStaffTarif(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list StaffTarif offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list StaffTarif limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.StaffTarif().GetList(c.Request.Context(), &models.StaffTarifGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.StaffTarif.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list StaffTarif resposne", http.StatusOK, resp)
}

// Update StaffTarif godoc
// @ID update_StaffTarif
// @Router /staffTarif/{id} [PUT]
// @Summary Update StaffTarif
// @Description Update StaffTarif
// @Tags StaffTarif
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param StaffTarif body models.UpdateStaffTarif true "UpdateStaffTarifRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateStaffTarif(c *gin.Context) {

	var (
		id               string = c.Param("id")
		updateStaffTarif models.UpdateStaffTarif
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateStaffTarif)
	if err != nil {
		h.handlerResponse(c, "error StaffTarif should bind json", http.StatusBadRequest, err.Error())
		return
	}

	updateStaffTarif.Id = id
	rowsAffected, err := h.strg.StaffTarif().Update(c.Request.Context(), &updateStaffTarif)
	if err != nil {
		h.handlerResponse(c, "storage.StaffTarif.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.StaffTarif.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.StaffTarif().GetByID(c.Request.Context(), &models.StaffTarifPrimaryKey{Id: updateStaffTarif.Id})
	if err != nil {
		h.handlerResponse(c, "storage.StaffTarif.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create StaffTarif resposne", http.StatusAccepted, resp)
}

// Delete StaffTarif godoc
// @ID delete_StaffTarif
// @Router /staffTarif/{id} [DELETE]
// @Summary Delete StaffTarif
// @Description Delete StaffTarif
// @Tags StaffTarif
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteStaffTarif(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.StaffTarif().Delete(c.Request.Context(), &models.StaffTarifPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.StaffTarif.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create StaffTarif resposne", http.StatusNoContent, nil)
}
