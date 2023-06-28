package utility

import (
	"time"
)

func GetCurrentTime() int {
	return int(time.Now().Unix())
}
