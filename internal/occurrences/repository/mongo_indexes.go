package repository

import (
	"context"

	mongodb "github.com/mateusfdl/go-poc/internal/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func HookSyncOccurrencesIndexes(
	ctx context.Context,
	db *mongodb.Mongo,
	log *zap.Logger,
) error {
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "actorId", Value: 1},
			{Key: "actorType", Value: 1},
			{Key: "sourceId", Value: 1},
			{Key: "sourceType", Value: 1},
		},
		Options: options.Index().SetName("actor_source_compound_index"),
	}

	_, err := db.DB.Collection("occurrences").Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Error("failed to create index on 'occurrences' collection", zap.Error(err))
		return err
	}

	log.Info("index created on 'occurrences' collection")
	return nil
}
