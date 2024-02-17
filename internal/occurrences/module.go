package occurrences

import (
	occurrencesGrpcServer "buf.build/gen/go/matheusslima/go-poc/grpc/go/occurrences/v1/occurrencesv1grpc"
	occurrenceGrpcInternal "github.com/mateusfdl/go-poc/internal/occurrences/grpc"
	"github.com/mateusfdl/go-poc/internal/occurrences/repository"
	"github.com/mateusfdl/go-poc/internal/occurrences/services"
	"go.uber.org/fx"
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
)
