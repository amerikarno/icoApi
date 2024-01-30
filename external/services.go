package external

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/google/uuid"
)

type ExternalServices struct{}

func NewExternalServices() *ExternalServices { return &ExternalServices{} }

func (e *ExternalServices) GenUuid() (uid string) {
	return uuid.New().String()
}

func (h *ExternalServices) HashString(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hashedBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashedBytes)
}
