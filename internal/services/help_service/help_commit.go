package help_service

import (
	"context"

	"github.com/sergeyiksanov/help-on-road/internal/models"
)

func (hs *HelpService) CommitHelp(ctx context.Context, userId int64, status models.HelpCallStatus) error {
	helps, err := hs.helpRepository.GetByUserId(ctx, userId)
	if err != nil {
		return err
	}

	help := helps[0]
	help.Status = status

	if err := hs.helpRepository.Update(ctx, userId, 0, help); err != nil {
		return err
	}

	// if err := hs.helpRepository.Delete(ctx, userId); err != nil {
	// return err
	// }

	return nil
}
