package submission

import (
	"slices"
	"time"
	"uuid"

	domain "github.com/markgredasov/viewcode-submissions/internal/domain"
	errs "github.com/markgredasov/viewcode-submissions/internal/errors"
)

const (
	TestsPassedUnitialized       = -1
	ExecutionTimeInMSUnitialized = -1
)

type Submission struct {
	ID                uuid.UUID
	UserID            uuid.UUID
	ProblemID         uuid.UUID
	TestsTotal        int
	TestsPassed       int
	Status            Status
	Language          domain.Language
	ExecutionTimeInMS int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func New(
	userID string,
	problemID string,
	testsTotal int,
	testsPassed int,
	language domain.Language,
	executionTimeInMs int,
) (Submission, error) {
	id := uuid.New()

	userIDValueObject, err := uuid.Parse(userID)
	if err != nil {
		return Submission{}, err
	}

	problemIDValueObject, err := uuid.Parse(problemID)
	if err != nil {
		return Submission{}, err
	}

	return Submission{
		ID:                id,
		UserID:            userIDValueObject,
		ProblemID:         problemIDValueObject,
		TestsTotal:        testsTotal,
		TestsPassed:       testsPassed,
		Status:            StatusPending,
		Language:          language,
		ExecutionTimeInMS: executionTimeInMs,
	}, nil
}

func (s *Submission) WithID(id uuid.UUID) *Submission {
	s.ID = id
	return s
}

func (s *Submission) GetAvailableNextStatuses() []Status {
	switch s.Status {
	case StatusPending:
		return []Status{StatusRunning}
	case StatusRunning:
		return []Status{StatusAccepted, StatusInternal, StatusTimeout, StatusWrongAnswer}
	default:
		return []Status{}
	}
}

func (s *Submission) AdvanceTo(next Status) error {
	if !next.IsValid() {
		return errs.InvalidArgument("unknown submissions status: " + next.String())
	}

	statuses := s.GetAvailableNextStatuses()
	if !slices.Contains(statuses, next) {
		return errs.InvalidArgument("cannot transition submission status from " + s.Status.String() + " to " + next.String())
	}

	s.Status = next
	return nil
}
