package model

import "time"

type Submission struct {
	ID                UUID
	UserID            UUID
	ProblemID         UUID
	TestsTotal        int
	TestsPassed       int
	Status            Status
	Language          Language
	ExecutionTimeInMS int
	CreatedAt         time.Time
}

func NewSubmission(
	userID string,
	problemID string,
	testsTotal int,
	testsPassed int,
	status string,
	language string,
	executionTimeInMs int,
	createdAt time.Time,
) (Submission, error) {
	id := NewUUID()

	userIDValueObject, err := ToUUID(userID)
	if err != nil {
		return Submission{}, err
	}

	problemIDValueObject, err := ToUUID(problemID)
	if err != nil {
		return Submission{}, err
	}

	statusValueObject, err := NewStatus(status)
	if err != nil {
		return Submission{}, err
	}

	languageValueObject, err := NewLanguage(language)
	if err != nil {
		return Submission{}, err
	}

	return Submission{
		ID:                id,
		UserID:            userIDValueObject,
		ProblemID:         problemIDValueObject,
		TestsTotal:        testsTotal,
		TestsPassed:       testsPassed,
		Status:            statusValueObject,
		Language:          languageValueObject,
		ExecutionTimeInMS: executionTimeInMs,
		CreatedAt:         createdAt,
	}, nil
}

func (s *Submission) WithID(id UUID) *Submission {
	s.ID = id
	return s
}
