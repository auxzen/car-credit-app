package model

import "time"

type Pembayaran struct {
	ID           int       `json:"id"`
	KontrakNo    string    `json:"kontrak_no"`
	AngsuranKe   int       `json:"angsuran_ke"`
	TanggalBayar time.Time `json:"tanggal_bayar" time_format:"2006-01-02"`
}
