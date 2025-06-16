package handler

import (
	"ims-finance/model"
	"ims-finance/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type KontrakHandler struct {
	usecase usecase.KontrakUsecase
}

func NewKontrakHandler(u usecase.KontrakUsecase) *KontrakHandler {
	return &KontrakHandler{usecase: u}
}

func (h *KontrakHandler) BuatKontrak(c *gin.Context) {
	var input model.Kontrak
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Default tanggal mulai = hari ini jika tidak diset
	if input.TanggalMulai.IsZero() {
		input.TanggalMulai = time.Now()
	}

	err := h.usecase.BuatKontrakDanJadwal(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal menyimpan kontrak"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Kontrak berhasil dibuat", "kontrak_no": input.KontrakNo})
}
