package api

import "time"

type Info struct {
	Version string    // version string
	Time    time.Time // commit time
}
