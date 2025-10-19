# Yasin Todo API

Bu proje, Go ve MongoDB kullanılarak geliştirilmiş bir **REST API Todo uygulamasıdır**.  
Tüm ödev kriterlerine uygundur.

**Prepared by: Yasin Sahyar**

## 🚀 Tamamlanan Adımlar

- ✅ Step 1: Go REST API temel yapısı oluşturuldu
- ✅ Step 2: CRUD işlemleri (GET, POST, PUT, DELETE)
- ✅ Step 3: Filtering, Sorting, Pagination, Validation
- ✅ Step 4: Swagger dokümantasyonu (`http://localhost:9000/swagger/index.html`)
- ✅ Step 5: Lokal çalıştırma (`go run main.go`)

## 📡 API Endpoint’leri

- `GET /tasks`
- `POST /tasks`
- `PUT /tasks/{id}`
- `DELETE /tasks/{id}`

## 🗄️ Database

- MongoDB (`yasin_todo` veritabanı, `tasks` koleksiyonu)

## 🏃‍♂️ Çalıştırma

```bash
go mod download
go run main.go
