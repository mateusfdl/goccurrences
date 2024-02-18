package badges

import (
	"github.com/mateusfdl/go-poc/internal/occurrences/entity"
	"go.uber.org/zap"
)

type AccountCreatedBadge struct {
  logger *zap.SugaredLogger
}

func (b *AccountCreatedBadge) EvaluateRule([]*entity.Occurrence) error {
  b.logger.Info("Account create badge rule evaluated")
  return nil
}
