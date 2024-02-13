package repository

import (
	"context"
	"log"

	mongodb "github.com/mateusfdl/go-poc/internal/mongo"
	"github.com/mateusfdl/go-poc/internal/occurrences/dto"
	"github.com/mateusfdl/go-poc/internal/occurrences/entity"
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
	dto dto.CreateOccurrenceDTO,
) (string, error) {
	doc, err := c.collection.InsertOne(ctx, dto)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return doc.InsertedID.(string), nil
}

func (c *MongoOccurrenceRepository) List(
	ctx context.Context,
	ID string,
) ([]entity.Occurrence, error) {
	var occurrences []entity.Occurrence
	cursor, err := c.collection.Find(ctx, dto.CreateOccurrenceDTO{UserID: ID})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if err = cursor.All(ctx, &occurrences); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return occurrences, nil
}
