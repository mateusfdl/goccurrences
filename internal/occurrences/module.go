package occurrences

import (
	"context"

	occurrencesGrpcServer "buf.build/gen/go/matheusslima/go-poc/grpc/go/occurrences/v1/occurrencesv1grpc"
	"github.com/mateusfdl/go-poc/internal/mongo"
	"github.com/mateusfdl/go-poc/internal/occurrences/emitter"
	occurrenceGrpcInternal "github.com/mateusfdl/go-poc/internal/occurrences/grpc"
	"github.com/mateusfdl/go-poc/internal/occurrences/repository"
	"github.com/mateusfdl/go-poc/internal/occurrences/services"
	"go.uber.org/fx"
	"go.uber.org/zap"
	grpcServer "google.golang.org/grpc"
)

var Module = fx.Module("occurrences",
	fx.Provide(occurrenceGrpcInternal.NewOccurrenceHandler),
	fx.Provide(services.NewOccurrenceService),
	fx.Provide(repository.NewOccurrenceRepository),
	fx.Provide(func(repo *repository.MongoOccurrenceRepository) repository.OccurrenceRepository { return repo }),
	fx.Invoke(
		func(s *grpcServer.Server, h *occurrenceGrpcInternal.OccurrenceHandler) {
			occurrencesGrpcServer.RegisterOccurrenceServiceServer(s, h)
		},
	),
	fx.Invoke(emitter.New),
	fx.Invoke(HookDatabaseIndexesSync),
)

func HookDatabaseIndexesSync(
	lc fx.Lifecycle,
	mongo *mongo.Mongo,
	logger *zap.Logger,
) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			err := repository.HookSyncOccurrencesIndexes(
				context.Background(),
				mongo,
				logger,
			)

			if err != nil {
				logger.Error("failed to sync indexes", zap.Error(err))
				return err
			}

			return nil
		}})
}
