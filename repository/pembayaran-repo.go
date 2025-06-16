package repository

import (
	"context"
	"database/sql"
	"ims-finance/model"
)

type PembayaranRepository interface {
	InsertPayment(payment model.Pembayaran) error
	ListByKontrak(kontrakNo string) ([]model.Pembayaran, error)
}

type pembayaranRepository struct {
	db *sql.DB
}

func (r *pembayaranRepository) InsertPayment(payment model.Pembayaran) error {
	query := `
		INSERT INTO pembayaran(kontrak_no, jumlah, tanggal)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.ExecContext(context.Background(), query,
		payment.KontrakNo,
		payment.AngsuranKe,
		payment.TanggalBayar,
	)

	return err
}

func (r *pembayaranRepository) ListByKontrak(kontrakNo string) ([]model.Pembayaran, error) {
	query := `
		SELECT id, kontrak_no, angsuran_ke, tanggal_bayar
		FROM pembayaran
		WHERE kontrak_no = $1
	`

	rows, err := r.db.QueryContext(context.Background(), query, kontrakNo)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var list []model.Pembayaran
	for rows.Next() {
		var p model.Pembayaran
		err := rows.Scan(&p.ID, &p.KontrakNo, &p.AngsuranKe, &p.TanggalBayar)
		if err != nil {
			return nil, err
		}
		list = append(list, p)
	}

	return list, nil
}

func NewPembayaranRepository(db *sql.DB) PembayaranRepository {
	return &pembayaranRepository{db: db}
}
