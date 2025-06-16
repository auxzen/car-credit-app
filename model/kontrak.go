package model

import "time"

type Kontrak struct {
	KontrakNo    string    `json:"kontrak_no"`
	ClientName   string    `json:"client_name"`
	OTR          float64   `json:"otr"`
	DownPayment  float64   `json:"down_payment"`
	TenorBulan   int       `json:"tenor_bulan"`
	TanggalMulai time.Time `json:"tanggal_mulai"`
}
