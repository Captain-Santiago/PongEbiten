package savegame

import (
	"time"
)

type Save struct {
	Name   string
	Points int
	Date   time.Time
}
