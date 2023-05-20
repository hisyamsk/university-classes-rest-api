package helper

import (
	"fmt"
	"time"
)

func GetTimeFromString(timestr string) string {
	t, err := time.Parse(time.RFC3339, timestr)
	PanicIfError(err)

	formattedTime := fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
	return formattedTime
}
