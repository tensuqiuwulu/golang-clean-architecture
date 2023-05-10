package utils

import (
	"strconv"
	"time"
)

func GenerateTableName(name string, date time.Time) string {
	tableName := name + strconv.Itoa(date.Year())
	if date.Month() < 10 {
		tableName += "_0" + strconv.Itoa(int(date.Month()))
	} else {
		tableName += "_" + strconv.Itoa(int(date.Month()))
	}

	return tableName
}
