package main

import (
	"github.com/joho/godotenv"
	"github.com/local-interloper/mcradar/internal/db"
	"github.com/local-interloper/mcradar/internal/scanning"
)

func main() {
	godotenv.Load()

	db.Init()

	scanning.BeginFullRangeScan()
}
