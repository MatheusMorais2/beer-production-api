package helpers

import (
	"fmt"
	"time"
)

func IsZero(time time.Time) bool {
	zeroTime := time
	fmt.Println(zeroTime)
	if time == zeroTime {
		return true
	}

	return false
}
