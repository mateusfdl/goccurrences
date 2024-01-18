package occurrences

import (
	"github.com/mateusfdl/go-poc/internal/occurrences/repository"
	"github.com/mateusfdl/go-poc/internal/occurrences/services"
	"go.uber.org/fx"
)

var Module = fx.Module("occurrences",
	fx.Provide(services.New),
	fx.Provide(repository.NewOccurrenceRepository),
	fx.Provide(func(repo *repository.MongoOccurrenceRepository) repository.OccurrenceRepository { return repo }),
)
