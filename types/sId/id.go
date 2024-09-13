package sId

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

// UUID attempts to generate a UUID up to 3 times.
// It returns a UUID on success or an error if it fails after 3 attempts.
func UUID() (uuid.UUID, error) {
	var id uuid.UUID
	var err error

	for i := 0; i < 3; i++ {
		id, err = uuid.NewV7()
		if err == nil {
			// If the UUID was generated successfully, return it
			return id, nil
		}

		time.Sleep(100 * time.Millisecond) // Sleep for 100 milliseconds before retrying
	}

	// If unable to generate a UUID after 3 attempts, return an error
	return uuid.UUID{}, fmt.Errorf("error attempting to generate a UUID on 3 attempt: %v", err)
}

func CustomUUID(length int) (string, error) {

	uuid, err := uuid.NewV7()
	if err != nil {
		return "", fmt.Errorf("error attempting to generate a UUID on 3 attempt: %v", err)
	}

	id := strings.Replace(uuid.String(), "-", "", -1)

	if length >= len(id) {
		return id, nil
	}

	return id[len(id)-length:], nil
}
