package menumakanan

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/onainadapdap1/golang_kantin/internal/api"
	"github.com/onainadapdap1/golang_kantin/internal/service/menumakanan"
	"github.com/onainadapdap1/golang_kantin/models"
)

type MenuMakananHandler interface {
	CreateMenuMakanan(c *gin.Context) 
	GetAllMenuMakanan(c *gin.Context)
}

type menuMakananHandler struct {
	serv menumakanan.MenuMakananService
}

func NewMenuMakananHandler(serv menumakanan.MenuMakananService) MenuMakananHandler {
	return &menuMakananHandler{serv: serv}
}

func (h *menuMakananHandler) CreateMenuMakanan(c *gin.Context)  {
	var menuMakanInput api.CreateMenuMakananInput
	if err := c.ShouldBind(&menuMakanInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error brewuu": err.Error()})
		return
	}

	parsedTime, err := time.Parse("2006-01-02", menuMakanInput.TanggalMakan)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	menuMakan := models.MenuMakanan {
		TanggalMakan: parsedTime,
		MenuPagi: menuMakanInput.MenuPagi,
		MenuSiang: menuMakanInput.MenuSiang,
		MenuMalam: menuMakanInput.MenuMalam,
		Foto1: menuMakanInput.Foto1,
		Foto2: menuMakanInput.Foto2,
		Foto3: menuMakanInput.Foto3,
	}

	if err := h.serv.CreateMenuMakanan(menuMakan); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create menu makanan"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": "Successfully to create menu makanan"})
}

func (h *menuMakananHandler) GetAllMenuMakanan(c *gin.Context) {
	menus, err := h.serv.GetAllMenuMakanan()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no data is found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": menus,
	});
}