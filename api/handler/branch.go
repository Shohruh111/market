package handler

import (
	"market/models"
	"market/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Branch godoc
// @ID create_branch
// @Router /branch [POST]
// @Summary Create branch
// @Description Create Branch
// @Tags Branch
// @Accept json
// @Procedure json
// @Param Branch body models.CreateBranch true "CreateBranchRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateBranch(c *gin.Context) {

	var createBranch models.CreateBranch
	err := c.ShouldBindJSON(&createBranch)
	if err != nil {
		h.handlerResponse(c, "error Branch should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Branch().Create(c.Request.Context(), &createBranch)
	if err != nil {
		h.handlerResponse(c, "storage.Branch.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.Branch().GetByID(c.Request.Context(), &models.BranchPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Branch.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Branch resposne", http.StatusCreated, resp)
}

// GetByID Branch godoc
// @ID get_by_id_Branch
// @Router /branch/{id} [GET]
// @Summary Get By ID Branch
// @Description Get By ID Branch
// @Tags Branch
// @Accept json
// @Procedure json
// @Param id path string true "id"
func (h *handler) GetByIdBranch(c *gin.Context) {

	var id string = c.Param("id")
	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Branch().GetByID(c.Request.Context(), &models.BranchPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Branch.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id Branch resposne", http.StatusOK, resp)
}

// GetList Branch godoc
// @ID get_list_Branch
// @Router /branch [GET]
// @Summary Get List Branch
// @Description Get List Branch
// @Tags Branch
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Param address query string false "address"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListBranch(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list Branch offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list Branch limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Branch().GetList(c.Request.Context(), &models.BranchGetListRequest{
		Offset: offset,
		Limit:  limit,
		SearchByName: c.Query("search"),
		SearchByAddress: c.Query("address"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.Branch.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list Branch resposne", http.StatusOK, resp)
}

// Update Branch godoc
// @ID update_Branch
// @Router /branch/{id} [PUT]
// @Summary Update Branch
// @Description Update Branch
// @Tags Branch
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param Branch body models.UpdateBranch true "UpdateBranchRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateBranch(c *gin.Context) {

	var (
		id           string = c.Param("id")
		updateBranch models.UpdateBranch
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateBranch)
	if err != nil {
		h.handlerResponse(c, "error Branch should bind json", http.StatusBadRequest, err.Error())
		return
	}

	updateBranch.Id = id
	rowsAffected, err := h.strg.Branch().Update(c.Request.Context(), &updateBranch)
	if err != nil {
		h.handlerResponse(c, "storage.Branch.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.Branch.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Branch().GetByID(c.Request.Context(), &models.BranchPrimaryKey{Id: updateBranch.Id})
	if err != nil {
		h.handlerResponse(c, "storage.Branch.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Branch resposne", http.StatusAccepted, resp)
}

// Delete Branch godoc
// @ID delete_Branch
// @Router /branch/{id} [DELETE]
// @Summary Delete Branch
// @Description Delete Branch
// @Tags Branch
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteBranch(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Branch().Delete(c.Request.Context(), &models.BranchPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Branch.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Branch resposne", http.StatusNoContent, nil)
}
