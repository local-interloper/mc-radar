# mc-radar

![Go](https://img.shields.io/badge/Go-1.26-00ADD8?style=flat-square&logo=go&logoColor=white)
![Python](https://img.shields.io/badge/Python-3.14-3776AB?style=flat-square&logo=python&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-18-336791?style=flat-square&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Compose-2496ED?style=flat-square&logo=docker&logoColor=white)
![License](https://img.shields.io/badge/License-GPL--3.0-blue?style=flat-square)
![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20macOS-lightgrey?style=flat-square)

> A concurrent Minecraft server scanner that sweeps the entire IPv4 space, catalogs every public server it finds, and exposes the data through a REST API.

---

## What it does

mc-radar is a monorepo containing two services:

### mcradar (scanner)

Splits the full 32-bit IPv4 address space into a configurable number of parallel goroutines, connects to port `25565` on each address using a hand-rolled Minecraft Java Edition protocol implementation, and stores every live server it finds into a PostgreSQL database.

For each discovered server it records:

- IP address
- Server version
- Online / max player counts
- Online player sample (names + UUIDs)
- Whether the server runs in online mode (legit) or offline mode (cracked)

### mcradar-api

A FastAPI service that exposes the collected data over a REST API with bearer token authentication, pagination, and filtering.

## Requirements

- Docker Compose (recommended)

Or, to run manually:

- Go 1.26+ (scanner)
- Python 3.14+ with [uv](https://github.com/astral-sh/uv) (API)
- PostgreSQL instance

## Getting started

### With Docker Compose (recommended)

```bash
git clone https://github.com/local-interloper/mc-radar
cd mc-radar
cp .example.env .env
# Edit .env with your desired credentials
docker compose up --build
```

This starts the scanner, the API, and a PostgreSQL 18 database.

### Manually

**Scanner:**

```bash
cd mcradar
cp ../.example.env .env
# Edit .env with your PostgreSQL connection details
go build -o mcradar ./cmd/mcradar/main.go
./mcradar
```

**API:**

```bash
cd mcradar-api
# Set environment variables (see Configuration below)
uv sync
uv run fastapi run app/main.py
```

The API listens on port `8000`.

## Configuration

All configuration is done via environment variables. Copy `.example.env` to `.env` and adjust as needed:

| Variable                   | Description                              | Default       |
|----------------------------|------------------------------------------|---------------|
| `API_KEY`                  | Bearer token for API auth                | —             |
| `POSTGRES_HOST`            | PostgreSQL host                          | `mc-radar-db` |
| `POSTGRES_PASSWORD`        | PostgreSQL password                      | —             |
| `POSTGRES_DB`              | PostgreSQL database name                 | `postgres`    |
| `POSTGRES_MAX_CONNECTIONS` | Max PostgreSQL connections               | —             |
| `APP_WORKERS`              | Number of parallel scan goroutines       | —             |
| `APP_TIMEOUT_MS`           | Connection timeout per host (ms)         | —             |

## API

All endpoints require an `Authorization: Bearer <API_KEY>` header.

### `POST /api/servers`

Returns a paginated list of discovered servers.

**Request body:**

```json
{
  "pagination": {
    "first": 0,
    "last": 50
  },
  "filters": {
    "version": "1.21",
    "type": "Legit"
  }
}
```

- `pagination.first` / `pagination.last` — row range (offset / limit style)
- `filters.version` — regex matched against the version string (optional)
- `filters.type` — one of `"Legit"`, `"Cracked"`, or `"Unknown"` (optional)

**Response:**

```json
{
  "data": [
    {
      "ip": "1.2.3.4",
      "createdAt": "2025-01-01T00:00:00Z",
      "updatedAt": "2025-01-01T00:00:00Z",
      "version": "1.21.4",
      "type": "Legit",
      "onlinePlayers": 3,
      "maxPlayers": 20
    }
  ],
  "total": 1
}
```

## Tech stack

### mcradar
- **[GORM](https://gorm.io/)** — ORM with PostgreSQL driver (`pgx`)
- **[godotenv](https://github.com/joho/godotenv)** — `.env` file loading
- **Minecraft Java Edition protocol** — implemented from scratch (VarInt, VarLong, McString, packets)
- **`sync.WaitGroup`** — concurrent range scanning with known-server deduplication cache

### mcradar-api
- **[FastAPI](https://fastapi.tiangolo.com/)** — async web framework
- **[psycopg v3](https://www.psycopg.org/psycopg3/)** — PostgreSQL driver with connection pooling
- **[pypika](https://github.com/kayak/pypika)** — SQL query builder
- **[uv](https://github.com/astral-sh/uv)** — package and project manager

### Infrastructure
- **Docker Compose** — containerized deployment with health-checked PostgreSQL 18
