package tools

import "github.com/google/uuid"

func GenUUID() string {
	return uuid.New().String()
}
