-- Tabel Kontrak Kredit
CREATE TABLE kontrak (
    kontrak_no VARCHAR(20) PRIMARY KEY,
    client_name VARCHAR(100) NOT NULL,
    otr NUMERIC(15, 2) NOT NULL, -- Harga mobil
    down_payment NUMERIC(15, 2) NOT NULL,
    tenor_bulan INT NOT NULL, -- Lama cicilan dalam bulan
    tanggal_mulai DATE NOT NULL
);

-- Tabel Jadwal Angsuran
CREATE TABLE jadwal_angsuran (
    id SERIAL PRIMARY KEY,
    kontrak_no VARCHAR(20) REFERENCES kontrak(kontrak_no) ON DELETE CASCADE,
    angsuran_ke INT NOT NULL,
    angsuran_per_bulan NUMERIC(15, 2) NOT NULL,
    tanggal_jatuh_tempo DATE NOT NULL
);

-- Tabel Pembayaran
CREATE TABLE pembayaran (
    id SERIAL PRIMARY KEY,
    kontrak_no VARCHAR(20) REFERENCES kontrak(kontrak_no) ON DELETE CASCADE,
    angsuran_ke INT NOT NULL,
    tanggal_bayar DATE NOT NULL
);
