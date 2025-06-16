package handler

import (
	"ims-finance/model"
	"ims-finance/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PembayaranHandler struct {
	usecase usecase.PembayaranUsecase
}

func NewPembayaranHandler(u usecase.PembayaranUsecase) *PembayaranHandler {
	return &PembayaranHandler{usecase: u}
}

func (h *PembayaranHandler) CatatPembayaran(c *gin.Context) {
	var input model.Pembayaran
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.usecase.CatatPembayaran(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mencatat pembayaran"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pembayaran berhasil dicatat"})
}
