package submissionshandler

import (
	"time"

	"github.com/markgredasov/viewcode-submissions/internal/model"
	submissions_v1 "github.com/markgredasov/viewcode-submissions/pkg/api/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func fromRequestToModel(req *submissions_v1.CreateRequest) (model.Submission, error) {
	statusModel, err := model.NewStatusFromGRPC(req.GetStatus())
	if err != nil {
		return model.Submission{}, err
	}

	languageModel, err := model.NewLanguageFromGRPC(req.GetLanguage())
	if err != nil {
		return model.Submission{}, err
	}

	submission, err := model.NewSubmission(
		req.GetUserId(),
		req.GetProblemId(),
		int(req.GetTestsTotal()),
		int(req.GetTestsPassed()),
		statusModel.String(),
		languageModel.String(),
		int(req.GetExecutionTimeMs()),
		time.Now(),
	)

	if err != nil {
		return model.Submission{}, err
	}

	return submission, nil
}

func fromModelToResponse(m model.Submission) (*submissions_v1.CreateResponse, error) {
	statusModel, err := m.Status.GRPCModel()
	if err != nil {
		return nil, err
	}

	languageModel, err := m.Language.GRPCModel()
	if err != nil {
		return nil, err
	}

	reqSubmission := &submissions_v1.Submission{
		Id:              m.ID.String(),
		UserId:          m.UserID.String(),
		ProblemId:       m.ProblemID.String(),
		TestsTotal:      int32(m.TestsTotal),
		TestsPassed:     int32(m.TestsPassed),
		Status:          statusModel,
		Language:        languageModel,
		ExecutionTimeMs: int32(m.ExecutionTimeInMS),
		CreatedAt:       timestamppb.New(m.CreatedAt),
	}

	resp := &submissions_v1.CreateResponse{
		Submission: reqSubmission,
	}

	return resp, nil
}
