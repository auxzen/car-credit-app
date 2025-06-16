package usecase

import (
	"ims-finance/model"
	"ims-finance/repository"
	"time"
)

type PembayaranUsecase interface {
	CatatPembayaran(p model.Pembayaran) error
}

type pembayaranUsecaseImpl struct {
	pembayaranRepo repository.PembayaranRepository
}

func NewPembayaranUsecase(pRepo repository.PembayaranRepository) PembayaranUsecase {
	return &pembayaranUsecaseImpl{
		pembayaranRepo: pRepo,
	}
}

func (u *pembayaranUsecaseImpl) CatatPembayaran(p model.Pembayaran) error {
	if p.TanggalBayar.IsZero() {
		p.TanggalBayar = time.Now()
	}
	return u.pembayaranRepo.InsertPayment(p)
}
