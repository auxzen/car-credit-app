# App Information 
This app for simulation credit car called ims-finance

## Stack Technology 
1. Go Language
2. PostgreSQL
3. GIN Gonic

## How to run application 
```
Create file .env with this structure code:

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your-password
DB_NAME=your-db

```

running API:
```
localhost:8080/<end-point>
```


### | The  End Point 
```
POST /kontrak
Content-Type: application/json

{
  "kontrak_no": "AGR00001",
  "client_name": "SUGUS",
  "otr": 240000000,
  "down_payment": 48000000,
  "tenor_bulan": 12
}
```
```
POST /pembayaran
Content-Type: application/json

{
  "kontrak_no": "AGR00001",
  "angsuran_ke": 6,
  "tanggal_bayar": "2024-06-25"
}
```
```
GET /kontrak/AGR00001/denda?tanggal=2024-08-14

```