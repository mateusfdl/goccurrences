package repository

import (
	"context"
	"errors"
	"log"

	mongodb "github.com/mateusfdl/go-poc/internal/mongo"
	"github.com/mateusfdl/go-poc/internal/occurrences/dto"
	"github.com/mateusfdl/go-poc/internal/occurrences/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoOccurrenceRepository struct {
	collection *mongo.Collection
}

func NewOccurrenceRepository(db *mongodb.Mongo) *MongoOccurrenceRepository {
	return &MongoOccurrenceRepository{collection: db.DB.Collection("occurrences")}
}

func (c *MongoOccurrenceRepository) Create(
	ctx context.Context,
	dto *dto.CreateOccurrenceDTO,
) (string, error) {
	doc, err := c.collection.InsertOne(ctx, dto)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	oid, ok := doc.InsertedID.(primitive.ObjectID)

	if !ok {
		log.Fatal("error casting to object id")
		return "", errors.New("error casting to object id")

	}

	return oid.Hex(), nil
}

func (c *MongoOccurrenceRepository) List(
	ctx context.Context,
	ID string,
	Limit uint32,
	Skip uint32,
) ([]entity.Occurrence, error) {
	var occurrences []entity.Occurrence
	stages := bson.A{
		bson.D{{Key: "$match", Value: bson.D{{Key: "user_id", Value: ID}}}},
	}

	cursor, err := c.collection.Aggregate(ctx, stages)
	if err != nil {
		return nil, ErrAggregationPipeline
	}
	if err = cursor.All(context.TODO(), &occurrences); err != nil {
		return nil, ErrListUserOccurrences

	}

	return occurrences, nil
}
