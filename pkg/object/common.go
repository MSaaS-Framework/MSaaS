package object

import (
	"github.com/google/uuid"
)

type GeneralSpec struct {
	UUID        uuid.UUID
	ProjectID   uuid.UUID
	Name        string
	Object      string
	Status      string
	Description string
}
