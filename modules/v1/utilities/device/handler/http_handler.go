package product

import (
	"ClearningPatternGO/modules/v1/utilities/device/repository"
	"ClearningPatternGO/modules/v1/utilities/device/service"
	api "ClearningPatternGO/pkg/api_response"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DeviceHandler interface {
	ListProduct(c *gin.Context)
}

type deviceHandler struct {
	productService service.Service
}

func NewDeviceHandler(productService service.Service) *deviceHandler {
	return &deviceHandler{productService}
}

func Handler(db *gorm.DB) *deviceHandler {
	Repository := repository.NewRepository(db)
	Service := service.NewService(Repository)
	Handler := NewDeviceHandler(Service)
	return Handler
}

func (h *deviceHandler) ListDevice(c *gin.Context) {
	listProduct, err := h.productService.ListDevice()

	if err != nil {
		response := api.APIRespon("Failed to get data all device", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := api.APIRespon("Success to get data all device", http.StatusOK, "success", listProduct)
	c.JSON(http.StatusOK, response)
}
