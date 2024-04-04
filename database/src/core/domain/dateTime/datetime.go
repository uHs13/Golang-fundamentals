package datetime

import "time"

const (
	saoPauloTimezone = "America/Sao_Paulo"
	dateFormat       = "2006-01-02 15:04:05"
)

func BuildFormattedCurrentDate() string {
	return (time.Now()).Format(dateFormat)
}
