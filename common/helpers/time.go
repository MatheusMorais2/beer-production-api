package helpers

import (
	"time"
)

func IsZero(time time.Time) bool {
	zeroTime := time
	if time == zeroTime {
		return true
	}

	return false
}
