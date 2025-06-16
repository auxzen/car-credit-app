package repository

import (
	"context"
	"database/sql"
	"ims-finance/model"
)

type KontrakRepository interface {
	Create(kontrak model.Kontrak) error
	GetByNo(KontrackNo string) (*model.Kontrak, error)
}

type kontrakRepository struct {
	db *sql.DB
}

func (r *kontrakRepository) Create(kontrak model.Kontrak) error {
	query := `
		INSERT INTO kontrak (kontrak_no, 
			client_name, 
			otr, 
			down_payment, 
			tenor_bulan, 
			tanggal_mulai)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.ExecContext(context.Background(), query,
		kontrak.KontrakNo,
		kontrak.ClientName,
		kontrak.OTR,
		kontrak.DownPayment,
		kontrak.TenorBulan,
		kontrak.TanggalMulai,
	)

	return err
}

func (r *kontrakRepository) GetByNo(KontrackNo string) (*model.Kontrak, error) {
	query := `
		SELECT kontrak_no, client_name, otr, down_payment, tenor_bulan, tanggal mulai
		FROM kontrak
		WHERE kontrak_no = $1
	`

	row := r.db.QueryRowContext(context.Background(), query, KontrackNo)

	var k model.Kontrak
	err := row.Scan(&k.KontrakNo,
		&k.ClientName,
		&k.OTR,
		&k.DownPayment,
		&k.TenorBulan,
		&k.TanggalMulai,
	)
	if err != nil {
		return nil, err
	}

	return &k, nil
}

func NewKontrakRepository(db *sql.DB) KontrakRepository {
	return &kontrakRepository{db: db}
}
