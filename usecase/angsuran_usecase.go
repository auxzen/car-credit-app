package usecase

import (
	"ims-finance/repository"
	"time"
)

type DendaDetail struct {
	AngsuranKe        int
	HariKeterlambatan int
	TotalDenda        float64
}

type DendaUsecase interface {
	HitungDenda(kontrakNo string, sampaiTanggal time.Time) ([]DendaDetail, error)
}

type dendaUsecaseImpl struct {
	angsuranRepo   repository.AngsuranRepository
	pembayaranRepo repository.PembayaranRepository
}

func NewDendaUsecase(aRepo repository.AngsuranRepository, pRepo repository.PembayaranRepository) DendaUsecase {
	return &dendaUsecaseImpl{
		angsuranRepo:   aRepo,
		pembayaranRepo: pRepo,
	}
}

func (u *dendaUsecaseImpl) HitungDenda(kontrakNo string, sampaiTanggal time.Time) ([]DendaDetail, error) {
	angsuranList, err := u.angsuranRepo.ListByKontrak(kontrakNo)
	if err != nil {
		return nil, err
	}
	pembayaranList, err := u.pembayaranRepo.ListByKontrak(kontrakNo)
	if err != nil {
		return nil, err
	}

	paidMap := map[int]bool{}
	for _, p := range pembayaranList {
		paidMap[p.AngsuranKe] = true
	}

	var result []DendaDetail

	for _, a := range angsuranList {
		if a.TanggalJatuhTempo.After(sampaiTanggal) {
			continue
		}
		if paidMap[a.AngsuranKe] {
			continue
		}

		delay := int(sampaiTanggal.Sub(a.TanggalJatuhTempo).Hours() / 24)
		denda := float64(delay) * 0.001 * a.AngsuranPerBulan

		result = append(result, DendaDetail{
			AngsuranKe:        a.AngsuranKe,
			HariKeterlambatan: delay,
			TotalDenda:        denda,
		})
	}

	return result, nil
}
