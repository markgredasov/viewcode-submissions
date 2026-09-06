package submission

import (
	"time"

	"github.com/markgredasov/viewcode-submissions/internal/model"
)

const (
	TestsPassedUnitialized       = -1
	ExecutionTimeInMSUnitialized = -1
)

type Submission struct {
	ID                model.UUID
	UserID            model.UUID
	ProblemID         model.UUID
	TestsTotal        int
	TestsPassed       int
	Status            Status
	Language          model.Language
	ExecutionTimeInMS int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func New(
	userID string,
	problemID string,
	testsTotal int,
	testsPassed int,
	language model.Language,
	executionTimeInMs int,
) (Submission, error) {
	id := model.NewUUID()

	userIDValueObject, err := model.ToUUID(userID)
	if err != nil {
		return Submission{}, err
	}

	problemIDValueObject, err := model.ToUUID(problemID)
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

func (s *Submission) WithID(id model.UUID) *Submission {
	s.ID = id
	return s
}
