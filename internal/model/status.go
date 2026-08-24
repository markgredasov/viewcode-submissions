package model

import (
	"fmt"

	errs "github.com/markgredasov/viewcode-submissions/internal/errors"
)

type Status string

const (
	StatusPending     = Status("PENDING")
	StatusRunning     = Status("RUNNING")
	StatusAccepted    = Status("ACCEPTED")
	StatusWrongAnswer = Status("WRONG_ANSWER")
	StatusTimeout     = Status("TIMEOUT")
	StatusInternal    = Status("INTERNAL")
)

func NewStatus(status string) (Status, error) {
	statusValueObject := Status(status)
	if !statusValueObject.IsValid() {
		return "", fmt.Errorf("status: %s: %w", status, errs.ErrInvalidArgument)
	}

	return statusValueObject, nil
}

func (s Status) IsValid() bool {
	switch s {
	case StatusPending, StatusRunning, StatusAccepted, StatusWrongAnswer, StatusTimeout, StatusInternal:
		return true
	}

	return false
}
