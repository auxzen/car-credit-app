-- Insert Kontrak Pak Sugus
INSERT INTO kontrak (kontrak_no, client_name, otr, down_payment, tenor_bulan, tanggal_mulai)
VALUES 
('AGR00001', 'SUGUS', 240000000, 48000000, 12, '2024-01-01');

-- Insert Jadwal Angsuran 12 bulan mulai Januari 2024
INSERT INTO jadwal_angsuran (kontrak_no, angsuran_ke, angsuran_per_bulan, tanggal_jatuh_tempo)
VALUES
('AGR00001', 1, 12907000, '2024-01-25'),
('AGR00001', 2, 12907000, '2024-02-25'),
('AGR00001', 3, 12907000, '2024-03-25'),
('AGR00001', 4, 12907000, '2024-04-25'),
('AGR00001', 5, 12907000, '2024-05-25'),
('AGR00001', 6, 12907000, '2024-06-25'),
('AGR00001', 7, 12907000, '2024-07-25'),
('AGR00001', 8, 12907000, '2024-08-25'),
('AGR00001', 9, 12907000, '2024-09-25'),
('AGR00001',10, 12907000, '2024-10-25'),
('AGR00001',11, 12907000, '2024-11-25'),
('AGR00001',12, 12907000, '2024-12-25');

-- Insert Pembayaran sampai dengan Mei 2024 (angsuran ke-1 s/d ke-5)
INSERT INTO pembayaran (kontrak_no, angsuran_ke, tanggal_bayar)
VALUES
('AGR00001', 1, '2024-01-25'),
('AGR00001', 2, '2024-02-25'),
('AGR00001', 3, '2024-03-25'),
('AGR00001', 4, '2024-04-25'),
('AGR00001', 5, '2024-05-25');


-- QUERY ANGSURAN JATUH TEMPO HINGGA 14 AGUSTUS 2024
SELECT
    k.kontrak_no,
    k.client_name,
    SUM(j.angsuran_per_bulan) AS total_angsuran_jatuh_tempo
FROM
    kontrak k
JOIN
    jadwal_angsuran j ON k.kontrak_no = j.kontrak_no
WHERE
    k.client_name = 'SUGUS'
    AND j.tanggal_jatuh_tempo <= '2024-08-14'
GROUP BY
    k.kontrak_no, k.client_name;


-- QUERY KETELAMBATAN
SELECT 
    j.kontak_no,
    j.client_name,
    j.angsuran_ke,
    GREATEST(DATE_PART('day', DATE '2024-08-14' - j.tanggal_jatuh_tempo), 0) AS hari_ketermbatan,
    ROUND(
        GREATEST(DATE_PART('day', DATE '2024-08-14' - j.tanggal_jatuh_tempo), 0) * 0.001 * j.angsuran_per_bulan
    ) AS total_denda
FROM
    jadwal_angsuran j
JOIN 
    kontrak k ON j.kontrak_no = k.kontrak_no
LEFT JOIN
    pembayaran p ON j.kontak_no = p.kontak_no AND j.angsuran_ke = p.angsuran_ke
WHERE
    k.client_name = 'SUGUS'
    AND j.tanggal_jatuh_tempo <= '2024-08-14'
    AND p.id IS NULL -- Belum ada pembayaran
ORDER BY
    j.angsuran_ke;
