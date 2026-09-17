package submissionshandler

import (
	"github.com/markgredasov/viewcode-submissions/internal/model"
	"github.com/markgredasov/viewcode-submissions/internal/model/submission"
	submissions_v1 "github.com/markgredasov/viewcode-submissions/pkg/api/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func createToSubmission(req *submissions_v1.CreateRequest) (submission.Submission, error) {
	languageModel, err := model.NewLanguageFromGRPC(req.GetLanguage())
	if err != nil {
		return submission.Submission{}, err
	}

	s, err := submission.New(req.GetUserId(), req.GetProblemId(), int(req.GetTestsTotal()),
		submission.TestsPassedUnitialized, languageModel, submission.ExecutionTimeInMSUnitialized)

	if err != nil {
		return submission.Submission{}, err
	}

	return s, nil
}

func toSubmission(m submission.Submission) (*submissions_v1.Submission, error) {
	statusModel, err := m.Status.GRPCModel()
	if err != nil {
		return nil, err
	}

	languageModel, err := m.Language.GRPCModel()
	if err != nil {
		return nil, err
	}

	return &submissions_v1.Submission{
		Id:              m.ID.String(),
		UserId:          m.UserID.String(),
		ProblemId:       m.ProblemID.String(),
		TestsTotal:      int32(m.TestsTotal),
		TestsPassed:     int32(m.TestsPassed),
		Status:          statusModel,
		Language:        languageModel,
		ExecutionTimeMs: int32(m.ExecutionTimeInMS),
		CreatedAt:       timestamppb.New(m.CreatedAt),
		UpdatedAt:       timestamppb.New(m.UpdatedAt),
	}, nil
}

func toSubmissions(m []submission.Submission) ([]*submissions_v1.Submission, error) {
	submissions := make([]*submissions_v1.Submission, len(m))
	for i := range m {
		s, err := toSubmission(m[i])
		if err != nil {
			return nil, err
		}

		submissions[i] = s
	}
	return submissions, nil
}

func createToResponse(m submission.Submission) (*submissions_v1.CreateResponse, error) {
	s, err := toSubmission(m)
	if err != nil {
		return nil, err
	}

	resp := &submissions_v1.CreateResponse{
		Submission: s,
	}

	return resp, nil
}

func getByIDToResponse(m submission.Submission) (*submissions_v1.GetByIDResponse, error) {
	s, err := toSubmission(m)
	if err != nil {
		return nil, err
	}

	resp := &submissions_v1.GetByIDResponse{
		Submission: s,
	}

	return resp, nil
}

func listAllToResponse(m []submission.Submission) (*submissions_v1.ListAllResponse, error) {
	s, err := toSubmissions(m)
	if err != nil {
		return nil, err
	}

	resp := &submissions_v1.ListAllResponse{
		Submissions: s,
	}

	return resp, nil
}

func setRunningToResponse(m submission.Submission) (*submissions_v1.SetRunningResponse, error) {
	s, err := toSubmission(m)
	if err != nil {
		return nil, err
	}

	resp := &submissions_v1.SetRunningResponse{
		Submission: s,
	}

	return resp, nil
}

func setResultToResponse(m submission.Submission) (*submissions_v1.SetResultResponse, error) {
	s, err := toSubmission(m)
	if err != nil {
		return nil, err
	}

	resp := &submissions_v1.SetResultResponse{
		Submission: s,
	}

	return resp, nil
}
