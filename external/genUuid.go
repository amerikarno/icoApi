package external

import "github.com/google/uuid"

func GenUuid() (uid string) {
	return uuid.New().String()
}
