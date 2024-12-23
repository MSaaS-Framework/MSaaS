package object

import (
	"github.com/google/uuid"
)

type GeneralSpec struct {
	UUID        uuid.UUID
	ProjectID   uuid.UUID
	UserID      uuid.UUID
	Name        string
	Type        string
	Status      string
	Description string
}
