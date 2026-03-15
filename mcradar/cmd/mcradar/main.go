package main

import (
	"sync"

	"github.com/joho/godotenv"
	"github.com/local-interloper/mc-radar/mcradar/internal/db"
	"github.com/local-interloper/mc-radar/mcradar/internal/scanning"
	"github.com/local-interloper/mc-radar/mcradar/internal/settings"
)

func main() {
	wg := &sync.WaitGroup{}

	godotenv.Load()

	settings.Init()

	db.Init()

	scanning.BeginFullRangeScan(wg)

	wg.Wait()
}
