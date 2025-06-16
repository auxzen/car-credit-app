package repository

import (
	"context"
	"database/sql"
	"ims-finance/model"
)

type AngsuranRepository interface {
	InsertLayaway(jadwal []model.JadwalAngsuran) error
	ListByKontrak(KontrakNo string) ([]model.JadwalAngsuran, error)
}

type angsuranRepository struct {
	db *sql.DB
}

func (r *angsuranRepository) InsertLayaway(jadwal []model.JadwalAngsuran) error {
	query := `
		INSERT INTO jadwal_angsuran(kontrak_no, 
			angsuran_ke, 
			angsuran_per_bulan, 
			tanggal_jatuh_tempo)
		VALUES ($1,$2,$3,$4)`

	tx, err := r.db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	for _, j := range jadwal {
		_, err := tx.ExecContext(context.Background(), query, j.KontrakNo, j.AngsuranKe, j.AngsuranPerBulan, j.TanggalJatuhTempo)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (r *angsuranRepository) ListByKontrak(KontrakNo string) ([]model.JadwalAngsuran, error) {
	query := `SELECT id, kontrak_no, angsuran_ke, angsuran_per_bulan, tanggal_jatuh_tempo 
	FROM jadwal_angsuran 
	WHERE kontrak_no = $1 
	ORDER BY angsuran_ke`

	rows, err := r.db.QueryContext(context.Background(), query, KontrakNo)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var list []model.JadwalAngsuran
	for rows.Next() {
		var j model.JadwalAngsuran
		err := rows.Scan(&j.ID, &j.KontrakNo, &j.AngsuranKe, &j.AngsuranPerBulan, &j.TanggalJatuhTempo)
		if err != nil {
			return nil, err
		}
		list = append(list, j)
	}

	return list, nil
}

func NewAngsuranRepository(db *sql.DB) AngsuranRepository {
	return &angsuranRepository{db: db}
}
