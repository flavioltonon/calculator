package valueobject

import "github.com/google/uuid"

type RequestID string

func NewRequestID() RequestID {
	return RequestID(uuid.NewString())
}

func (vo RequestID) String() string {
	return string(vo)
}
