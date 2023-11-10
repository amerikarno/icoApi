package external

import "github.com/google/uuid"

type ExternalUuid struct {}

func NewExternalUuid() *ExternalUuid { return &ExternalUuid{} }

func (e *ExternalUuid) GenUuid() (uid string) {
	return uuid.New().String()
}
