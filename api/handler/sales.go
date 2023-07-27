package handler

import (
	"market/models"
	"market/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	Success              = "Success"
	Cancel               = "Cancel"
	PaymentType1         = "Card"
	PaymentType2         = "Cash"
	TypeStaffTransaction = "-"
	TarifType            = "PERCENT"
)

// Create Sales godoc
// @ID create_Sales
// @Router /sales [POST]
// @Summary Create Sales
// @Description Create Sales
// @Tags Sales
// @Accept json
// @Procedure json
// @Param Sales body models.CreateSales true "CreateSalesRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateSales(c *gin.Context) {

	var (
		createSales models.CreateSales
		amount      int
	)
	err := c.ShouldBindJSON(&createSales)
	if err != nil {
		h.handlerResponse(c, "error Sales should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Sales().Create(c.Request.Context(), &createSales)
	if err != nil {
		h.handlerResponse(c, "storage.Sales.create", http.StatusInternalServerError, err.Error())
		return
	}

	staff, err := h.strg.Staff().GetByID(c, &models.StaffPrimaryKey{Id: createSales.CashierId})
	if err != nil {
		h.handlerResponse(c, "storage.Sales.create.Staff.GetById", http.StatusInternalServerError, err.Error())
		return
	}
	tarif, err := h.strg.StaffTarif().GetByID(c, &models.StaffTarifPrimaryKey{Id: staff.TarifId})
	if err != nil {
		h.handlerResponse(c, "storage.Sales.create.StaffTarif.GetById", http.StatusInternalServerError, err.Error())
		return
	}

	if createSales.PaymentType == PaymentType1 {
		if tarif.Type == TarifType {
			amount = (createSales.Price * tarif.AmountForCard) / 100
		} else {
			amount = tarif.AmountForCard
		}
		_, err := h.strg.StaffTransaction().Create(c, &models.CreateStaffTransaction{
			SalesId: id,
			Type:    "+",
			Text:    "success",
			Amount:  amount,
			StaffId: createSales.CashierId,
		})
		if err != nil {
			h.handlerResponse(c, "storage.Sales.Create.StaffTransaction.create", http.StatusInternalServerError, err.Error())
			return
		}
		_, err = h.strg.Staff().Update(c, &models.UpdateStaff{
			Id:      createSales.CashierId,
			Name:    staff.Name,
			Type:    staff.Type,
			Balance: staff.Balance + amount,
		})
		if err != nil {
			h.handlerResponse(c, "storage.Sales.Create.Staff.update", http.StatusInternalServerError, err.Error())
			return
		}
	} else if createSales.PaymentType == PaymentType2 {
		if tarif.Type == TarifType {
			amount = (createSales.Price * tarif.AmountForCash) / 100
		} else {
			amount = tarif.AmountForCash
		}
		_, err := h.strg.StaffTransaction().Create(c, &models.CreateStaffTransaction{
			SalesId: id,
			Type:    "+",
			Text:    "success",
			Amount:  amount,
			StaffId: createSales.CashierId,
		})
		if err != nil {
			h.handlerResponse(c, "storage.Sales.Create.StaffTransaction.create", http.StatusInternalServerError, err.Error())
			return
		}
		_, err = h.strg.Staff().Update(c, &models.UpdateStaff{
			Id:      createSales.CashierId,
			Name:    staff.Name,
			Type:    staff.Type,
			Balance: staff.Balance + amount,
		})
		if err != nil {
			h.handlerResponse(c, "storage.Sales.Create.Staff.update", http.StatusInternalServerError, err.Error())
			return
		}
	}

	if createSales.AsistentId != "" {
		staff, err := h.strg.Staff().GetByID(c, &models.StaffPrimaryKey{Id: createSales.AsistentId})
		if err != nil {
			h.handlerResponse(c, "storage.Sales.create.Staff.GetById", http.StatusInternalServerError, err.Error())
			return
		}
		tarif, err := h.strg.StaffTarif().GetByID(c, &models.StaffTarifPrimaryKey{Id: staff.TarifId})
		if err != nil {
			h.handlerResponse(c, "storage.Sales.create.StaffTarif.GetById", http.StatusInternalServerError, err.Error())
			return
		}
		if createSales.PaymentType == PaymentType1 {
			if tarif.Type == TarifType {
				amount = (createSales.Price * tarif.AmountForCard) / 100
			} else {
				amount = tarif.AmountForCard
			}
			_, err := h.strg.StaffTransaction().Create(c, &models.CreateStaffTransaction{
				SalesId: id,
				Type:    "+",
				Text:    "success",
				Amount:  amount,
				StaffId: createSales.AsistentId,
			})
			if err != nil {
				h.handlerResponse(c, "storage.Sales.Create.StaffTransaction.create", http.StatusInternalServerError, err.Error())
				return
			}
			_, err = h.strg.Staff().Update(c, &models.UpdateStaff{
				Id:      createSales.AsistentId,
				Name:    staff.Name,
				Type:    staff.Type,
				Balance: staff.Balance + amount,
			})
			if err != nil {
				h.handlerResponse(c, "storage.Sales.Create.Staff.update", http.StatusInternalServerError, err.Error())
				return
			}
		} else if createSales.PaymentType == PaymentType2 {
			if tarif.Type == TarifType {
				amount = (createSales.Price * tarif.AmountForCash) / 100
			} else {
				amount = tarif.AmountForCash
			}
			_, err := h.strg.StaffTransaction().Create(c, &models.CreateStaffTransaction{
				SalesId: id,
				Type:    "+",
				Text:    "success",
				Amount:  amount,
				StaffId: createSales.AsistentId,
			})
			if err != nil {
				h.handlerResponse(c, "storage.Sales.Create.StaffTransaction.create", http.StatusInternalServerError, err.Error())
				return
			}
			_, err = h.strg.Staff().Update(c, &models.UpdateStaff{
				Id:      createSales.AsistentId,
				Name:    staff.Name,
				Type:    staff.Type,
				Balance: staff.Balance + amount,
			})
			if err != nil {
				h.handlerResponse(c, "storage.Sales.Create.Staff.update", http.StatusInternalServerError, err.Error())
				return
			}
		}
	}

	resp, err := h.strg.Sales().GetByID(c.Request.Context(), &models.SalesPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Sales.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Sales resposne", http.StatusCreated, resp)
}

// GetByID Sales godoc
// @ID get_by_id_Sales
// @Router /sales/{id} [GET]
// @Summary Get By ID Sales
// @Description Get By ID Sales
// @Tags Sales
// @Accept json
// @Procedure json
// @Param id path string true "id"
func (h *handler) GetByIdSales(c *gin.Context) {

	var id string = c.Param("id")
	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Sales().GetByID(c.Request.Context(), &models.SalesPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Sales.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id Sales resposne", http.StatusOK, resp)
}

// GetList Sales godoc
// @ID get_list_Sales
// @Router /sales [GET]
// @Summary Get List Sales
// @Description Get List Sales
// @Tags Sales
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Param branchid query string false "branchID"
// @Param assistentid query string false "asistentID"
// @Param paymenttype query string false "paymentType"
// @Param cashierid query string false "cashierID"
// @Param status query string false "status"
// @Param price query string false "price"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListSales(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list Sales offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list Sales limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Sales().GetList(c.Request.Context(), &models.SalesGetListRequest{
		Offset:             offset,
		Limit:              limit,
		SearchByClientName: c.Query("search"),
		SearchByBranchId:   c.Query("branchID"),
		ShopAsistentId:     c.Query("asistentID"),
		PaymentType:        c.Query("paymentType"),
		CashierId:          c.Query("cashierID"),
		Status:             c.Query("status"),
		Price:              c.GetInt("price"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.Sales.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list Sales resposne", http.StatusOK, resp)
}

// Update Sales godoc
// @ID update_Sales
// @Router /sales/{id} [PUT]
// @Summary Update Sales
// @Description Update Sales
// @Tags Sales
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param Sales body models.UpdateSales true "UpdateSalesRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateSales(c *gin.Context) {

	var (
		id          string = c.Param("id")
		updateSales models.UpdateSales
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateSales)
	if err != nil {
		h.handlerResponse(c, "error Sales should bind json", http.StatusBadRequest, err.Error())
		return
	}

	if updateSales.Status == Cancel {
		transactions, err := h.strg.StaffTransaction().GetBySalesID(c, &models.PrimaryKey{SalesId: updateSales.Id})
		if err != nil {
			h.handlerResponse(c, "storage.Sales.update.StaffTransaction.GetBySalesID", http.StatusInternalServerError, err.Error())
			return
		}
		for i := 0; i < transactions.Count; i++ {
			_, err = h.strg.StaffTransaction().Update(c, &models.UpdateStaffTransaction{
				Id:      transactions.StaffTransaction[i].Id,
				Type:    TypeStaffTransaction,
				Text:    Cancel,
				Amount:  transactions.StaffTransaction[i].Amount,
				StaffId: transactions.StaffTransaction[i].StaffId,
			})
			if err != nil {
				h.handlerResponse(c, "storage.Sales.update.StaffTransaction.Update", http.StatusInternalServerError, err.Error())
				return
			}
			staff, err := h.strg.Staff().GetByID(c, &models.StaffPrimaryKey{Id: transactions.StaffTransaction[i].StaffId})
			if err != nil {
				h.handlerResponse(c, "storage.Sales.update.Staff.GeByID", http.StatusInternalServerError, err.Error())
				return
			}
			_, err = h.strg.Staff().Update(c, &models.UpdateStaff{
				Id:      staff.Id,
				Name:    staff.Name,
				Type:    staff.Type,
				Balance: staff.Balance - transactions.StaffTransaction[i].Amount,
			})
			if err != nil {
				h.handlerResponse(c, "storage.Sales.update.Staff.Update", http.StatusInternalServerError, err.Error())
				return
			}
		}
	}

	updateSales.Id = id
	rowsAffected, err := h.strg.Sales().Update(c.Request.Context(), &updateSales)
	if err != nil {
		h.handlerResponse(c, "storage.Sales.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.Sales.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Sales().GetByID(c.Request.Context(), &models.SalesPrimaryKey{Id: updateSales.Id})
	if err != nil {
		h.handlerResponse(c, "storage.Sales.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Sales resposne", http.StatusAccepted, resp)
}

// Delete Sales godoc
// @ID delete_Sales
// @Router /sales/{id} [DELETE]
// @Summary Delete Sales
// @Description Delete Sales
// @Tags Sales
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteSales(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Sales().Delete(c.Request.Context(), &models.SalesPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Sales.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Sales resposne", http.StatusNoContent, nil)
}

// Sort Sales godoc
// @ID sort_Sales
// @Router /sortsales [GET]
// @Summary Sort Sales
// @Description Sort Sales
// @Tags Sales
// @Accept json
// @Procedure json
// @Param from query string true "from"
// @Param to query string true "to"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) SortSales(c *gin.Context) {

	resp, err := h.strg.Sales().SortBySalesAmount(c, &models.SalesSortRequest{
		CreatedAtFrom: c.Query("from"),
		CreatedAtTo:   c.Query("to"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.Sales.sortSales", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "sort sales", http.StatusAccepted, resp)
}
