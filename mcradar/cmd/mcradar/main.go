package main

import (
	"github.com/joho/godotenv"
	"github.com/local-interloper/mc-radar/mcradar/internal/db"
	"github.com/local-interloper/mc-radar/mcradar/internal/scanning"
)

func main() {
	godotenv.Load()

	db.Init()

	scanning.BeginFullRangeScan()
}
