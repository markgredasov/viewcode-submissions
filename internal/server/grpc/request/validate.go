package request

import (
	"fmt"

	errs "github.com/markgredasov/viewcode-submissions/internal/errors"
)

type validateRequest interface {
	ValidateAll() error
}

func Validate(req validateRequest) error {
	if err := req.ValidateAll(); err != nil {
		return fmt.Errorf("validation failed: %w: %w", errs.ErrInvalidArgument, err)
	}

	return nil
}
