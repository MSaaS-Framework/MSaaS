package object

import "github.com/google/uuid"

type Service struct {
	General   *GeneralSpec
	Neighbors map[string]uuid.UUID
	Databases map[string]uuid.UUID
	API       uuid.UUID
}
