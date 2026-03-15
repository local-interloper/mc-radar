package main

import (
	"github.com/joho/godotenv"
	"github.com/local-interloper/mc-radar/mcradar/internal/db"
	"github.com/local-interloper/mc-radar/mcradar/internal/scanning"
	"github.com/local-interloper/mc-radar/mcradar/internal/settings"
)

func main() {
	godotenv.Load()

	settings.Init()

	db.Init()

	scanning.BeginFullRangeScan()
}
