package main

import (
	"ims-finance/config"
	"ims-finance/handler"
	"ims-finance/repository"
	"ims-finance/router"
	"ims-finance/usecase"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(".env file tidak ditemukan, pakai default environment")
	}

	//  koneksi DB
	config.ConnectDB()
	db := config.ConnectDB()

	//  repository
	kontrakRepo := repository.NewKontrakRepository(db)
	angsuranRepo := repository.NewAngsuranRepository(db)
	pembayaranRepo := repository.NewPembayaranRepository(db)

	//  usecase
	kontrakUsecase := usecase.NewKontrakUsecase(kontrakRepo, angsuranRepo)
	pembayaranUsecase := usecase.NewPembayaranUsecase(pembayaranRepo)
	dendaUsecase := usecase.NewDendaUsecase(angsuranRepo, pembayaranRepo)

	//  handler
	kontrakHandler := handler.NewKontrakHandler(kontrakUsecase)
	pembayaranHandler := handler.NewPembayaranHandler(pembayaranUsecase)
	dendaHandler := handler.NewDendaHandler(dendaUsecase)

	// Routing
	r := router.SetupRouter(kontrakHandler, pembayaranHandler, dendaHandler)

	// Jalankan server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
