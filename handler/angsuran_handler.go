package handler

import (
	"ims-finance/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type DendaHandler struct {
	usecase usecase.DendaUsecase
}

func NewDendaHandler(u usecase.DendaUsecase) *DendaHandler {
	return &DendaHandler{usecase: u}
}

func (h *DendaHandler) CekDenda(c *gin.Context) {
	kontrakNo := c.Param("kontrak_no")

	// Bisa ambil dari query params atau default ke 14 Agustus 2024
	tanggalStr := c.Query("tanggal")
	var tanggal time.Time
	var err error
	if tanggalStr == "" {
		tanggal, _ = time.Parse("2006-01-02", "2024-08-14")
	} else {
		tanggal, err = time.Parse("2006-01-02", tanggalStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "format tanggal harus YYYY-MM-DD"})
			return
		}
	}

	dendaList, err := h.usecase.HitungDenda(kontrakNo, tanggal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal menghitung denda"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"kontrak_no": kontrakNo, "denda": dendaList})
}
