package utils

import (
	"fmt"
	"strconv"
	"time"
)

func AutoGenerateBookingCode() string {
	now := time.Now()
	year, month, day := now.Date()
	timestamp := now.Unix()

	return "BK" + strconv.Itoa(year)[2:] +
		fmt.Sprintf("%02d", int(month)) +
		fmt.Sprintf("%02d", day) +
		fmt.Sprintf("%06d", timestamp%1000000)
}
