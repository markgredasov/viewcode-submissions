package submissionsrepository

import (
	"time"

	"github.com/markgredasov/viewcode-submissions/internal/model"
	"github.com/markgredasov/viewcode-submissions/internal/model/submission"
)

type record struct {
	ID                string
	UserID            string
	ProblemID         string
	TestsTotal        int
	TestsPassed       int
	Status            string
	Language          string
	ExecutionTimeInMS int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func toModel(r record) (submission.Submission, error) {
	languageModel, err := model.NewLanguage(r.Language)
	if err != nil {
		return submission.Submission{}, err
	}

	m, err := submission.New(r.UserID, r.ProblemID, r.TestsTotal, r.TestsPassed, languageModel, r.ExecutionTimeInMS)
	if err != nil {
		return submission.Submission{}, err
	}

	uid, err := model.ToUUID(r.ID)
	if err != nil {
		return submission.Submission{}, err
	}
	m.ID = uid

	status, err := submission.NewStatus(r.Status)
	if err != nil {
		return submission.Submission{}, err
	}
	m.Status = status

	m.CreatedAt = r.CreatedAt
	m.UpdatedAt = r.UpdatedAt

	return m, nil
}
