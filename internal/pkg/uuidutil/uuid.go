package uuidutil

import "github.com/google/uuid"

func GenUuid() string {
	uuidValue := uuid.New()
	return uuidValue.String()
}
