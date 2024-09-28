package object

import (
	"github.com/google/uuid"
)

type Project struct {
	General   *GeneralSpec
	Services  map[string]uuid.UUID
	Databases map[string]uuid.UUID
	APIs      map[string]uuid.UUID
}
