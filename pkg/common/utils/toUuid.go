package utils

import (
	"github.com/gofrs/uuid"
)

func ToUuid(id string) (uid uuid.UUID) {
	return uuid.Must(uuid.FromString(id))
}