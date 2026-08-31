package submissionsrepository

import (
	"time"

	"github.com/markgredasov/viewcode-submissions/internal/model"
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
}

func fromRecordToModel(r record) (model.Submission, error) {
	m, err := model.NewSubmission(r.UserID, r.ProblemID, r.TestsTotal, r.TestsPassed, r.Status,
		r.Language, r.ExecutionTimeInMS, r.CreatedAt)
	if err != nil {
		return model.Submission{}, err
	}

	uid, err := model.ToUUID(r.ID)
	if err != nil {
		return model.Submission{}, err
	}

	return *m.WithID(uid), nil
}
