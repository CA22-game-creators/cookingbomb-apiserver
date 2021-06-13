package account

import "github.com/google/uuid"

type OutputData struct {
	SessionToken uuid.UUID
	Err          error
}
