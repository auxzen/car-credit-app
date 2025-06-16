package usecase

import (
	"ims-finance/model"
	"ims-finance/repository"
)

type KontrakUsecase interface {
	BuatKontrakDanJadwal(k model.Kontrak) error
}

type kontrakUsecaseImpl struct {
	kontrakRepo  repository.KontrakRepository
	angsuranRepo repository.AngsuranRepository
}

func NewKontrakUsecase(kRepo repository.KontrakRepository, aRepo repository.AngsuranRepository) KontrakUsecase {
	return &kontrakUsecaseImpl{
		kontrakRepo:  kRepo,
		angsuranRepo: aRepo,
	}
}

func (u *kontrakUsecaseImpl) BuatKontrakDanJadwal(k model.Kontrak) error {
	// Simpan kontrak
	err := u.kontrakRepo.Create(k)
	if err != nil {
		return err
	}

	// Hitung sisa pinjaman = OTR - DP
	sisaPinjaman := k.OTR - k.DownPayment

	var bunga float64
	switch {
	case k.TenorBulan <= 12:
		bunga = 0.12
	case k.TenorBulan <= 24:
		bunga = 0.14
	default:
		bunga = 0.165
	}

	// hitung angsuran per bulan berdasarakan poko + bunga
	totalBayar := sisaPinjaman + (sisaPinjaman * bunga)
	angsuranPerBulan := totalBayar / float64(k.TenorBulan)

	// Buat jadwal angsuran bulanan
	var jadwals []model.JadwalAngsuran
	for i := 1; i <= k.TenorBulan; i++ {
		jatuhTempo := k.TanggalMulai.AddDate(0, i, 0).AddDate(0, 0, 24-k.TanggalMulai.Day()) // Fix ke tgl 25
		jadwal := model.JadwalAngsuran{
			KontrakNo:         k.KontrakNo,
			AngsuranKe:        i,
			AngsuranPerBulan:  angsuranPerBulan,
			TanggalJatuhTempo: jatuhTempo,
		}
		jadwals = append(jadwals, jadwal)
	}

	// Simpan ke DB
	return u.angsuranRepo.InsertLayaway(jadwals)
}
