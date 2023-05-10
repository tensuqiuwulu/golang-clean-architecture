package utils

import (
	guuid "github.com/google/uuid"
)

func RandomUUID() string {
	return guuid.NewString()
}
