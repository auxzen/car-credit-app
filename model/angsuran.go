package model

import "time"

type JadwalAngsuran struct {
	ID                int       `json:"id"`
	KontrakNo         string    `json:"kontrak_no"`
	AngsuranKe        int       `json:"angsuran_ke"`
	AngsuranPerBulan  float64   `json:"angsuran_per_bulan"`
	TanggalJatuhTempo time.Time `json:"tanggal_jatuh_tempo"`
}
