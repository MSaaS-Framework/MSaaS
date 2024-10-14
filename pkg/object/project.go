package object

import (
	"github.com/google/uuid"
)

type Project struct {
	General *GeneralSpec

	User       *User
	HostPath   string
	GitHubRepo string

	Services  map[string]uuid.UUID
	Databases map[string]uuid.UUID
	APIs      map[string]uuid.UUID
}
