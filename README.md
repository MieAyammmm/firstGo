

```markdown
# Recipe Application

Aplikasi manajemen resep dengan backend Golang, frontend Vue.js (TypeScript), dan MongoDB Atlas sebagai database.

## Fitur

- ✅ Menambah resep baru
- ✅ Melihat daftar semua resep
- ✅ Melihat detail resep
- ✅ Mengedit resep yang ada
- ✅ Menghapus resep
- ✅ Validasi input

## Teknologi yang Digunakan

### Backend

- **Golang** (Gin framework)
- **MongoDB Atlas** (Database cloud)
- **MongoDB Go Driver**

### Frontend

- **Vue.js 3** (Composition API)
- **TypeScript**
- **Axios** (HTTP client)
- **Vue Router**

## Struktur Proyek
```

recipe-app/
├── backend/ # Golang backend
│ ├── controllers/ # Controller untuk recipe
│ ├── models/ # Model data
│ ├── routes/ # Definisi routes
│ ├── main.go # Entry point backend
│ └── go.mod # Go modules
│
├── frontend/ # Vue.js frontend
│ ├── public/ # Assets public
│ ├── src/ # Source code
│ │ ├── api/ # API services
│ │ ├── assets/ # Assets
│ │ ├── components/ # Komponen Vue
│ │ ├── router/ # Konfigurasi router
│ │ ├── types/ # TypeScript interfaces
│ │ ├── views/ # Halaman views
│ │ ├── App.vue # Komponen utama
│ │ └── main.ts # Entry point frontend
│ └── package.json # Dependencies frontend
│
└── README.md # Dokumentasi ini

````

## Prasyarat

- Go 1.16+
- Node.js 14+
- MongoDB Atlas account (atau MongoDB lokal)
- Git

## Instalasi dan Konfigurasi

### Backend Setup

1. Clone repositori:
   ```bash
   git clone https://github.com/MieAyammmm/firstGo.git
   cd recipe-app/backend
````

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Buat file `.env` di folder backend:

   ```env
   MONGODB_URI=mongodb+srv://<username>:<password>@cluster0.mongodb.net/recipe_db?retryWrites=true&w=majority
   PORT=8000
   ```

4. Jalankan backend:
   ```bash
   go run main.go
   ```

### Frontend Setup

1. Masuk ke folder frontend:

   ```bash
   cd ../frontend
   ```

2. Install dependencies:

   ```bash
   npm install
   ```

3. Buat file `.env` di folder frontend:

   ```env
   VUE_APP_API_URL=http://localhost:8000
   ```

4. Jalankan frontend:
   ```bash
   npm run dev
   ```

## Konfigurasi MongoDB Atlas

1. Buat cluster gratis di [MongoDB Atlas](https://www.mongodb.com/atlas/database)
2. Tambahkan user database dengan hak akses read-write
3. Tambahkan IP Anda ke whitelist network access
4. Dapatkan connection string dan masukkan ke `.env` backend

## Endpoint API

| Method | Endpoint     | Deskripsi            |
| ------ | ------------ | -------------------- |
| GET    | /            | Health check         |
| POST   | /recipes     | Buat resep baru      |
| GET    | /recipes     | Dapatkan semua resep |
| GET    | /recipes/:id | Dapatkan resep by ID |
| PUT    | /recipes/:id | Update resep         |
| DELETE | /recipes/:id | Hapus resep          |

## Model Data

```go
type Recipe struct {
    ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    Title       string             `json:"title" bson:"title"`
    Ingredients []string           `json:"ingredients" bson:"ingredients"`
    Steps       []string           `json:"steps" bson:"steps"`
    Duration    int                `json:"duration" bson:"duration"` // dalam menit
    Difficulty string             `json:"difficulty" bson:"difficulty"`
}
```
