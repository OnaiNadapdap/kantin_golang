package barang

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/onainadapdap1/golang_kantin/internal/api"
	"github.com/onainadapdap1/golang_kantin/internal/service/barang"
	"github.com/onainadapdap1/golang_kantin/models"
)

type BarangHandler interface {
	CreateBarang(c *gin.Context)
	ShowBarang(c *gin.Context)
	HideBarang(c *gin.Context)
	GetPengumuman(c *gin.Context)
	// GetAllBarang(c *gin.Context)
}

type barangHandler struct {
	serv barang.BarangService
}

func NewBarangHandler(serv barang.BarangService) BarangHandler {
	return &barangHandler{serv: serv}
}

func (h *barangHandler) CreateBarang(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(models.User)
	var barangInput api.CreateInputBarang
	if err := c.ShouldBind(&barangInput); err != nil {
		// {"error brewuu":"json: cannot unmarshal string into Go struct field CreateInputBarang.user_id of type uint"}
		c.JSON(http.StatusBadRequest, gin.H{"error brewuu": err.Error()})
		return
	}
	log.Println("barang input : ", barangInput)
	log.Println("tipe user id : ", reflect.TypeOf(currentUser.ID))

	parsedTime, err := time.Parse("2006-01-02", barangInput.ExpiryDate)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	barang := models.Barang{
		Kategori:    barangInput.Kategori,
		UserID:      currentUser.ID,
		Name:        barangInput.Name,
		Description: barangInput.Description,
		ExpiryDate:  parsedTime,
		File:        barangInput.File,
	}

	if err := h.serv.CreateBarang(&barang); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send barang", "err": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "barang sended successfully"})

}

func (h *barangHandler) ShowBarang(c *gin.Context) {
	// id := c.Params("id")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	fmt.Println("hello world")
	log.Println("id : ", id)
	err = h.serv.ShowBarang(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "can't show barang!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success to show barang"})
}

func (h *barangHandler) HideBarang(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid om"})
		return
	}

	err = h.serv.HideBarang(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "can't hide barang!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success to hide barang"})
}

func (h *barangHandler) GetPengumuman(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("perPage", "10"))
	fmt.Println("logging here handler")
	pengumuman, err := h.serv.GetPengumuman(page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch pengumuman"})
		return
	}

	// c.JSON(http.StatusOK, pengumuman)
	c.JSON(http.StatusOK, gin.H{"success": true, "data": pengumuman})
}
