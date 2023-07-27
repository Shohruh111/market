package api

import (
	_ "market/api/docs"
	"market/api/handler"
	"market/config"
	"market/pkg/logger"
	"market/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewApi(r *gin.Engine, cfg *config.Config, strg storage.StorageI, logger logger.LoggerI) {
	handler := handler.NewHandler(cfg, strg, logger)

	r.GET("/branch", handler.GetListBranch)
	r.GET("/branch/:id", handler.GetByIdBranch)
	r.POST("branch", handler.CreateBranch)
	r.PUT("/branch/:id", handler.UpdateBranch)
	r.DELETE("/branch/:id", handler.DeleteBranch)

	r.GET("/sales", handler.GetListSales)
	r.GET("/sortsales", handler.SortSales)
	r.GET("/sales/:id", handler.GetByIdSales)
	r.POST("sales", handler.CreateSales)
	r.PUT("/sales/:id", handler.UpdateSales)
	r.DELETE("/sales/:id", handler.DeleteSales)

	r.GET("/staff", handler.GetListStaff)
	r.GET("/topstaff", handler.GetListTop)
	r.GET("/staff/:id", handler.GetByIdStaff)
	r.POST("staff", handler.CreateStaff)
	r.PUT("/staff/:id", handler.UpdateStaff)
	r.DELETE("/staff/:id", handler.DeleteStaff)

	r.GET("/staffTarif", handler.GetListStaffTarif)
	r.GET("/staffTarif/:id", handler.GetByIdStaffTarif)
	r.POST("staffTarif", handler.CreateStaffTarif)
	r.PUT("/staffTarif/:id", handler.UpdateStaffTarif)
	r.DELETE("/staffTarif/:id", handler.DeleteStaffTarif)

	r.GET("/staffTransaction", handler.GetListStaffTransaction)
	r.GET("/staffTransaction/:id", handler.GetByIdStaffTransaction)
	r.POST("staffTransaction", handler.CreateStaffTransaction)
	r.PUT("/staffTransaction/:id", handler.UpdateStaffTransaction)
	r.DELETE("/staffTransaction/:id", handler.DeleteStaffTransaction)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
