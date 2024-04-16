package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/onainadapdap1/golang_kantin/config"
	pengumumanHandler "github.com/onainadapdap1/golang_kantin/internal/handler/pengumuman"
	pengumumanRepo "github.com/onainadapdap1/golang_kantin/internal/repository/pengumuman"
	"github.com/onainadapdap1/golang_kantin/internal/service/auth"
	pengumumanServ "github.com/onainadapdap1/golang_kantin/internal/service/pengumuman"

	feedbackHandler "github.com/onainadapdap1/golang_kantin/internal/handler/feedback"
	feedbackRepo "github.com/onainadapdap1/golang_kantin/internal/repository/feedback"
	feedbackServ "github.com/onainadapdap1/golang_kantin/internal/service/feedback"

	barangHandler "github.com/onainadapdap1/golang_kantin/internal/handler/barang"
	barangRepo "github.com/onainadapdap1/golang_kantin/internal/repository/barang"
	barangServ "github.com/onainadapdap1/golang_kantin/internal/service/barang"

	menuMakanHandler "github.com/onainadapdap1/golang_kantin/internal/handler/menumakanan"
	menuMakanRepo "github.com/onainadapdap1/golang_kantin/internal/repository/menumakanan"
	menuMakanServ "github.com/onainadapdap1/golang_kantin/internal/service/menumakanan"

	allergyReportHandler "github.com/onainadapdap1/golang_kantin/internal/handler/allergyreport"
	allergyReportRepo "github.com/onainadapdap1/golang_kantin/internal/repository/allergyreport"
	allergyReportServ "github.com/onainadapdap1/golang_kantin/internal/service/allergyreport"

	userHandler "github.com/onainadapdap1/golang_kantin/internal/handler/user"
	userRepo "github.com/onainadapdap1/golang_kantin/internal/repository/user"
	userServ "github.com/onainadapdap1/golang_kantin/internal/service/user"

	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	// config.LoadEnv()
	DB = config.ConnectToDB()
}

func main() {
	authService := auth.NewAuthService()
	userRepo := userRepo.NewUserRepository(DB)
	userServ := userServ.NewUserService(userRepo)
	userHandler := userHandler.NewUserHandler(userServ, authService)

	pengumumanRepo := pengumumanRepo.NewPengumumanRepository(DB)
	pengumumanServ := pengumumanServ.NewPengumumanService(pengumumanRepo)
	pengumumanHandler := pengumumanHandler.NewPengumumanHandler(pengumumanServ)

	feedbackRepo := feedbackRepo.NewFeedbackRepository(DB)
	feedbackServ := feedbackServ.NewFeedbackService(feedbackRepo)
	feedbackHandler := feedbackHandler.NewFeedbackHandler(feedbackServ)

	barangRepo := barangRepo.NewBarangRepository(DB)
	barangServ := barangServ.NewBarangService(barangRepo)
	barangHandler := barangHandler.NewBarangHandler(barangServ)

	menuMakanRepo := menuMakanRepo.NewMenuMakananRepo(DB)
	menuMakanServ := menuMakanServ.NewMenuMakananServ(menuMakanRepo)
	menuMakanHandler := menuMakanHandler.NewMenuMakananHandler(menuMakanServ)

	allergyReportRepo := allergyReportRepo.NewAllergyReportRepo(DB)
	allergyReportServ := allergyReportServ.NewAllergyReportServ(allergyReportRepo)
	allergyReportHandler := allergyReportHandler.NewAllergyReportHandler(allergyReportServ)

	router := gin.Default()
	api := router.Group("/api/v1")
	// Membuat pengumuman baru
	// newPengumuman := models.Pengumuman{
	//     TanggalBerakhir:  time.Now(),
	//     TanggalPembuatan: time.Now(),
	//     Deskripsi:        "Contoh pengumuman",
	//     Published:        true,
	// }

	// // Menyimpan pengumuman ke dalam database
	// DB.Create(&newPengumuman)

	api.POST("/login", userHandler.Login)
	api.POST("/pengumuman", pengumumanHandler.CreatePengumuman)
	api.GET("/pengumuman", pengumumanHandler.GetAllPengumuman)
	api.POST("/feedback", feedbackHandler.CreateFeedback)
	api.GET("/feedback", feedbackHandler.GetAllMyFeedback)
	api.POST("/barangs", barangHandler.CreateBarang)
	api.GET("/show-barangs/:id", barangHandler.ShowBarang)
	api.GET("/hide-barangs/:id", barangHandler.HideBarang)
	api.GET("/barangs", barangHandler.GetPengumuman) //pagination
	api.POST("/menu-makanans", menuMakanHandler.CreateMenuMakanan)
	api.GET("/menu-makanans", menuMakanHandler.GetAllMenuMakanan)
	api.POST("/allergy-reports", allergyReportHandler.CreateAllergyReport)

	router.Run(envPortOr("8000"))
}

func envPortOr(port string) string {
	envPort := os.Getenv("PORT")
	if envPort != "" {
		return ":" + envPort
	}
	return ":" + port
}
