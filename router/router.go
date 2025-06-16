package router

import (
	"ims-finance/handler"
	"ims-finance/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	kontrakHandler *handler.KontrakHandler,
	pembayaranHandler *handler.PembayaranHandler,
	dendaHandler *handler.DendaHandler,
) *gin.Engine {
	r := gin.Default()

	// Gunakan middleware
	r.Use(middleware.LoggerMiddleware())
	r.Use(gin.Recovery()) // atau CustomRecovery

	// Endpoint Kontrak
	r.POST("/kontrak", kontrakHandler.BuatKontrak)

	// Endpoint Pembayaran
	r.POST("/pembayaran", pembayaranHandler.CatatPembayaran)

	// Endpoint Denda
	r.GET("/kontrak/:kontrak_no/denda", dendaHandler.CekDenda)

	return r
}
