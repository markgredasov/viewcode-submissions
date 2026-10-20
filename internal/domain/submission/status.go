package submission

import (
	"fmt"

	errs "github.com/markgredasov/viewcode-submissions/internal/errors"
	submissions_v1 "github.com/markgredasov/viewcode-submissions/pkg/api/v1"
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

var (
	statusModelToGRPCModel = map[Status]submissions_v1.Status{
		StatusPending:     submissions_v1.Status_STATUS_PENDING,
		StatusRunning:     submissions_v1.Status_STATUS_RUNNING,
		StatusAccepted:    submissions_v1.Status_STATUS_ACCEPTED,
		StatusWrongAnswer: submissions_v1.Status_STATUS_WRONG_ANSWER,
		StatusTimeout:     submissions_v1.Status_STATUS_TIMEOUT,
		StatusInternal:    submissions_v1.Status_STATUS_INTERNAL,
	}

	//nolint:exhaustive // unspecified not needed in map
	statusGRPCModelToModel = map[submissions_v1.Status]Status{
		submissions_v1.Status_STATUS_PENDING:      StatusPending,
		submissions_v1.Status_STATUS_RUNNING:      StatusRunning,
		submissions_v1.Status_STATUS_ACCEPTED:     StatusAccepted,
		submissions_v1.Status_STATUS_WRONG_ANSWER: StatusWrongAnswer,
		submissions_v1.Status_STATUS_TIMEOUT:      StatusTimeout,
		submissions_v1.Status_STATUS_INTERNAL:     StatusInternal,
	}
)

func NewStatus(status string) (Status, error) {
	statusValueObject := Status(status)
	if !statusValueObject.IsValid() {
		return "", fmt.Errorf("status: %s: %w", status, errs.ErrInvalidArgument)
	}

	return statusValueObject, nil
}

func ConvertStatusFromGRPC(grpcModel submissions_v1.Status) (Status, error) {
	statusModel, ok := statusGRPCModelToModel[grpcModel]
	if !ok {
		return "", fmt.Errorf("status: %s: %w", grpcModel.String(), errs.ErrInvalidArgument)
	}

	return statusModel, nil
}

func (s Status) String() string {
	return string(s)
}

func (s Status) IsValid() bool {
	switch s {
	case StatusPending, StatusRunning, StatusAccepted, StatusWrongAnswer, StatusTimeout, StatusInternal:
		return true
	}

	return false
}

func (s Status) IsTerminal() bool {
	switch s {
	case StatusAccepted, StatusWrongAnswer, StatusTimeout, StatusInternal:
		return true
	}
	return false
}

func (s Status) GRPCModel() (submissions_v1.Status, error) {
	grpcModel, ok := statusModelToGRPCModel[s]
	if !ok {
		return submissions_v1.Status_STATUS_UNSPECIFIED, fmt.Errorf("no such status: %s: %w", s.String(), errs.ErrInvalidArgument)
	}
	return grpcModel, nil
}
