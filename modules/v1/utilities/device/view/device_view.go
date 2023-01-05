package view

import (
	"ClearningPatternGO/modules/v1/utilities/device/repository"
	"ClearningPatternGO/modules/v1/utilities/device/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type deviceView struct {
	productService service.Service
}

func NewDeviceView(productService service.Service) *deviceView {
	return &deviceView{productService}
}

func View(db *gorm.DB) *deviceView {
	Repository := repository.NewRepository(db)
	Service := service.NewService(Repository)
	View := NewDeviceView(Service)
	return View
}

func (h *deviceView) Index(c *gin.Context) {
	// listProduct, err := h.productService.ListProduct()
	listDevice, err := h.productService.ListDevice()
	fmt.Println(listDevice)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "device_index.html", gin.H{
		"list": listDevice,
	})
}
