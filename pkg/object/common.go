package object

import (
	"github.com/google/uuid"
)

type Spec struct {
	UUID        uuid.UUID
	ProjectID   uuid.UUID
	Name        string
	TypeOf      string
	Status      string
	Description string
}
