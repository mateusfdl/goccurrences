package listeners

import (
	"context"
	"fmt"

	"github.com/mateusfdl/go-poc/internal/achievements/badges"
	"github.com/mateusfdl/go-poc/internal/occurrences/dto"
	"github.com/mateusfdl/go-poc/internal/occurrences/emitter"
	"github.com/mateusfdl/go-poc/internal/occurrences/events"
	"github.com/mateusfdl/go-poc/internal/occurrences/services"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("badges-listeners",
	fx.Provide(NewListeners),
	fx.Invoke(HookListeners),
)

type Listeners struct {
	accountCreatedBadge *badges.AccountCreatedBadge
	occurrenceService   *services.OccurrenceService
	logger              *zap.SugaredLogger
}

func NewListeners(
	accountCreatedBadge *badges.AccountCreatedBadge,
	occurrenceService *services.OccurrenceService,
	logger *zap.SugaredLogger,
) *Listeners {
	return &Listeners{
		accountCreatedBadge: accountCreatedBadge,
		occurrenceService:   occurrenceService,
		logger:              logger,
	}
}

func (l *Listeners) RegisterListeners() {
	emitter.AddListener(l.AccountCreatedEventHandlers)
}

func (l *Listeners) AccountCreatedEventHandlers(e events.AccountCreatedEvent) error {
	l.logger.Info(fmt.Sprintf("AccountCreatedEventHandlers: %v", e))
	uOccurrences, err := l.occurrenceService.UserOccurrences(
		context.Background(),
		&dto.ListUserOccurrenceDTO{ActorID: e.UserID},
	)
	if err != nil {
		return err
	}
	l.accountCreatedBadge.EvaluateRule(uOccurrences)

	return nil
}

func HookListeners(
	lc fx.Lifecycle,
	listeners *Listeners,
	logger *zap.SugaredLogger,
) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			listeners.RegisterListeners()
			logger.Info("Listeners registered")
			return nil
		},
		OnStop: func(context.Context) error {
			logger.Info("Listeners cleared")
			return nil
		},
	})
}
