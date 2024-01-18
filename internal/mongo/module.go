package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var (
	Module    = fx.Module("mongo", providers, invokers)
	providers = fx.Provide(
		New,
	)
	invokers = fx.Invoke(
		HookConnection,
	)
)

type Mongo struct {
	DB *mongo.Database
}

func New(logger *zap.Logger) *Mongo {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/rewards-poc?authSource=admin&retryWrites=true&w=majority")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		logger.Error("failed to connect to mongo")
	}

	db := client.Database("rewards-poc")
	return &Mongo{DB: db}
}

func HookConnection(lc fx.Lifecycle, client *mongo.Client, logger *zap.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			err := client.Ping(ctx, nil)
			if err != nil {
				logger.Error("mongo is dead", zap.Error(err))
				return err
			}

			logger.Info("mongo is alive")

			return nil
		},
		OnStop: func(ctx context.Context) error {
			err := client.Disconnect(ctx)
			if err != nil {
				logger.Error("failed to gracefully disconnect from mongo", zap.Error(err))
				return err
			}

			return nil
		},
	})
}
